package aoc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadInput downloads the requested input file.
func (c *Client) DownloadInput(year, day int) (io.ReadCloser, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: c.SessionToken})

	resp, err := c.HTTPClient.Do(req) //nolint:bodyclose
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, DownloadError{StatusCode: resp.StatusCode}
	}

	return resp.Body, nil
}

// DownloadAndSaveInput downloads and saves the requested input file.
func (c *Client) DownloadAndSaveInput(year, day int, targetFile string) error {
	file, err := c.DownloadInput(year, day)
	if err != nil {
		return err
	}

	defer file.Close()

	f, err := os.Create(targetFile)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = io.Copy(f, file)

	return err
}
