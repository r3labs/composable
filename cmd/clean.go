/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/r3labs/composable/build"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean environment",
	Long:  `clean all images and volumes associated with a compose envrionment`,
	Run:   build.Clean,
}

func init() {
	RootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolP("force", "f", true, "Forces cleanup")
}
