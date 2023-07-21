package nl80211

import (
	"fmt"
	"testing"

	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)


func TestNestedAttributeSetter(t *testing.T) {
	msg, err := NewNl80211Message(
		unix.NL80211_CMD_UNSPEC, 
		NestedAttributeSetter(
			NL80211_ATTR_KEY,
			AttributeSetter[uint32](NL80211_ATTR_KEY_TYPE, 10),
			AttributeSetter[uint8](NL80211_ATTR_KEY_DEFAULT, 20),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	ad, err := netlink.NewAttributeDecoder(msg.Data)
	if err != nil {
		t.Fatal(err)
	}
	for ad.Next() {
		switch ad.Type() {
		case uint16(NL80211_ATTR_KEY):
			ad.Nested(func(nad *netlink.AttributeDecoder) error {
				for nad.Next() {
					switch nad.Type() {
					case uint16(NL80211_ATTR_KEY_TYPE):
						val := nad.Uint32()
						if val != 10 {
							t.Fatal(fmt.Errorf("expected 10, got %d", val))
						}
					case uint16(NL80211_ATTR_KEY_DEFAULT):
						val := nad.Uint8()
						if val != 20 {
							t.Fatal(fmt.Errorf("expected 20, got %d", val))
						}
					}
				}
				return nil
			})
		}
	}
}