/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package build

import (
	"fmt"
	"sync"

	"github.com/r3labs/composable/git"
	"github.com/r3labs/composable/yaml"
)

// CloneRepos clones and checks out the correct branch for a repo
func CloneRepos(d *yaml.Definition, bpath string) error {
	var wg sync.WaitGroup
	wg.Add(len(d.Repos))

	for i := 0; i < len(d.Repos); i++ {
		go func(wg *sync.WaitGroup, d *yaml.Definition, i int) {
			defer wg.Done()

			fmt.Printf("  %s\n", d.Repos[i].Name())
			r, err := git.Clone(d.Repos[i].URL(), bpath)
			if err != nil {
				panic(err)
			}

			err = r.Sync(d.Repos[i].Branch())
			if err != nil {
				panic(err)
			}

			//d.Repos[i].gitRepo = r
		}(&wg, d, i)
	}

	wg.Wait()

	return nil
}
