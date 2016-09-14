/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	mode, opts := GetOptions()

	switch mode {
	case "gen", "generate":
		fmt.Println("Generating Output Definition")

		def, err := cloneRepos(&opts)
		handleErr(err)

		err = generateOutputFile(&opts, def)
		handleErr(err)
	case "rel", "release":
		opts.username, opts.password = login()

		fmt.Println("")
		fmt.Printf("Releasing Version %s\n", opts.releasever)

		def, err := cloneRepos(&opts)
		handleErr(err)

		err = buildAndPush(&opts, def)
		handleErr(err)

		err = generateOutputFile(&opts, def)
		handleErr(err)
	case "destroy":
		fmt.Println("Destroying")
	}
}
