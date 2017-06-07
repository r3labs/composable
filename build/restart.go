/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"os"

	"github.com/r3labs/composable/docker/compose"
	"github.com/r3labs/composable/yaml"
	"github.com/spf13/cobra"
)

func Restart(cmd *cobra.Command, services []string) {
	composeEnv, _ := cmd.Flags().GetString("compose-env")
	composeFile, _ := cmd.Flags().GetString("compose-file")

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

	if len(services) < 1 {
		for s := range dc.Services {
			services = append(services, s)
		}
	}

	err = c.Stop(services)
	if err != nil {
		fatal(err)
	}

	err = c.Start(services)
	if err != nil {
		fatal(err)
	}
}
