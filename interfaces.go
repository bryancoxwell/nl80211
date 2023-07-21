package nl80211

import (
	"net"

	"github.com/mdlayher/netlink"
	"github.com/mdlayher/genetlink"
)

type InterfaceType uint32

const (
	NL80211_IFTYPE_UNSPECIFIED InterfaceType = iota
	NL80211_IFTYPE_ADHOC
	NL80211_IFTYPE_STATION
	NL80211_IFTYPE_AP
	NL80211_IFTYPE_AP_VLAN
	NL80211_IFTYPE_WDS
	NL80211_IFTYPE_MONITOR
	NL80211_IFTYPE_MESH_POINT
	NL80211_IFTYPE_P2P_CLIENT
	NL80211_IFTYPE_P2P_GO
	NL80211_IFTYPE_P2P_DEVICE
	NL80211_IFTYPE_OCB
	NL80211_IFTYPE_NAN
	NUM_NL80211_IFTYPES
	NL80211_IFTYPE_MAX = NUM_NL80211_IFTYPES - 1
)

type Interface struct {
	Index uint32
	Name string
	WiphyIndex uint32
	Type InterfaceType	
	Mac net.HardwareAddr
	WiphyFreq uint32
}

func InterfaceFromNetlinkMessage(msg genetlink.Message) (Interface, error) {
	ifi := Interface{}
	ad, err := netlink.NewAttributeDecoder(msg.Data)
	if err != nil {
		return ifi, err
	}
	for ad.Next() {
		switch ad.Type() {
		case uint16(NL80211_ATTR_IFINDEX):
			ifi.Index = ad.Uint32()
		case uint16(NL80211_ATTR_IFNAME):
			ifi.Name = ad.String()
		case uint16(NL80211_ATTR_WIPHY):
			ifi.WiphyIndex = ad.Uint32()
		case uint16(NL80211_ATTR_IFTYPE):
			ifi.Type = InterfaceType(ad.Uint32())
		case uint16(NL80211_ATTR_MAC):
			ifi.Mac = ad.Bytes()
		case uint16(NL80211_ATTR_WIPHY_FREQ):
			ifi.WiphyFreq = ad.Uint32()
		}
	}
	return ifi, nil
}
