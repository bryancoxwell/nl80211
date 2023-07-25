package nl80211

import (
	"fmt"

	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

// DumpInterfaces returns a list of all wifi interfaces
func (c *Conn) DumpInterfaces() ([]*InterfaceInfo, error) {
	req, err := NewNl80211Message(unix.NL80211_CMD_GET_INTERFACE)
	if err != nil {
		return nil, err
	}

	res, err := c.conn.Execute(req, c.nl80211Family.ID, netlink.Request|netlink.Dump)
	if err != nil {
		return nil, err
	}

	var ifis []*InterfaceInfo
	for _, msg := range res {
		ifi, err := InterfaceFromNetlinkMessage(msg)
		if err != nil {
			return nil, err
		}
		ifis = append(ifis, &ifi)
	}

	return ifis, nil
}

// InterfaceByName returns the wifi interface with the given name
func (c *Conn) InterfaceByName(ifname string) (*InterfaceInfo, error) {
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

// DumpWiphys returns a list of all wifi physical devices
// func (c *Conn) DumpWiphys() ([]*Wiphy, error) {
// 	req, err := NewNl80211Message(unix.NL80211_CMD_GET_WIPHY)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := c.conn.Execute(req, c.nl80211Family.ID, netlink.Request|netlink.Dump)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var wifis []*Wiphy
// 	for _, msg := range res {
// 		w, err := WiphyFromNetlinkMessage(msg)
// 		if err != nil {
// 			return nil, err
// 		}
// 		wifis = append(wifis, &w)
// 	}
// 	return wifis, nil
// }

// SetInterfaceFreq sets the frequency of the given interface with the given channel mode
func (c *Conn) SetInterfaceFreq(ifindex uint32, freq, chmode uint32) error {
	msg, err := NewNl80211Message(
		unix.NL80211_CMD_SET_CHANNEL,
		SetAttrIfindex(ifindex),
		SetAttrWiphyFreq(freq),
		SetAttrWiphyChannelType(chmode),
	)
	if err != nil {
		return fmt.Errorf("SetInterfaceFreq: %w", err)
	}
	_, err = c.conn.Send(msg, c.nl80211Family.ID, netlink.Request|netlink.Acknowledge)
	if err != nil {
		return fmt.Errorf("SetInterfaceFreq: %w", err)
	}
	return nil
}

// SetInterfaceChannel sets the channel of the given interface with the given channel mode
func (c *Conn) SetInterfaceChannel(ifindex uint32, channel, chmode uint32) error {
	freq, err := wifiChannelToFreq(channel)
	if err != nil {
		return fmt.Errorf("SetInterfaceChannel: %w", err)
	}
	return c.SetInterfaceFreq(ifindex, freq, chmode)
}

// SetInterfaceType sets the interface down, changes the interface type, and then brings the interface back up.
func (c *Conn) SetInterfaceType(ifindex, iftype uint32) error {
	req, err := NewNl80211Message(
		unix.NL80211_CMD_SET_INTERFACE, 
		SetAttrIfindex(ifindex), 
		SetAttrIftype(iftype),
	)
	if err != nil {
		return err
	}

	_, err = c.conn.Execute(req, c.nl80211Family.ID, netlink.Request|netlink.Acknowledge)
	return err
}