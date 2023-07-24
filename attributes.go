package nl80211

import (
	"net"

	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

// AttributeType is a type that can be used as a netlink attribute.
type AttributeType interface {
	bool | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | ~[]byte | string
}

// Nl80211AttributeSetter is a function that takes a netlink.AttributeEncoder. It is used to set attributes in a netlink message.
type Nl80211AttributeSetter func(*netlink.AttributeEncoder)

// AttributeSetter returns a function that takes a netlink.AttributeEncoder and calls its method for the given type with the given value.
func AttributeSetter[T AttributeType](typ int, value T) Nl80211AttributeSetter {
	switch v := any(value).(type) {
	case bool:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Flag(uint16(typ), v)
		}
	case int8:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Int8(uint16(typ), v)
		}
	case int16:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Int16(uint16(typ), v)
		}
	case int32:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Int32(uint16(typ), v)
		}
	case int64:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Int64(uint16(typ), v)
		}
	case uint8:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Uint8(uint16(typ), v)
		}
	case uint16:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Uint16(uint16(typ), v)
		}
	case uint32:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Uint32(uint16(typ), v)
		}
	case uint64:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Uint64(uint16(typ), v)
		}
	case []byte:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.Bytes(uint16(typ), v)
		}
	case string:
		return func(attrEncoder *netlink.AttributeEncoder) {
			attrEncoder.String(uint16(typ), v)
		}
	default:
		panic("invalid type")
	}
}

// NestedAttributeSetter returns a function that takes a netlink.AttributeEncoder and calls its Nested method with the given type and a function that calls the given attribute setters.
func NestedAttributeSetter(typ int, attrs ...Nl80211AttributeSetter) Nl80211AttributeSetter {
	return func(attrEncoder *netlink.AttributeEncoder) {
		attrEncoder.Nested(uint16(typ), func(attrEncoder *netlink.AttributeEncoder) error {
			for _, attr := range attrs {
				attr(attrEncoder)
			}
			return nil
		})
	}
}

// SetAttrWiphy returns a function that takes a netlink.AttributeEncoder and calls its Uint32 method with the given value and the NL80211_ATTR_WIPHY attribute type.
func SetAttrWiphy(wiphy uint32) Nl80211AttributeSetter {
	return AttributeSetter[uint32](unix.NL80211_ATTR_WIPHY, wiphy)
}

// SetAttrWiphyName returns a function that takes a netlink.AttributeEncoder and calls its String method with the given value and the NL80211_ATTR_WIPHY_NAME attribute type.
func SetAttrWiphyName(name string) Nl80211AttributeSetter {
	return AttributeSetter[string](unix.NL80211_ATTR_WIPHY_NAME, name)
}

// SetAttrIfindex returns a function that takes a netlink.AttributeEncoder and calls its Uint32 method with the given value and the NL80211_ATTR_IFINDEX attribute type.
func SetAttrIfindex(ifindex uint32) Nl80211AttributeSetter {
	return AttributeSetter[uint32](unix.NL80211_ATTR_IFINDEX, ifindex)
}

// SetAttrIfname returns a function that takes a netlink.AttributeEncoder and calls its String method with the given value and the NL80211_ATTR_IFNAME attribute type.
func SetAttrIfname(ifname string) Nl80211AttributeSetter {
	return AttributeSetter[string](unix.NL80211_ATTR_IFNAME, ifname)
}

// SetAttrIftype returns a function that takes a netlink.AttributeEncoder and calls its Uint32 method with the given value and the NL80211_ATTR_IFTYPE attribute type.
func SetAttrIftype(iftype int) Nl80211AttributeSetter {
	return AttributeSetter[uint32](unix.NL80211_ATTR_IFTYPE, uint32(iftype))
}

// SetAttrMac returns a function that takes a netlink.AttributeEncoder and calls its Bytes method with the given value and the NL80211_ATTR_MAC attribute type.
func SetAttrMac(mac net.HardwareAddr) Nl80211AttributeSetter {
	return AttributeSetter[net.HardwareAddr](unix.NL80211_ATTR_MAC, mac)
}

// SetAttrBeaconInterval returns a function that takes a netlink.AttributeEncoder and calls its Uint32 method with the given value and the NL80211_ATTR_BEACON_INTERVAL attribute type.
func SetAttrBeaconInterval(beaconInterval uint32) Nl80211AttributeSetter {
	return AttributeSetter[uint32](unix.NL80211_ATTR_BEACON_INTERVAL, beaconInterval)
}

// SetAttrWiphyFreq returns a function that takes a netlink.AttributeEncoder and calls its Uint32 method with the given value and the NL80211_ATTR_WIPHY_FREQ attribute type.
func SetAttrWiphyFreq(freq uint32) Nl80211AttributeSetter {
	return AttributeSetter[uint32](unix.NL80211_ATTR_WIPHY_FREQ, freq)
}