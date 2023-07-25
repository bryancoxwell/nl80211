package nl80211

import (
	"net"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

type InterfaceInfo struct {
	Index uint32
	Name string
	WiphyIndex uint32
	IfType int
	Mac net.HardwareAddr
	WiphyFreq uint32
}

func InterfaceFromNetlinkMessage(msg genetlink.Message) (InterfaceInfo, error) {
	ifi := InterfaceInfo{}
	ad, err := netlink.NewAttributeDecoder(msg.Data)
	if err != nil {
		return ifi, err
	}
	for ad.Next() {
		switch ad.Type() {
		case unix.NL80211_ATTR_IFINDEX:
			ifi.Index = ad.Uint32()
		case unix.NL80211_ATTR_IFNAME:
			ifi.Name = ad.String()
		case unix.NL80211_ATTR_WIPHY:
			ifi.WiphyIndex = ad.Uint32()
		case unix.NL80211_ATTR_IFTYPE:
			ifi.IfType = int(ad.Uint32())
		case unix.NL80211_ATTR_MAC:
			ifi.Mac = ad.Bytes()
		case unix.NL80211_ATTR_WIPHY_FREQ:
			ifi.WiphyFreq = ad.Uint32()
		}
	}
	return ifi, nil
}
