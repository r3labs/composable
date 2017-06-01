/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "kill a service",
	Long:  `kill a service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kill called")
	},
}

func init() {
	RootCmd.AddCommand(killCmd)
}
