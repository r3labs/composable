/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "scale a service",
	Long:  `scale a service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scale called")
	},
}

func init() {
	RootCmd.AddCommand(scaleCmd)
}
