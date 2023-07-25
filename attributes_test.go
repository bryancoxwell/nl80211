package nl80211

import (
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/mdlayher/netlink"
)

// testAttributeEncoder returns an attribute encoder with a fixed byte order
func testAttributeEncoder() *netlink.AttributeEncoder {
	ae := netlink.NewAttributeEncoder()
	ae.ByteOrder = binary.LittleEndian
	return ae
}

func TestFlagAttributeSetter(t *testing.T) {
	type test struct {
		input bool
		attrId int
		want []byte
	}

	tests := []test{
		{true, 3, []byte{4, 0, 3, 0}},
		{false, 4, []byte{}},
	}

	for _, test := range tests {
		ae := testAttributeEncoder()
		f := AttributeSetter[bool](test.attrId, test.input)
		f(ae)
		got, err := ae.Encode()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("got %v, want %v", got, test.want)
		}
	}
}

func TestInt8AttributeSetter(t *testing.T) {
	type test struct {
		input int8
		attrId int
		want []byte
	}

	tests := []test{
		{3, 10, []byte{5, 0, 10, 0, 3, 0, 0, 0}},
		{5, 12, []byte{5, 0, 12, 0, 5, 0, 0, 0}},
	}

	for _, test := range tests {
		ae := testAttributeEncoder()
		f := AttributeSetter[int8](test.attrId, test.input)
		f(ae)
		got, err := ae.Encode()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("got %v, want %v", got, test.want)
		}
	}
}

func TestInt16AttributeSetter(t *testing.T) {
	type test struct {
		input int16
		attrId int
		want []byte
	}

	tests := []test{
		{3, 10, []byte{6, 0, 10, 0, 3, 0, 0, 0}},
		{5, 12, []byte{6, 0, 12, 0, 5, 0, 0, 0}},
	}

	for _, test := range tests {
		ae := testAttributeEncoder()
		f := AttributeSetter[int16](test.attrId, test.input)
		f(ae)
		got, err := ae.Encode()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("got %v, want %v", got, test.want)
		}
	}
}

func TestInt32AttributeSetter(t *testing.T) {
	type test struct {
		input int32
		attrId int
		want []byte
	}

	tests := []test{
		{3, 6, []byte{8, 0, 6, 0, 3, 0, 0, 0}},
		{5, 20, []byte{8, 0, 20, 0, 5, 0, 0, 0}},
	}

	for _, test := range tests {
		ae := testAttributeEncoder()
		f := AttributeSetter[int32](test.attrId, test.input)
		f(ae)
		got, err := ae.Encode()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("got %v, want %v", got, test.want)
		}
	}
}

func TestStringAttributeSetter(t *testing.T) {
	type test struct {
		input string
		attrId int
		want []byte
	}
	tests := []test{
		{"foobar", 2, []byte{11, 0, 2, 0, 102, 111, 111, 98, 97, 114, 0, 0}},
	}

	for _, test := range tests {
		ae := testAttributeEncoder()
		f := AttributeSetter[string](test.attrId, test.input)
		f(ae)
		got, err := ae.Encode()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("got %v, want %v", got, test.want)
		}
	}
}

func TestNestedAttributeSetter(t *testing.T) {
	ta := AttributeSetter[bool](3, true) 
	tb := AttributeSetter[uint32](5, 20)
	tc := AttributeSetter[string](2, "foobar")
	
	ae := testAttributeEncoder()

	f := NestedAttributeSetter(
		66, 
		ta,
		tb,
		tc,
	)
	
	f(ae)
	got, err := ae.Encode()
	want := []byte{28, 0, 66, 128, 4, 0, 3, 0, 8, 0, 5, 0, 20, 0, 0, 0, 11, 0, 2, 0, 102, 111, 111, 98, 97, 114, 0, 0}
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %v, want %v", got, want)
	}
}