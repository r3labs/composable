/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "composable",
	Short: "Composable is a tool for working with building and developing on docker, based off of libcompose",
	Long:  `Composable is a tool for working with building and developing on docker, based off of libcompose`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	//cobra.OnInitialize(initConfig)
}

func main() {
	os.Exit(1)
}
