package client

import (
	"context"
	"net/http"
	"os"

	"github.com/docker/docker/api"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Client struct {
	dc *client.Client
}

// New ...
func New(host string) (*Client, error) {
	var hc *http.Client

	cli, err := client.NewClient(host, api.DefaultVersion, hc, nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		dc: cli,
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

// CreateImage ...
func (c *Client) BuildImage(name, path string) error {
	buildContext, err := os.Open(path)
	defer buildContext.Close()
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Tags: []string{name},
	}

	_, err = c.dc.ImageBuild(context.Background(), buildContext, opts)
	return err
}
