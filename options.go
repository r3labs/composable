/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"flag"
	"os"
)

// Options stores this applications configuration and options
type Options struct {
	home       string
	host       string
	output     string
	definition string
	template   string
}

// GetOptions reads options from cli arguments
func GetOptions() (string, Options) {
	opts := Options{}
	mode := os.Args[1]

	// Clear mode for parsing flags
	os.Args = append(os.Args[:1], os.Args[1+1:]...)

	flag.StringVar(&opts.host, "h", "unix:///var/run/docker.sock", "Docker host to target. Defaults to current host")
	flag.StringVar(&opts.home, "d", "/tmp/composable/", "Deployment directory where all repos are checked out")
	flag.StringVar(&opts.output, "o", "docker-compose.yml", "Output file for docker-compose")
	flag.Parse()

	opts.definition = flag.Arg(0)
	opts.template = flag.Arg(1)

	return mode, opts
}
