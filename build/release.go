/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/r3labs/composable/config"
	"github.com/r3labs/composable/docker/client"
	"github.com/r3labs/composable/yaml"
	"github.com/spf13/cobra"
)

func Release(cmd *cobra.Command, args []string) {
	buildPath, _ := cmd.Flags().GetString("build-path")
	dockerOrg, _ := cmd.Flags().GetString("docker-org")
	dockerUser, _ := cmd.Flags().GetString("docker-user")
	dockerPassword, _ := cmd.Flags().GetString("docker-password")
	dockerHost, _ := cmd.Flags().GetString("docker-host")
	dockerRegistry, _ := cmd.Flags().GetString("docker-registry")
	version, _ := cmd.Flags().GetString("version")

	if buildPath == "" {
		fatal(errors.New("no build path specified"))
	}

	if len(args) < 2 {
		fatal(errors.New("release must specify a definition and template file"))
	}

	cli, err := client.New(dockerHost, dockerRegistry)
	if err != nil {
		fatal(err)
	}

	if dockerPassword == "" {
		dockerPassword = config.GetPassword("please enter your docker registry password")
	}

	err = cli.Login(dockerUser, dockerPassword)
	if err != nil {
		fatal(err)
	}

	d, err := yaml.LoadDefinition(args[0])
	if err != nil {
		fatal(err)
	}

	d.BuildPath = buildPath
	d.Release.Org = dockerOrg
	d.Release.Version = version
	d.Release.Registry = dockerRegistry
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

	fmt.Println("building images:")
	for _, s := range d.Repos {
		berr := build(cli, s)
		if berr != nil {
			fatal(berr)
		}
	}

	fmt.Println("pushing images:")
	for _, s := range d.Repos {
		perr := push(cli, s)
		if perr != nil {
			fatal(perr)
		}
	}

	fmt.Println("generating output")
	err = d.GenerateOutput("docker-compose.yml")
	if err != nil {
		fatal(err)
	}
}

func build(cli *client.Client, s yaml.Repo) error {
	fmt.Printf("building image: %s\n", s.Image())
	output, err := cli.BuildImage(s.Image(), s.BuildPath())
	if err != nil {
		fatal(err)
	}
	rd := bufio.NewReader(output)
	for {
		data, err := rd.ReadBytes(10)
		if err != nil {
			fmt.Println("")
			break
		}
		output := client.ProcessOutput(data)
		fmt.Print(string(output.Stream))
	}

	return nil
}

func push(cli *client.Client, s yaml.Repo) error {
	fmt.Printf("pushing image: %s\n", s.Image())
	output, err := cli.PushImage(s.Image())
	if err != nil {
		fatal(err)
	}
	rd := bufio.NewReader(output)
	for {
		data, err := rd.ReadBytes(10)
		if err != nil {
			fmt.Println("")
			break
		}
		output := client.ProcessOutput(data)
		fmt.Print(string(output.Stream))
	}

	return nil
}
