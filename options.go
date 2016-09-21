/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Options stores this applications configuration and options
type Options struct {
	home         string
	host         string
	output       string
	definition   string
	template     string
	releasever   string
	org          string
	maxworkers   int
	username     string
	password     string
	globalbranch string
	isRelease    bool
	overrides    map[string]string
}

// GetOptions reads options from cli arguments
func GetOptions() (string, Options) {
	var overrides string

	opts := Options{}
	mode := os.Args[1]

	if len(os.Args) > 2 && mode != "--help" {
		// Clear mode for parsing flags
		os.Args = append(os.Args[:1], os.Args[1+1:]...)
	}

	flag.StringVar(&opts.host, "h", "unix:///var/run/docker.sock", "Docker host to target. Defaults to current host")
	flag.StringVar(&opts.home, "d", "/tmp/composable/", "Deployment directory where all repos are checked out")
	flag.StringVar(&opts.output, "o", "docker-compose.yml", "Output file for docker-compose")
	flag.StringVar(&overrides, "b", "", "Override a repo's branch, specified by repo name, comma delimited")
	flag.StringVar(&opts.globalbranch, "G", "", "Globally override all git branches")
	flag.StringVar(&opts.org, "org", "", "Docker hub organisation target for release")
	flag.StringVar(&opts.releasever, "version", "", "Version to release")
	flag.IntVar(&opts.maxworkers, "w", runtime.NumCPU(), "number of build workers for a release, defaults to number of cpu's")
	flag.Parse()

	opts.definition = flag.Arg(0)
	opts.template = flag.Arg(1)
	opts.overrides = GetOverrides(overrides)

	switch mode {
	case "rel", "release":
		opts.isRelease = true
		if opts.releasever == "" {
			panic("No release version specified!")
		}
		if opts.org == "" {
			panic("No org specified!")
		}
	}

	_, err := os.Stat(opts.home)
	if err != nil {
		panic(fmt.Sprintf("Deployment directory '%s' does not exist!", opts.home))
	}

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
