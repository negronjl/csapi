// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "CSAPI": swagger Resource Client
//
// Command:
// $ goagen
// --design=github.com/negronjl/csapi/design
// --out=$(GOPATH)/src/github.com/negronjl/csapi
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// DownloadSwaggerJSON downloads swagger.json and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadSwaggerJSON(ctx context.Context, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: "/swagger.json"}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// DownloadSwaggerYaml downloads swagger.yaml and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadSwaggerYaml(ctx context.Context, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: "/swagger.yaml"}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}
