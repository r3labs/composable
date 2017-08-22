/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/r3labs/composable/build"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "destroys a compose environment",
	Long:  `destroys a compose environment`,
	Run:   build.Down,
}

func init() {
	RootCmd.AddCommand(downCmd)
	downCmd.Flags().BoolP("clean", "R", false, "Cleans all containers and images on a successful down")
}
