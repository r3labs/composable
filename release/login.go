/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package release

import (
	"fmt"

	"github.com/howeyc/gopass"
	"github.com/r3labs/composable/docker/host"
)

func Login(hostname string) (string, string) {
	var username string
	var password string

	fmt.Println("Please enter your docker hub credentials")
	fmt.Printf("Username: ")
	fmt.Scanf("%s", &username)

	fmt.Printf("Password: ")
	pass, _ := gopass.GetPasswdMasked()
	password = string(pass)

	dh, err := host.New(hostname)
	if err != nil {
		panic(err)
	}

	err = dh.SetAuthCredentials(username, password)
	if err != nil {
		panic(err)
	}

	return username, password
}
