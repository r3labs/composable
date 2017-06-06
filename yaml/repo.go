package yaml

import "github.com/r3labs/composable/safe"

// Repo definition
type Repo map[string]interface{}

func (r *Repo) Name() string {
	return safe.String((*r)["name"])
}

func (r *Repo) Branch() string {
	return safe.String((*r)["branch"])
}

func (r *Repo) BuildPath() string {
	return safe.String((*r)["build"])
}

func (r *Repo) URL() string {
	return safe.String((*r)["path"])
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
