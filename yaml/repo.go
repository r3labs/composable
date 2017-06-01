package yaml

import "github.com/r3labs/composable/safe"

// Repo definition
type Repo map[string]interface{}

/*
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
*/

func (r *Repo) Name() string {
	return safe.String((*r)["name"])
}

func (r *Repo) Branch() string {
	return safe.String((*r)["branch"])
}

func (r *Repo) SetEnv(k, v string) {
	if (*r)["environment"] == nil {
		(*r)["environment"] = make(map[string]string)
	}
	(*r)["environment"].(map[string]string)[k] = v
}

func (r *Repo) SetBranch(branch string) {
	(*r)["branch"] = branch
}
