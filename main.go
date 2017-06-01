/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/r3labs/composable/cmd"
	"github.com/spf13/viper"
)

func main() {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(home)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".composable.yml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(viper.ConfigFileUsed())
		err = ioutil.WriteFile(viper.ConfigFileUsed(), []byte{}, 0644)
		if err != nil {
			os.Exit(1)
		}

		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	/*
		mode, opts := GetOptions()

		switch mode {
		case "gen", "generate":
			generate(&opts)
		case "rel", "release":
			release(&opts)
		case "destroy":
			fmt.Println("Destroying")
		}
	*/
}
