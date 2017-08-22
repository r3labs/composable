/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	yaml "gopkg.in/yaml.v2"
)

func CreateConfig(cpath string) error {
	_, err := os.Stat(cpath)
	if os.IsNotExist(err) {
		err = ioutil.WriteFile(cpath, []byte{}, 0644)
		if err != nil {
			os.Exit(1)
		}
	}

	return nil
}

func WriteConfig(cpath string, config map[string]interface{}) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(cpath, data, 0644)
}

func GetPassword(msg string) string {
	fmt.Printf(msg + ": ")
	pass, _ := gopass.GetPasswdMasked()
	password := string(pass)
	return password
}

func Set(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("set must contain arguments, i.e. 'build.env composable-test'")
	}

	val, err := strconv.Atoi(args[1])
	if err != nil {
		viper.Set(args[0], args[1])
	} else {
		viper.Set(args[0], val)
	}
}
