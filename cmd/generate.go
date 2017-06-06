/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/r3labs/composable/build"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate a docker compose file",
	Long:  `generates a docker compose yaml file`,
	Run:   build.Generate,
}

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("overrides", "b", "", "Overides branch for specified repos")
	// excludes
	generateCmd.Flags().StringP("global-branch", "G", "", "Overides the branch for all repos, excluding ones specified in overides")
	generateCmd.Flags().StringP("environment", "E", "", "Sets an environmental variable for all docker containers")
}
