/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// OutputStream stores data from output streams
type OutputStream []byte

func (o *OutputStream) Write(data []byte) (int, error) {
	*o = append(*o, data...)
	return len(data), nil
}

// Output returns the
func (o *OutputStream) Output() string {
	return string(*o)
}
