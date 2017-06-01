/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

type Config struct {
	Environment    string `json:"environment"`
	BuildPath      string `json:"build_path"`
	DockerOrg      string `json:"docker_org"`
	DockerUsername string `json:"docker_username"`
}

func GetConfig() *Config {

}
