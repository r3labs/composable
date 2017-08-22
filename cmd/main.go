/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/r3labs/composable/config"
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
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cpath := path.Join(home, ".composable.yaml")
	config.CreateConfig(cpath)

	viper.AddConfigPath(home)
	viper.SetConfigName(".composable")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetDefault("build.compose-env", "composable")
	viper.SetDefault("build.compose-file", "./docker-compose.yml")
	viper.SetDefault("build.path", "./composable")
	viper.SetDefault("docker.host", "unix:///var/run/docker.sock")

	RootCmd.PersistentFlags().StringP("compose-env", "e", viper.GetString("build.compose-env"), "Name of the compose environment")
	RootCmd.PersistentFlags().StringP("compose-file", "o", viper.GetString("build.compose-file"), "Path to docker-compose.yml")
	RootCmd.PersistentFlags().StringP("build-path", "P", viper.GetString("build.path"), "Path to deploy git repos")
	RootCmd.PersistentFlags().StringP("docker-org", "O", viper.GetString("docker.org"), "Docker organisation used for releases")
	RootCmd.PersistentFlags().StringP("docker-user", "U", viper.GetString("docker.user"), "Docker user used for releases")
	RootCmd.PersistentFlags().StringP("docker-host", "H", viper.GetString("docker.host"), "Docker host used for builds")

	viper.BindPFlag("build.compose-env", RootCmd.PersistentFlags().Lookup("compose-env"))
	viper.BindPFlag("build.compose-file", RootCmd.PersistentFlags().Lookup("compose-file"))
	viper.BindPFlag("build.path", RootCmd.PersistentFlags().Lookup("build-path"))
}
