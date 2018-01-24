/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/r3labs/composable/build"
	"github.com/spf13/cobra"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "run a release of the project",
	Long: `run a release of the target project.
	This will build and push all images to docker hub.
	On completion a docker-compose yaml file will be generated.`,
	Run: build.Release,
}

func init() {
	RootCmd.AddCommand(releaseCmd)
	releaseCmd.Flags().StringP("version", "v", "", "Release version")
	releaseCmd.Flags().StringP("global-branch", "G", "", "Overides the branch for all repos. Used when producing a pre-release build")
}
