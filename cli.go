/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	"github.com/r3labs/composable/dockerhost"
)

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func login(host string) (string, string) {
	var username string
	var password string

	fmt.Println("Please enter your docker hub credentials")
	fmt.Printf("Username: ")
	fmt.Scanf("%s", &username)

	fmt.Printf("Password: ")
	pass, _ := gopass.GetPasswdMasked()
	password = string(pass)

	dh, err := dockerhost.New(host)
	if err != nil {
		panic(err)
	}

	err = dh.SetAuthCredentials(username, password)
	if err != nil {
		panic(err)
	}

	return username, password
}

func generate(opts *Options) {
	fmt.Println("Generating Output Definition")

	def, err := LoadDefiniton(opts.definition, opts)
	if err != nil {
		exit(err)
	}

	fmt.Println(" cloning repos")
	err = def.CloneRepos()
	if err != nil {
		exit(err)
	}

	fmt.Println(" generating output definition")
	err = def.GenerateOutput()
	if err != nil {
		exit(err)
	}
}

func release(opts *Options) {
	fmt.Printf("Releasing Version %s\n", opts.releasever)

	if opts.username == "" && opts.password == "" {
		opts.username, opts.password = login(opts.host)
	}

	def, err := LoadDefiniton(opts.definition, opts)
	if err != nil {
		exit(err)
	}

	fmt.Println(" cloning repos")
	err = def.CloneRepos()
	if err != nil {
		exit(err)
	}

	fmt.Println(" building docker images")
	err = def.BuildImages()
	if err != nil {
		exit(err)
	}

	fmt.Println(" publishing docker images")
	err = def.UploadImages()
	if err != nil {
		exit(err)
	}

	fmt.Println(" generating output definition")
	err = def.GenerateOutput()
	if err != nil {
		exit(err)
	}
}
