/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/r3labs/composable/cmd"
	"github.com/r3labs/composable/config"
	"github.com/spf13/viper"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := config.WriteConfig(viper.ConfigFileUsed(), viper.AllSettings()); err != nil {
		log.Println(err)
	}
}
