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
		generate(&opts)
	case "rel", "release":
		release(&opts)
	case "destroy":
		fmt.Println("Destroying")
	}
}
