/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"sync"

	"github.com/howeyc/gopass"
)

func login() (string, string) {
	var username string
	var password string

	fmt.Println("Please enter your docker hub credentials")
	fmt.Printf("Username: ")
	fmt.Scanf("%s", &username)

	fmt.Printf("Password: ")
	pass, _ := gopass.GetPasswdMasked()
	password = string(pass)

	return username, password
}

func syncRepo(repo *Repo, destination string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("  " + repo.Name)
	// Clone the repo, checkout the branches
	g := NewGitRepo(repo.Path, destination)
	err := g.Clone()
	if err != nil {
		fmt.Println("Could not sync repo " + repo.Name)
		panic(err)
	}

	// Fetch correct branch and update
	err = g.Fetch()
	if err != nil {
		panic(err)
	}

	err = g.Checkout(repo.Branch)
	if err != nil {
		fmt.Println("Could not checkout repo branch " + repo.Name + ":" + repo.Branch)
		panic(err)
	}

	err = g.Pull()
	if err != nil {
		panic(err)
	}

	repo.gitRepo = g
}

func cloneRepos(opts *Options) (*Definition, error) {
	def, err := loadDefiniton(opts.definition)
	if err != nil {
		return nil, err
	}

	for repo, branch := range opts.overrides {
		def.OverrideBranch(repo, branch)
	}

	// Clone Repos
	fmt.Println(" syncing repos")
	var wg sync.WaitGroup
	wg.Add(len(def.Repos))

	for i := 0; i < len(def.Repos); i++ {
		go syncRepo(def.Repos[i], opts.home, &wg)
	}

	wg.Wait()

	return def, nil
}

func generateOutputFile(opts *Options, def *Definition) error {
	tpl, err := loadTemplate(opts.template)
	if err != nil {
		return nil
	}

	// Build output definition
	fmt.Println(" connecting to docker host")
	dh, err := NewDockerHost(opts.host)
	if err != nil {
		return err
	}

	fmt.Println(" updating docker images library")
	err = dh.UpdateImages()
	if err != nil {
		return err
	}

	for _, repo := range def.Repos {
		var image string

		commit, cerr := repo.gitRepo.CommitID()
		if cerr != nil {
			return err
		}

		if opts.isRelease {
			image = fmt.Sprintf("%s/%s:%s", opts.org, repo.Name, opts.releasever)
		} else {
			image = fmt.Sprintf("%s:%s", repo.Name, commit)
		}

		s := Service{
			Image:        image,
			Entrypoint:   repo.Entrypoint,
			Restart:      repo.Restart,
			Volumes:      repo.Volumes,
			Ports:        repo.Ports,
			Links:        repo.Links,
			Dependencies: repo.Dependencies,
			Environment:  repo.Environment,
		}

		if !dh.ImageExists(image) {
			s.Build = repo.gitRepo.DeployPath()
		}

		tpl.Services[repo.Name] = s
	}

	fmt.Println(" writing output")
	err = tpl.WriteFile(opts.output)
	if err != nil {
		return err
	}

	return nil
}

func buildAndPush(opts *Options, def *Definition) error {
	dh, err := NewDockerHost(opts.host)
	if err != nil {
		fmt.Println(err)
	}
	dh.SetAuthCredentials(opts.username, opts.password)

	fmt.Println(" building images")
	for _, repo := range def.Repos {
		name := fmt.Sprintf("%s/%s:%s", opts.org, repo.Name, opts.releasever)
		fmt.Println("  " + name)
		_, berr := dh.BuildImage(name, repo.gitRepo.DeployPath())
		if berr != nil {
			return berr
		}
	}

	fmt.Println(" uploading images")
	for _, repo := range def.Repos {
		fmt.Println("  " + name)
		name := fmt.Sprintf("%s/%s:%s", opts.org, repo.Name, opts.releasever)
		_, uerr := dh.PushImage(name)
		if uerr != nil {
			return uerr
		}
	}

	return nil
}
