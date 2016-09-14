/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"errors"
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

// DockerHost holds the docker client we need to talk to docker
type DockerHost struct {
	Host   string
	Images []string
	client *docker.Client
	auth   *docker.AuthConfiguration
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

// SetAuthCredentials sets which account should be used for actions like pushing to docker hub
func (d *DockerHost) SetAuthCredentials(username, password string) error {
	d.auth = &docker.AuthConfiguration{
		Username: username,
		Password: password,
	}

	status, err := d.client.AuthCheck(d.auth)
	if err != nil {
		return err
	}

	fmt.Println(status.Status)

	return nil
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

// BuildImage builds a docker image and tags it
func (d *DockerHost) BuildImage(name, path string) (string, error) {
	var ostream OutputStream

	// Create a image from dockerfile
	opts := docker.BuildImageOptions{
		Name:         name,
		ContextDir:   path,
		OutputStream: &ostream,
	}

	err := d.client.BuildImage(opts)

	return ostream.Output(), err
}

// PushImage pushes a built image to docker hub
func (d *DockerHost) PushImage(name string) (string, error) {
	var ostream OutputStream

	opts := docker.PushImageOptions{
		Name:         name,
		OutputStream: &ostream,
	}

	if d.auth == nil {
		return "", errors.New("No authentication information")
	}

	err := d.client.PushImage(opts, *d.auth)

	return ostream.Output(), err
}
