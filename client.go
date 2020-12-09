package aoc

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// HTTPClient is a simple interface for http.Client.
// Reason: mock client in tests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client for Advent of Code.
type Client struct {
	SessionToken string
	HTTPClient   HTTPClient

	logger *logrus.Logger
}

// RequestTimout is the timeout of a request in seconds.
const RequestTimout = 10

// NewClient creates a new client.
func NewClient(token string) Client {
	return Client{
		SessionToken: token,
		HTTPClient: &http.Client{
			Timeout: time.Second * RequestTimout,
		},
		logger: logrus.New(),
	}
}

// LogLevel sets logger level.
func (c *Client) LogLevel(level logrus.Level) {
	c.logger.SetLevel(level)
}
