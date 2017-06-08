/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Tar ...
func Tar(buildDir string) (io.Reader, error) {
	files := make(map[string]os.FileInfo)

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	err := filepath.Walk(buildDir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		files[path] = f
		return nil
	})
	if err != nil {
		return nil, err
	}

	for path, file := range files {
		hdr := &tar.Header{
			Name: file.Name(),
			Mode: int64(file.Mode()),
			Size: file.Size(),
		}

		err = tw.WriteHeader(hdr)
		if err != nil {
			return nil, err
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		_, err = tw.Write(data)
		if err != nil {
			return nil, err
		}
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return bytes.NewReader(buf.Bytes()), nil
}
