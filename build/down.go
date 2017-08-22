/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"os"

	"github.com/r3labs/composable/docker/compose"
	"github.com/spf13/cobra"
)

func Down(cmd *cobra.Command, args []string) {
	composeEnv, _ := cmd.Flags().GetString("compose-env")
	composeFile, _ := cmd.Flags().GetString("compose-file")
	clean, _ := cmd.Flags().GetBool("clean")

	_, err := os.Stat(composeFile)
	if err != nil {
		fatal(errors.New("could not locate docker-compose.yml"))
	}

	c, err := compose.New(composeEnv, composeFile)
	if err != nil {
		fatal(err)
	}

	err = c.Down(clean)
	if err != nil {
		fatal(err)
	}
}
