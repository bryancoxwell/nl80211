package nl80211

import (
	"errors"
	"fmt"
	"os"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
)

type Conn struct {
	conn            *genetlink.Conn
	nl80211Family genetlink.Family
}

func Dial(config *netlink.Config) (*Conn, error) {
	c, err := genetlink.Dial(nil)
	if err != nil {
		return nil, err
	}
	family, err := c.GetFamily("nl80211")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("nl80211 not supported by kernel")
		}
	}
	return &Conn{conn: c, nl80211Family: family}, nil
}

func (c *Conn) Close() error {
	return c.conn.Close()
}