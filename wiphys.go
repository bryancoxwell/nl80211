package nl80211

// import (
// 	"fmt"

// 	"github.com/mdlayher/genetlink"
// 	"github.com/mdlayher/netlink"
// 	"golang.org/x/sys/unix"
// )

// type Wiphy struct {
// 	Index uint32
// }

// func WiphyFromNetlinkMessage(msg genetlink.Message) (Wiphy, error) {
// 	w := Wiphy{}
// 	ad, err := netlink.NewAttributeDecoder(msg.Data)
// 	if err != nil {
// 		return w, err
// 	}
// 	for ad.Next() {
// 		// fmt.Println(ad.Type())
// 		switch ad.Type() {
// 		case unix.NL80211_ATTR_WIPHY:
// 			w.Index = ad.Uint32()
// 		case unix.NL80211_ATTR_WIPHY_BANDS:
// 			ad.Nested(
// 				func(nad *netlink.AttributeDecoder) error {
// 					for nad.Next() {
// 						fmt.Println("-----------")
// 						if nad.Type() == unix.NL80211_BAND_5GHZ {
// 							fmt.Println("5GHz")
// 							nad.Nested(
// 								func(nnad *netlink.AttributeDecoder) error {
// 									for nnad.Next() {
// 										if nnad.Type() == unix.NL80211_BAND_ATTR_FREQS {
// 											nnad.Nested(
// 												func(nnnad *netlink.AttributeDecoder) error {
// 													for nnnad.Next() {
// 														nnnad.Nested(
// 															func(nnnnad *netlink.AttributeDecoder) error {
// 																for nnnnad.Next() {
// 																	if nnnnad.Type() == unix.NL80211_FREQUENCY_ATTR_FREQ {
// 																		fmt.Println(nnnnad.Uint32())
// 																	}
// 																}
// 																return nil
// 															},
// 														)
// 													}
// 													return nil
// 												},
// 											)
// 										}
// 									}
// 									return nil
// 								},
// 							)
// 						}
// 					}
// 					return nil
// 				},
// 			)
// 		}
// 	}
// 	fmt.Println("-----------")
// 	return w, nil
// }

// type WiphyBand struct {
// 	Band int
	
// }