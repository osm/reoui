package reolink

import (
	"time"

	"github.com/osm/reoui/config"
)

type Client struct {
	// Basic connection details.
	Name     string
	address  string
	username string
	password string

	// Whether or not to use low or high quality for the stream.
	lowStreamQuality bool

	// Access token, acquired from the reolink camera and when it exipres.
	token     string
	expiresAt time.Time
}

func NewClient(c *config.Camera) *Client {
	return &Client{
		Name:             c.Name,
		address:          c.Address,
		username:         c.Username,
		password:         c.Password,
		lowStreamQuality: c.LowStreamQuality,
	}
}
