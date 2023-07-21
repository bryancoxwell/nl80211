package nl80211

import (
	"errors"
	"fmt"
	"os"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
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

func (c *Conn) DumpInterfaces() ([]*Interface, error) {
	req, err := NewNl80211Message(unix.NL80211_CMD_GET_INTERFACE)
	if err != nil {
		return nil, err
	}

	res, err := c.conn.Execute(req, c.nl80211Family.ID, netlink.Request|netlink.Dump)
	if err != nil {
		return nil, err
	}

	var ifis []*Interface
	for _, msg := range res {
		ifi, err := InterfaceFromNetlinkMessage(msg)
		if err != nil {
			return nil, err
		}
		ifis = append(ifis, &ifi)
	}

	return ifis, nil
}

func (c *Conn) InterfaceByName(ifname string) (*Interface, error) {
	ifis, err := c.DumpInterfaces()
	if err != nil {
		return nil, err
	}
	for _, ifi := range ifis {
		if ifi.Name == ifname {
			return ifi, nil
		}
	}
	return nil, fmt.Errorf("interface %s not found", ifname)
}

func (c *Conn) SetInterfaceFreq(ifi *Interface, freq uint32) error {
	msg, err := NewNl80211Message(
		unix.NL80211_CMD_SET_CHANNEL,
		SetAttrIfindex(uint32(ifi.Index)),
		SetAttrWiphyFreq(freq),
	)
	if err != nil {
		return fmt.Errorf("SetInterfaceFreq: %w", err)
	}
	_, err = c.conn.Execute(msg, c.nl80211Family.ID, netlink.Request|netlink.Acknowledge)
	if err != nil {
		return fmt.Errorf("SetInterfaceFreq: %w", err)
	}
	return nil
}

// setInterfaceType sets the interface down, changes the interface type, and then brings the interface back up.
func (c *Conn) SetInterfaceType(iftype InterfaceType, ifi *Interface) error {
	req, err := NewNl80211Message(
		unix.NL80211_CMD_SET_INTERFACE, 
		SetAttrIfindex(uint32(ifi.Index)), 
		SetAttrIftype(iftype),
	)
	if err != nil {
		return err
	}

	_, err = c.conn.Execute(req, c.nl80211Family.ID, netlink.Request|netlink.Acknowledge)
	return err
}

func (c *Conn) SetModeMonitor(ifi *Interface) error {
	return c.SetInterfaceType(NL80211_IFTYPE_MONITOR, ifi)
}