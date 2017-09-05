/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"fmt"
	"os"

	"github.com/r3labs/composable/docker/client"
	"github.com/r3labs/composable/docker/compose"
	"github.com/r3labs/composable/yaml"
	"github.com/spf13/cobra"
)

func Clean(cmd *cobra.Command, args []string) {
	composeEnv, _ := cmd.Flags().GetString("compose-env")
	composeFile, _ := cmd.Flags().GetString("compose-file")
	dockerHost, _ := cmd.Flags().GetString("docker-host")
	dockerRegistry, _ := cmd.Flags().GetString("docker-registry")
	force, _ := cmd.Flags().GetBool("force")

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

	cli, err := client.New(dockerHost, dockerRegistry)
	if err != nil {
		fatal(err)
	}

	err = c.Down(true)
	if err != nil && !force {
		fatal(err)
	}

	for _, service := range dc.Services {
		fmt.Println("removing image: " + service.Image())
		err := cli.DeleteImage(service.Image())
		if err != nil {
			fatal(err)
		}
	}
}
