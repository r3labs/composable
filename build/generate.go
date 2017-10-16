/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"fmt"

	"github.com/r3labs/composable/yaml"
	"github.com/spf13/cobra"
)

func Generate(cmd *cobra.Command, args []string) {
	environment, _ := cmd.Flags().GetString("environment")
	overrides, _ := cmd.Flags().GetString("overrides")
	excludes, _ := cmd.Flags().GetString("excludes")
	global, _ := cmd.Flags().GetString("global-branch")
	output, _ := cmd.Flags().GetString("compose-file")
	buildpath, _ := cmd.Flags().GetString("build-path")
	edition, _ := cmd.Flags().GetString("edition")

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

	if edition == "community" {
		for _, repo := range d.Repos {
			if repo["edition"] == "enterprise" {
				d.ExcludeRepo(repo.Name())
			}
		}
	}

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
}
