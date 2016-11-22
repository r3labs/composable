/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/r3labs/composable/dockerhost"
	"github.com/r3labs/composable/git"
	"gopkg.in/yaml.v2"
)

// Definition of repos
type Definition struct {
	Repos []*Repo  `yaml:"repos"`
	opts  *Options `yaml:"-"`
}

// Repo definition
type Repo struct {
	Name         string            `yaml:"name"`
	Path         string            `yaml:"path"`
	Branch       string            `yaml:"branch"`
	Entrypoint   string            `yaml:"entrypoint,omitempty"`
	Restart      string            `yaml:"restart"`
	Volumes      []string          `yaml:"volumes"`
	Ports        []string          `yaml:"ports"`
	Links        []string          `yaml:"links"`
	Dependencies []string          `yaml:"depends"`
	Environment  map[string]string `yaml:"environment"`
	gitRepo      *git.Repo         `yaml:"-"`
}

// LoadDefiniton the input definition
func LoadDefiniton(path string, opts *Options) (*Definition, error) {
	d := Definition{
		opts: opts,
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &d, err
	}

	err = yaml.Unmarshal(data, &d)
	if err != nil {
		return &d, err
	}

	// Ommit/Exclude repos
	for _, repo := range opts.excludes {
		d.ExcludeRepo(repo)
	}

	// Override branches
	if opts.globalbranch != "" {
		for _, repo := range d.Repos {
			d.OverrideBranch(repo.Name, opts.globalbranch)
		}
	}

	for repo, branch := range d.opts.overrides {
		d.OverrideBranch(repo, branch)
	}

	return &d, nil
}

// CloneRepos clones and checks out the correct branch for a repo
func (d *Definition) CloneRepos() error {
	var wg sync.WaitGroup
	wg.Add(len(d.Repos))

	for i := 0; i < len(d.Repos); i++ {
		go func(wg *sync.WaitGroup, d *Definition, i int) {
			defer wg.Done()

			fmt.Printf("  %s\n", d.Repos[i].Name)
			r, err := git.Clone(d.Repos[i].Path, d.opts.home)
			if err != nil {
				panic(err)
			}

			err = r.Sync(d.Repos[i].Branch)
			if err != nil {
				panic(err)
			}

			d.Repos[i].gitRepo = r
		}(&wg, d, i)
	}

	wg.Wait()

	return nil
}

// OverrideBranch updates a repo's branch
func (d *Definition) OverrideBranch(repo, branch string) {
	for i := 0; i < len(d.Repos); i++ {
		if d.Repos[i].Name == repo {
			d.Repos[i].Branch = branch
		}
	}
}

// ExcludeRepo removes a repo from a list based on name
func (d *Definition) ExcludeRepo(repo string) {
	wildcard := strings.Contains(repo, "*")
	if wildcard {
		repo = strings.Replace(repo, "*", "", -1)
	}

	for i := len(d.Repos) - 1; i >= 0; i-- {
		if wildcard && strings.Contains(d.Repos[i].Name, repo) || d.Repos[i].Name == repo {
			d.Repos = append(d.Repos[:i], d.Repos[i+1:]...)
		}
	}
}

// BuildImages builds all images defined on the definition
func (d *Definition) BuildImages() error {
	dh, err := dockerhost.New(d.opts.host)
	if err != nil {
		return err
	}
	dh.SetAuthCredentials(d.opts.username, d.opts.password)

	err = dh.UpdateImages()
	if err != nil {
		return err
	}

	for _, repo := range d.Repos {
		name := fmt.Sprintf("%s/%s:%s", d.opts.org, repo.Name, d.opts.releasever)
		if dh.ImageExists(name) {
			continue
		}
		fmt.Println("  " + name)
		_, err := dh.BuildImage(name, repo.gitRepo.DeployPath())
		if err != nil {
			return err
		}
	}

	return nil
}

// UploadImages uploads all images defined on the definition
func (d *Definition) UploadImages() error {
	dh, err := dockerhost.New(d.opts.host)
	if err != nil {
		return err
	}
	dh.SetAuthCredentials(d.opts.username, d.opts.password)

	for _, repo := range d.Repos {
		name := fmt.Sprintf("%s/%s:%s", d.opts.org, repo.Name, d.opts.releasever)
		fmt.Println("  " + name)

		_, err := dh.PushImage(name)
		if err != nil {
			return err
		}
	}

	return nil
}

// GenerateOutput creates a file from the definition and template.yml
func (d *Definition) GenerateOutput() error {
	tpl, err := loadTemplate(d.opts.template)
	if err != nil {
		return err
	}

	dh, err := dockerhost.New(d.opts.host)
	if err != nil {
		return err
	}

	dh.UpdateImages()

	for _, repo := range d.Repos {
		var image string

		commit, cerr := repo.gitRepo.CommitID()
		if cerr != nil {
			return err
		}

		if d.opts.isRelease {
			image = fmt.Sprintf("%s/%s:%s", d.opts.org, repo.Name, d.opts.releasever)
		} else if d.opts.devmode && repo.gitRepo.HasChanges() {
			t := time.Now()
			year, month, day := t.Date()
			image = fmt.Sprintf("%s:%s-%d%d%d-%d%d%d", repo.Name, commit, year, month, day, t.Hour(), t.Minute(), t.Second())
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

	err = tpl.WriteFile(d.opts.output)
	if err != nil {
		return err
	}

	return nil
}
