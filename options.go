/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"flag"
	"os"
	"strings"
)

// Options stores this applications configuration and options
type Options struct {
	home       string
	host       string
	output     string
	definition string
	template   string
	overrides  map[string]string
}

// GetOptions reads options from cli arguments
func GetOptions() (string, Options) {
	var overrides string

	opts := Options{}
	mode := os.Args[1]

	// Clear mode for parsing flags
	os.Args = append(os.Args[:1], os.Args[1+1:]...)

	flag.StringVar(&opts.host, "h", "unix:///var/run/docker.sock", "Docker host to target. Defaults to current host")
	flag.StringVar(&opts.home, "d", "/tmp/composable/", "Deployment directory where all repos are checked out")
	flag.StringVar(&opts.output, "o", "docker-compose.yml", "Output file for docker-compose")
	flag.StringVar(&overrides, "b", "", "Override a repo's branch, specified by repo name, comma delimited")
	flag.Parse()

	opts.definition = flag.Arg(0)
	opts.template = flag.Arg(1)
	opts.overrides = GetOverrides(overrides)

	return mode, opts
}

// GetOverrides based on cli arguments
func GetOverrides(overrides string) map[string]string {
	o := make(map[string]string)
	if overrides != "" {
		for _, data := range strings.Split(overrides, ",") {
			x := strings.Split(data, ":")
			if len(data) > 1 {
				// name = repo branch
				o[x[0]] = x[1]
			}
		}
	}

	return o
}
