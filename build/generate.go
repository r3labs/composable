/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"errors"
	"fmt"
	"os"

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

	for _, repo := range d.Repos {
		repo.SetEnv("ERNEST_EDITION", edition)
	}

	if edition == "community" {
		for i := len(d.Repos) - 1; i >= 0; i-- {
			if d.Repos[i]["edition"] == "enterprise" {
				d.ExcludeRepo(d.Repos[i].Name())
			}
		}
	}

	fmt.Println("cloning repos:")
	// Check if buildpath exists
	if _, err := os.Stat(buildpath); os.IsNotExist(err) {
		fatal(errors.New("Specified '" + buildpath + "' folder does not exist"))
	}
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
