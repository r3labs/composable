/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package yaml

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/r3labs/composable/git"

	"gopkg.in/yaml.v2"
)

// Definition of repos
type Definition struct {
	Release struct {
		Version  string
		Org      string
		Registry string
	}
	Template  string
	BuildPath string
	Repos     []Repo `yaml:"repos"`
}

// LoadDefiniton the input definition
func LoadDefinition(path string) (*Definition, error) {
	var d Definition

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &d, err
	}

	err = yaml.Unmarshal(data, &d)
	if err != nil {
		return &d, err
	}

	return &d, nil
}

func (d *Definition) Environment(environment string) {
	if environment != "" {
		envs := strings.Split(environment, ",")
		for _, repo := range d.Repos {
			for _, env := range envs {
				e := strings.Split(env, "=")
				repo.SetEnv(e[0], e[1])
			}
		}
	}
}

func (d *Definition) Overrides(overrides, excludes, global string) {
	// Ommit/Exclude repos
	for _, repo := range strings.Split(excludes, ",") {
		d.ExcludeRepo(repo)
	}

	// Override branches
	if global != "" {
		for _, repo := range d.Repos {
			d.OverrideBranch(repo.Name(), global)
		}
	}

	for repo, branch := range GetOverrides(overrides) {
		d.OverrideBranch(repo, branch)
	}
}

// OverrideBranch updates a repo's branch
func (d *Definition) OverrideBranch(repo, branch string) {
	for i := 0; i < len(d.Repos); i++ {
		if d.Repos[i].Name() == repo {
			d.Repos[i].SetBranch(branch)
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
		if wildcard && strings.Contains(d.Repos[i].Name(), repo) || d.Repos[i].Name() == repo {
			d.Repos = append(d.Repos[:i], d.Repos[i+1:]...)
		}
	}
}

func GetOverrides(overrides string) map[string]string {
	o := make(map[string]string)

	if overrides != "" {
		for _, data := range strings.Split(overrides, ",") {
			x := strings.Split(data, ":")
			if len(data) > 1 {
				// name = repo branch
				o[x[0]] = x[1]
			}
		}
	}

	return o
}

// GenerateOutput creates a file from the definition and template.yml
func (d *Definition) GenerateOutput(output string) error {
	editions := []string{"community", "enterprise"}

	for _, edition := range editions {
		tpl, err := LoadTemplate(d.Template)
		if err != nil {
			return err
		}

		err = d.ParseRepos()
		if err != nil {
			return err
		}

		tpl.Version = "2"

		for _, repo := range d.Repos {
			if edition == "community" && repo["edition"] == "enterprise" {
				continue
			}

			tpl.Services[repo.Name()] = copyRepo(repo)

			// clean map of any unsupported field
			delete(tpl.Services[repo.Name()], "name")
			delete(tpl.Services[repo.Name()], "path")
			delete(tpl.Services[repo.Name()], "branch")
			delete(tpl.Services[repo.Name()], "edition")
			if d.Release.Version != "" {
				delete(tpl.Services[repo.Name()], "build")
			}
		}

		if edition == "enterprise" {
			output = strings.TrimSuffix(output, ".yml") + ".enterprise.yml"
		}

		err = tpl.WriteFile(output)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Definition) ParseRepos() error {
	for i := range d.Repos {
		var image string

		r := git.Repo{
			Repo:        d.Repos[i].Name(),
			Destination: d.BuildPath,
		}

		commit, err := r.CommitID()
		if err != nil {
			return err
		}

		switch {
		case d.Release.Registry != "https://index.docker.io/v1/" && d.Release.Version != "":
			image = fmt.Sprintf("%s/%s/%s:%s", d.Release.Registry, d.Release.Org, d.Repos[i].Name(), d.Release.Version)
		case d.Release.Version != "":
			image = fmt.Sprintf("%s/%s:%s", d.Release.Org, d.Repos[i].Name(), d.Release.Version)
		default:
			image = fmt.Sprintf("%s:%s", d.Repos[i].Name(), commit)
		}

		d.Repos[i]["image"] = image
		d.Repos[i]["build"] = r.DeployPath()
	}

	return nil
}

func copyRepo(r Repo) Repo {
	rx := make(Repo)

	for k, v := range r {
		rx[k] = v
	}

	return rx
}
