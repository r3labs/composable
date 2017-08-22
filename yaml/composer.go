/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Composer file
type Composer struct {
	Version  string          `yaml:"version"`
	Services map[string]Repo `yaml:"services"`
	Networks yaml.MapSlice   `yaml:"networks"`
}

// Load the composer template
func LoadTemplate(path string) (*Composer, error) {
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
