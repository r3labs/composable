/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd ...
var RootCmd = &cobra.Command{
	Use:   "composable",
	Short: "Composable is a tool for developing on docker, based off of libcompose",
	Long:  `Composable is a tool for developing on docker, based off of libcompose`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	//cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringP("environment", "E", "/tmp/", "Path to the environment")
	RootCmd.PersistentFlags().StringP("build-path", "f", "/tmp/composable", "Path to deploy git repos")
	RootCmd.PersistentFlags().StringP("docker-org", "O", "", "Docker organisation used for releases")
	RootCmd.PersistentFlags().StringP("docker-user", "U", "", "Docker user used for releases")
	viper.BindPFlag("environment", RootCmd.PersistentFlags().Lookup("environment"))
	viper.BindPFlag("build-path", RootCmd.PersistentFlags().Lookup("build-path"))
	viper.BindPFlag("docker-org", RootCmd.PersistentFlags().Lookup("docker-org"))
	viper.BindPFlag("docker-user", RootCmd.PersistentFlags().Lookup("docker-user"))
	viper.SetDefault("environment", "/tmp/")
	viper.SetDefault("build-path", "/tmp/composable")
}
