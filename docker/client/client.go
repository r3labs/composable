/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/docker/docker/api"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/registry"
)

// Client ...
type Client struct {
	registry string
	dc       *client.Client
	auth     *types.AuthConfig
}

// New ...
func New(host, reg string) (*Client, error) {
	var hc *http.Client

	if reg == "" {
		reg = registry.IndexServer
	}

	cli, err := client.NewClient(host, api.DefaultVersion, hc, nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		dc:       cli,
		registry: reg,
	}, nil
}

// HasImage ...
func (c *Client) HasImage(image string) (bool, error) {
	images, err := c.dc.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return false, err
	}

	for _, img := range images {
		for _, t := range img.RepoTags {
			if t == image {
				return true, nil
			}
		}
	}

	return false, nil
}

// DeleteImage ...
func (c *Client) DeleteImage(image string) error {
	opts := types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	}
	_, err := c.dc.ImageRemove(context.Background(), image, opts)
	return err
}

// BuildImage ...
func (c *Client) BuildImage(name, path string) (io.ReadCloser, error) {
	tar, err := Tar(path)
	if err != nil {
		return nil, err
	}

	opts := types.ImageBuildOptions{
		Tags: []string{name},
	}

	resp, err := c.dc.ImageBuild(context.Background(), tar, opts)
	return resp.Body, err
}

// PushImage ...
func (c *Client) PushImage(image string) (io.ReadCloser, error) {
	data, err := json.Marshal(c.auth)
	if err != nil {
		return nil, err
	}

	opts := types.ImagePushOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(data),
	}

	return c.dc.ImagePush(context.Background(), image, opts)
}

// Login ...
func (c *Client) Login(username, password string) error {
	c.auth = &types.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: c.registry,
	}

	_, err := c.dc.RegistryLogin(context.Background(), *c.auth)

	return err
}
