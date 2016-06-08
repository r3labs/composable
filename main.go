/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"os"
	"sync"
)

func syncRepo(repo *Repo, destination string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Clone the repo, checkout the branches
	g := NewGitRepo(repo.Path, destination)
	err := g.Clone()
	if err != nil {
		panic(err)
	}

	// Fetch correct branch and update
	err = g.Fetch()
	if err != nil {
		panic(err)
	}

	err = g.Checkout(repo.Branch)
	if err != nil {
		panic(err)
	}

	err = g.Pull()
	if err != nil {
		panic(err)
	}

	repo.gitRepo = g
}

func generateOutputFile(opts *Options) error {
	def, err := loadDefiniton(opts.definition)
	if err != nil {
		return err
	}

	tpl, err := loadTemplate(opts.template)
	if err != nil {
		return nil
	}

	for repo, branch := range opts.overrides {
		def.OverrideBranch(repo, branch)
	}

	// Clone Repos
	fmt.Println(" Syncing Repos")
	var wg sync.WaitGroup
	wg.Add(len(def.Repos))

	for i := 0; i < len(def.Repos); i++ {
		go syncRepo(def.Repos[i], opts.home, &wg)
	}

	wg.Wait()

	// Build output definition
	fmt.Println(" Connecting to Docker Host")
	dh, err := NewDockerHost(opts.host)
	if err != nil {
		return err
	}

	fmt.Println(" Updating Docker Images Library")
	err = dh.UpdateImages()
	if err != nil {
		return err
	}

	for _, repo := range def.Repos {
		commit, cerr := repo.gitRepo.CommitID()
		if cerr != nil {
			return err
		}

		image := fmt.Sprintf("%s:%s", repo.Name, commit)
		s := Service{
			Image:        image,
			Volumes:      repo.Volumes,
			Ports:        repo.Ports,
			Links:        repo.Links,
			Dependencies: repo.Dependencies,
		}

		if !dh.ImageExists(image) {
			s.Build = repo.gitRepo.deploymentPath
		}

		tpl.Services[repo.Name] = s
	}

	fmt.Println(" Writing Output")
	err = tpl.WriteFile(opts.output)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	mode, opts := GetOptions()

	switch mode {
	case "gen", "generate":
		fmt.Println("Generating Output Definition")
		err := generateOutputFile(&opts)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "destroy":
		fmt.Println("Destroying")
	}
}
