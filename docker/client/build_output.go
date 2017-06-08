/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import "encoding/json"

// BuildOutput : represents build output from a build stream
type BuildOutput struct {
	Stream string `json:"stream"`
}

// ProcessOutput : gets output values from a build stream
func ProcessOutput(data []byte) *BuildOutput {
	var o BuildOutput
	json.Unmarshal(data, &o)
	return &o
}
