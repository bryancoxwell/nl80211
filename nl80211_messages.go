package nl80211

import (
	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
)

// NewNl80211Message returns a generic netlink message
func NewNl80211Message(cmd int, attrs ...Nl80211AttributeSetter) (genetlink.Message, error) {
	attributeEncoder := netlink.NewAttributeEncoder()

	for _, attr := range attrs {
		attr(attributeEncoder)
	}
	data, err := attributeEncoder.Encode()
	if err != nil {
		return genetlink.Message{}, err
	}

	return genetlink.Message{
		Header: genetlink.Header{
			Command: uint8(cmd),
			Version: 1,
		},
		Data: data,
	}, nil
}
