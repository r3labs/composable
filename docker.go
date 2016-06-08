/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "github.com/fsouza/go-dockerclient"

// DockerHost holds the docker client we need to talk to docker
type DockerHost struct {
	Host   string
	Images []string
	client *docker.Client
}

// NewDockerHost returns a new docker host
func NewDockerHost(host string) (*DockerHost, error) {
	d := DockerHost{Host: host}

	client, err := docker.NewClient(host)
	if err != nil {
		return &d, err
	}
	d.client = client

	return &d, nil
}

// UpdateImages all of the docker images on the host
func (d *DockerHost) UpdateImages() error {
	images, err := d.client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		return err
	}

	for _, img := range images {
		if len(img.RepoTags) > 0 {
			d.Images = append(d.Images, img.RepoTags[0])
		}
	}

	return nil
}

// ImageExists returns true if an image is present on the docker host
func (d *DockerHost) ImageExists(name string) bool {
	for _, img := range d.Images {
		if img == name {
			return true
		}
	}
	return false
}
