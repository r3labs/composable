/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "docker configuration",
	Long:  `docker configuration options used for a release or build`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker called")
	},
}

func init() {
	setCmd.AddCommand(dockerCmd)
}
