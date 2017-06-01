/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package options

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Options stores this applications configuration and options
type Options struct {
	Build struct {
		Path       string
		Output     string
		Definition string
		Template   string
		GlobalEnv  string
		Overrides  map[string]string
		Excludes   []string
	}
	Git struct {
		MaxWorkers   int
		GlobalBranch string
	}
	Release struct {
		IsRelease bool
		Version   string
	}
	Docker struct {
		Org      string
		Host     string
		Username string
		Password string
	}
}

// GetOptions reads options from cli arguments
func GetOptions() (string, Options) {
	var opts Options
	var overrides string
	var excludes string

	mode := os.Args[1]

	if len(os.Args) > 2 && mode != "--help" {
		// Clear mode for parsing flags
		os.Args = append(os.Args[:1], os.Args[1+1:]...)
	}

	flag.StringVar(&opts.Docker.Host, "h", "unix:///var/run/docker.sock", "Docker host to target. Defaults to current host")
	flag.StringVar(&opts.Build.Path, "d", "/tmp/composable/", "Deployment directory where all repos are checked out")
	flag.StringVar(&opts.Build.Output, "o", "docker-compose.yml", "Output file for docker-compose")
	flag.StringVar(&overrides, "b", "", "Override a repo's branch, specified by repo name, comma delimited")
	flag.StringVar(&excludes, "exclude", "", "Ignore repos from the definition based on a matching value")
	flag.StringVar(&opts.Git.GlobalBranch, "G", "", "Globally override all git branches")
	flag.StringVar(&opts.Build.GlobalEnv, "E", "", "Globally add extra environment options")
	flag.StringVar(&opts.Docker.Org, "org", "", "Docker hub organisation target for release")
	flag.StringVar(&opts.Release.Version, "version", "", "Version to release")
	flag.IntVar(&opts.Git.MaxWorkers, "w", runtime.NumCPU(), "number of build workers for a release, defaults to number of cpu's")
	flag.StringVar(&opts.Docker.Username, "u", "", "docker hub username, used only for release")
	flag.StringVar(&opts.Docker.Password, "p", "", "docker hub password, used only for release")
	flag.Parse()

	opts.Build.Definition = flag.Arg(0)
	opts.Build.Template = flag.Arg(1)
	opts.Build.Overrides = GetOverrides(overrides)
	opts.Build.Excludes = strings.Split(excludes, ",")

	switch mode {
	case "rel", "release":
		opts.Release.IsRelease = true
		if opts.Release.Version == "" {
			panic("No release version specified!")
		}
		if opts.Docker.Org == "" {
			panic("No org specified!")
		}
	}

	_, err := os.Stat(opts.Build.Path)
	if err != nil {
		panic(fmt.Sprintf("Deployment directory '%s' does not exist!", opts.Build.Path))
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
