/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"os"

	"github.com/r3labs/composable/docker/client"
	"github.com/r3labs/composable/docker/compose"
	"github.com/r3labs/composable/yaml"
	"github.com/spf13/cobra"
)

func Up(cmd *cobra.Command, args []string) {
	var services []string

	composeEnv, _ := cmd.Flags().GetString("compose-env")
	composeFile, _ := cmd.Flags().GetString("compose-file")
	dockerHost, _ := cmd.Flags().GetString("docker-host")

	_, err := os.Stat(composeFile)
	if err != nil {
		fatal(errors.New("could not locate docker-compose.yml"))
	}

	dc, err := yaml.LoadTemplate(composeFile)
	if err != nil {
		fatal(err)
	}

	c, err := compose.New(composeEnv, composeFile)
	if err != nil {
		fatal(err)
	}

	cli, err := client.New(dockerHost)
	if err != nil {
		fatal(err)
	}

	for k, v := range dc.Services {
		exists, err := cli.HasImage(v.Image())
		if err != nil {
			fatal(err)
		}

		if v.BuildPath() != "" && !exists {
			services = append(services, k)
		}
	}

	err = c.Build(services, false)
	if err != nil {
		fatal(err)
	}

	err = c.Up()
	if err != nil {
		fatal(err)
	}
}
