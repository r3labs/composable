/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/r3labs/composable/docker/compose"
	"github.com/spf13/cobra"
)

func Scale(cmd *cobra.Command, services []string) {
	if len(services) < 1 {
		fatal(errors.New("scale must specify service(s)"))
	}

	opts := make(map[string]int)

	for _, s := range services {
		x := strings.Split(s, "=")
		if len(x) < 2 {
			fatal(errors.New("scale arguments must take the form of 'SERVICE=COUNT', i.e. 'web=2'"))
		}

		opts[x[0]], _ = strconv.Atoi(x[1])
	}

	composeEnv, _ := cmd.Flags().GetString("compose-env")
	composeFile, _ := cmd.Flags().GetString("compose-file")

	_, err := os.Stat(composeFile)
	if err != nil {
		fatal(errors.New("could not locate docker-compose.yml"))
	}

	c, err := compose.New(composeEnv, composeFile)
	if err != nil {
		fatal(err)
	}

	err = c.Scale(opts)
	if err != nil {
		fatal(err)
	}
}
