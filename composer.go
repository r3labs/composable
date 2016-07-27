/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Composer file
type Composer struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Networks yaml.MapSlice      `yaml:"networks"`
}

// Service definition
type Service struct {
	Image        string            `yaml:"image,omitempty"`
	Build        string            `yaml:"build,omitempty"`
	Entrypoint   string            `yaml:"entrypoint,omitempty"`
	Restart      string            `yaml:"restart,omitempty"`
	Ports        []string          `yaml:"ports,omitempty"`
	Volumes      []string          `yaml:"volumes,omitempty"`
	Links        []string          `yaml:"links,omitempty"`
	Dependencies []string          `yaml:"depends_on,omitempty"`
	Environment  map[string]string `yaml:"environment,omitempty"`
}

// Load the composer template
func loadTemplate(path string) (*Composer, error) {
	var c Composer

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &c, err
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return &c, err
	}

	return &c, nil
}

// WriteFile outputs the docker compose file
func (s *Composer) WriteFile(path string) error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
