package nl80211

import (
	"fmt"
	"testing"
	"time"

	// "github.com/mdlayher/netlink"
)

// func TestDumpWiphys(t *testing.T) {
// 	c, err := Dial(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	msg, err := NewNl80211Message(NL80211_CMD_GET_INTERFACE)
// 	if err != nil {
// 		panic(err)
// 	}
// 	responses, err := c.conn.Execute(msg, c.nl80211Family.ID, netlink.Request|netlink.Dump)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, response := range responses {
// 		ad, err := netlink.NewAttributeDecoder(response.Data)
// 		if err != nil {
// 			panic(err)
// 		}
// 		for ad.Next() {
// 			fmt.Println(ad.Type())
// 		}
// 		fmt.Println("---------")
// 	}
// }

// func TestDumpInterfaces(t *testing.T) {
// 	c, err := Dial(nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	ifis, err := c.DumpInterfaces()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, ifi := range ifis {
// 		fmt.Println(ifi)
// 	}
// }

func TestSetInterfaceFreq(t *testing.T) {
	c, err := Dial(nil)
	if err != nil {
		t.Fatal(err)
	}
	ifi, err := c.InterfaceByName("wlan1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ifi)
	start := time.Now()
	err = c.SetInterfaceFreq(ifi, 2412)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("SetInterfaceFreq took", time.Since(start))
}