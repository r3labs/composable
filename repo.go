/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Definition of repos
type Definition struct {
	Repos []*Repo `yaml:"repos"`
}

// Repo definition
type Repo struct {
	Name         string   `yaml:"name"`
	Path         string   `yaml:"path"`
	Branch       string   `yaml:"branch"`
	Volumes      []string `yaml:"volumes"`
	Ports        []string `yaml:"ports"`
	Links        []string `yaml:"link"`
	Dependencies []string `yaml:"depends"`
	gitRepo      *GitRepo `yaml:"-"`
}

// Load the input definition
func loadDefiniton(path string) (*Definition, error) {
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
