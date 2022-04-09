package udp

import (
	"github.com/racing-telemetry/f1-cli/internal"
	"net"
)

type Client struct {
	conn *net.UDPConn
}

func Serve(addr *net.UDPAddr) (*Client, error) {
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c *Client) ReadSocket() ([]byte, error) {
	buf := make([]byte, internal.BufferSize)
	_, _, err := c.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, err
	}

	return buf, err
}

func (c *Client) Close() error {
	return c.conn.Close()
}
