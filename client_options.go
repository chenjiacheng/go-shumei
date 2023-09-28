package shumei

import "time"

type ClientOptions func(*Client) error

func WithRegion(region string) ClientOptions {
	return func(c *Client) error {
		c.Region = region
		return nil
	}
}

func WithTimeout(timeout time.Duration) ClientOptions {
	return func(c *Client) error {
		c.Timeout = timeout
		return nil
	}
}
