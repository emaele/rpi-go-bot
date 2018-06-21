package commontypes

import (
	"bytes"
	"encoding/json"
	"net"
	"reflect"
	"testing"
)

var networkTests = []struct {
	text    []byte
	json    []byte
	ipNet   *net.IPNet
	expText []byte
	expJSON []byte
}{{
	[]byte("0.0.0.0/0"), []byte("\"0.0.0.0/0\""),
	&net.IPNet{IP: net.IP{0x0, 0x0, 0x0, 0x0}, Mask: net.IPMask{0x0, 0x0, 0x0, 0x0}},
	nil, nil,
}, {
	[]byte("192.168.0.0/24"), []byte("\"192.168.0.0/24\""),
	&net.IPNet{IP: net.IP{0xc0, 0xa8, 0x0, 0x0}, Mask: net.IPMask{0xff, 0xff, 0xff, 0x0}},
	nil, nil,
}, {
	[]byte("192.168.0.1"), []byte("\"192.168.0.1\""),
	&net.IPNet{IP: net.IP{0xc0, 0xa8, 0x0, 0x1}, Mask: net.IPMask{0xff, 0xff, 0xff, 0xff}},
	[]byte("192.168.0.1/32"), []byte("\"192.168.0.1/32\""),
}}

func TestNetworkUnmarshalText(t *testing.T) {
	for _, test := range networkTests {
		var n Network
		if n.UnmarshalText(test.text); !reflect.DeepEqual(n.IPNet, test.ipNet) {
			t.Errorf("%s should be %s got %s", string(test.text), test.ipNet, n.IPNet)
		}
	}
}

func TestNetworkUnmarshalJSON(t *testing.T) {
	for _, test := range networkTests {
		var n Network
		if n.UnmarshalJSON(test.json); !reflect.DeepEqual(n.IPNet, test.ipNet) {
			t.Errorf("%s should be %d got %d", string(test.json), test.ipNet, n.IPNet)
		}
	}
}

func TestNetworkMarshalText(t *testing.T) {
	for _, test := range networkTests {
		n := Network{test.ipNet}
		expect := test.text
		if test.expText != nil {
			expect = test.expText
		}
		if text, err := n.MarshalText(); err != nil || !bytes.Equal(text, expect) {
			t.Errorf("%d should be %s got %s (error: %v)", test.ipNet, string(expect), string(text), err)
		}
	}
}

func TestNetworkMarshalJSON(t *testing.T) {
	for _, test := range networkTests {
		n := Network{test.ipNet}
		expect := test.json
		if test.expJSON != nil {
			expect = test.expJSON
		}
		if text, err := n.MarshalJSON(); err != nil || !bytes.Equal(text, expect) {
			t.Errorf("%d should be %s got %s (error: %v)", test.ipNet, string(expect), string(text), err)
		}
	}
}

func TestNetworks(t *testing.T) {
	// Don't really need to test both at this text and json at this point
	jsonBlob := []byte(`["192.168.0.0/24", "192.168.1.1", "10.0.0.0/8"]`)
	tests := []struct {
		ip     net.IP
		result bool
	}{
		{net.IPv4(135, 104, 0, 0), false},
		{net.IPv4(192, 168, 0, 1), true},
		{net.IPv4(192, 168, 1, 1), true},
		{net.IPv4(192, 168, 1, 2), false},
		{net.IPv4(10, 20, 30, 40), true},
	}

	var networks Networks

	if !networks.Empty() {
		t.Fatal("Empty networks should be empty")
	}

	err := json.Unmarshal(jsonBlob, &networks)
	if err != nil {
		t.Errorf("Problem unmarshalling %s: %s", jsonBlob, err)
	}

	if networks.Empty() {
		t.Error("Full networks should not be empty")
	}

	for _, test := range tests {
		if result := networks.Contains(test.ip); result != test.result {
			t.Errorf("Contains returned %v for %s, %v was expected", result, test.ip, test.result)
		}
	}
}
