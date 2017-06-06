/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"fmt"
	"os"

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

	fmt.Println(composeEnv)
	fmt.Println(dockerHost)

	for k, v := range dc.Services {
		if v.BuildPath() != "" {
			services = append(services, k)
		}
	}

	if c.Project != nil {

	}

	/*

		err = c.Build(services, false)
		if err != nil {
			fatal(err)
		}

		err = c.Up()
		if err != nil {
			fatal(err)
		}
	*/
}
