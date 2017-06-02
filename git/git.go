/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package git

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

// Repo stores all information about a git repo
type Repo struct {
	Repo        string
	Destination string
}

// Clone sets up and clones a git repo
func Clone(repo, destination string) (*Repo, error) {
	r := Repo{
		Repo:        repo,
		Destination: destination,
	}

	err := r.clone()
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Name returns the repo's name
func (r *Repo) Name() string {
	path := strings.Split(r.Repo, "/")
	return strings.Replace(path[len(path)-1], ".git", "", -1)
}

// Exists checks if the repo exists in the destination
func (r *Repo) Exists() bool {
	_, err := os.Stat(r.DeployPath())
	if err != nil {
		return false
	}

	return true
}

// DeployPath gives the full path to the project/repo
func (r *Repo) DeployPath() string {
	return path.Join(r.Destination, r.Name())
}

// Fetch all branches from remote
func (r *Repo) Fetch() error {
	cmd := exec.Command("git", "fetch")
	cmd.Dir = r.DeployPath()

	_, err := cmd.Output()
	if err != nil {
		return errors.New("could not fetch repo data: " + r.Name())
	}

	return nil
}

// Checkout git branch
func (r *Repo) Checkout(branch string) error {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = r.DeployPath()

	_, err := cmd.Output()
	if err != nil {
		return errors.New("could not checkout branch " + r.Name() + ":" + branch)
	}

	return nil
}

// Pull from remote
func (r *Repo) Pull() error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = r.DeployPath()

	_, err := cmd.Output()
	if err != nil {
		return errors.New("could not pull repo changes: " + r.Name())
	}

	return nil
}

// CommitID returns the commit id for the currently checked out branch
func (r *Repo) CommitID() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = r.DeployPath()

	output, err := cmd.Output()
	if err != nil {
		return "", errors.New("could not get git revision id")
	}

	id := string(output)
	return strings.TrimSpace(id), nil
}

// HasChanges returns true if there are local changes to a repo
func (r *Repo) HasChanges() bool {
	cmd := exec.Command("git", "status")
	cmd.Dir = r.DeployPath()

	output, _ := cmd.Output()

	return !strings.Contains(string(output), "nothing to commit, working tree clean")
}

// Sync the repo
func (r *Repo) Sync(branch string) error {
	// Fetch correct branch and update
	err := r.Fetch()
	if err != nil {
		return err
	}

	err = r.Checkout(branch)
	if err != nil {
		return fmt.Errorf("could not checkout repo branch " + r.Name() + ":" + branch)
	}

	err = r.Pull()
	return err
}

// Clone the repositort into the destination
func (r *Repo) clone() error {
	// Clone the repo, if it doesn't exist
	if !r.Exists() {
		cmd := exec.Command("git", "clone", r.Repo)
		cmd.Dir = r.Destination

		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("could not clone repo %s", r.Name())
		}
	}
	return nil
}
