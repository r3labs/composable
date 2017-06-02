/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func createConfig(cpath string) error {
	_, err := os.Stat(cpath)
	if os.IsNotExist(err) {
		err = ioutil.WriteFile(cpath, []byte{}, 0644)
		if err != nil {
			os.Exit(1)
		}
	}

	return nil
}

func writeConfig(cpath string, config map[string]interface{}) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(cpath, data, 0644)
}
