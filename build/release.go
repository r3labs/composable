/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"fmt"

	"github.com/r3labs/composable/docker/client"
	"github.com/r3labs/composable/yaml"
	"github.com/spf13/cobra"
)

func Release(cmd *cobra.Command, args []string) {
	/*
		environment, _ := cmd.Flags().GetString("environment")
		overrides, _ := cmd.Flags().GetString("overrides")
		excludes, _ := cmd.Flags().GetString("excludes")
		global, _ := cmd.Flags().GetString("global-branch")
		output, _ := cmd.Flags().GetString("compose-file")
		buildpath, _ := cmd.Flags().GetString("build-path")

		if buildpath == "" {
			fatal(errors.New("no build path specified"))
		}

		if len(args) < 2 {
			fatal(errors.New("generate must specify a definition and template file"))
		}

		d, err := yaml.LoadDefinition(args[0])
		if err != nil {
			fatal(err)
		}

		d.Template = args[1]
		d.BuildPath = buildpath
		d.Environment(environment)
		d.Overrides(overrides, excludes, global)

		fmt.Println("cloning repos:")
		err = CloneRepos(d, buildpath)
		if err != nil {
			fatal(err)
		}

		fmt.Println("generating output")
		err = d.GenerateOutput(output)
		if err != nil {
			fatal(err)
		}
	*/
	buildPath, _ := cmd.Flags().GetString("build-path")
	dockerOrg, _ := cmd.Flags().GetString("docker-org")
	//dockerUser, _ := cmd.Flags().GetString("docker-user")
	dockerHost, _ := cmd.Flags().GetString("docker-host")
	version, _ := cmd.Flags().GetString("version")

	if buildPath == "" {
		fatal(errors.New("no build path specified"))
	}

	if len(args) < 2 {
		fatal(errors.New("release must specify a definition and template file"))
	}

	cli, err := client.New(dockerHost)
	if err != nil {
		fatal(err)
	}

	/*
		pwd := config.GetPassword("please enter your docker registry password")
		err = cli.Login(dockerUser, pwd)
		if err != nil {
			fatal(err)
		}
	*/

	d, err := yaml.LoadDefinition(args[0])
	if err != nil {
		fatal(err)
	}

	d.BuildPath = buildPath
	d.Release.Org = dockerOrg
	d.Release.Version = version
	d.Template = args[1]

	fmt.Println("cloning repos:")
	err = CloneRepos(d, buildPath)
	if err != nil {
		fatal(err)
	}

	err = d.ParseRepos()
	if err != nil {
		fatal(err)
	}

	for _, s := range d.Repos {
		fmt.Println(s.Image())
		fmt.Println(s.BuildPath())
		err = cli.BuildImage(s.Image(), s.BuildPath())
		if err != nil {
			fatal(err)
		}
	}

	fmt.Println("generating output")
	err = d.GenerateOutput("docker-compose.yml")
	if err != nil {
		fatal(err)
	}
}
