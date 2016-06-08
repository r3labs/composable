/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// GitRepo stores all information about a git repo
type GitRepo struct {
	Repo           string
	Destination    string
	deploymentPath string
}

// NewGitRepo sets up a git repo
func NewGitRepo(repo, destination string) *GitRepo {
	return &GitRepo{
		Repo:        repo,
		Destination: destination,
	}
}

// Name returns the repo's name
func (g *GitRepo) Name() string {
	path := strings.Split(g.Repo, "/")
	return strings.Replace(path[len(path)-1], ".git", "", -1)
}

// Exists checks if the repo exists in the destination
func (g *GitRepo) Exists() bool {
	_, err := os.Stat(g.deploymentPath)
	if err != nil {
		return false
	}
	return true
}

// Clone the repositort into the destination
func (g *GitRepo) Clone() error {
	g.deploymentPath = g.Destination + g.Name()

	// Clone the repo, if it doesn't exist
	if !g.Exists() {
		cmd := exec.Command("git", "clone", g.Repo)
		cmd.Dir = g.Destination
		_, err := cmd.Output()
		if err != nil {
			return errors.New("Could not clone repo")
		}
	}
	return nil
}

// Fetch all branches from remote
func (g *GitRepo) Fetch() error {
	cmd := exec.Command("git", "fetch")
	cmd.Dir = g.deploymentPath
	_, err := cmd.Output()
	if err != nil {
		return errors.New("Could not fetch repo data")
	}
	return nil
}

// Checkout git branch
func (g *GitRepo) Checkout(branch string) error {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = g.deploymentPath
	_, err := cmd.Output()
	if err != nil {
		return errors.New("Could not checkout branch")
	}
	return nil
}

// Pull from remote
func (g *GitRepo) Pull() error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = g.deploymentPath
	_, err := cmd.Output()
	if err != nil {
		return errors.New("Could not pull repo changes")
	}

	// Update the commit id

	return nil
}

// CommitID returns the commit id for the currently checked out branch
func (g *GitRepo) CommitID() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = g.deploymentPath
	output, err := cmd.Output()
	if err != nil {
		return "", errors.New("Could not get git revision id")
	}

	id := string(output)
	return strings.TrimSpace(id), nil
}
