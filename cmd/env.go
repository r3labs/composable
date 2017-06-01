/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "set environment configuration",
	Long:  `set environment configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("env called")
	},
}

func init() {
	setCmd.AddCommand(envCmd)
}
