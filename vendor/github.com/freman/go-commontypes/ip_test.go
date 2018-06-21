package commontypes

import (
	"bytes"
	"net"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

var ipTests = []struct {
	text    []byte
	json    []byte
	ip      net.IP
	expText []byte
	expJSON []byte
}{{
	[]byte("0.0.0.0"), []byte("\"0.0.0.0\""),
	net.IP{0x0, 0x0, 0x0, 0x0},
	nil, nil,
}, {
	[]byte("192.168.0.0"), []byte("\"192.168.0.0\""),
	net.IP{0xc0, 0xa8, 0x0, 0x0},
	nil, nil,
}, {
	[]byte("192.168.0.1"), []byte("\"192.168.0.1\""),
	net.IP{0xc0, 0xa8, 0x0, 0x1},
	[]byte("192.168.0.1"), []byte("\"192.168.0.1\""),
}}

func TestIpUnmarshalText(t *testing.T) {
	spew.Config.DisableMethods = true
	spew.Config.DisablePointerMethods = true
	for _, test := range ipTests {
		var i IP
		if i.UnmarshalText(test.text); !i.Equal(test.ip) {
			spew.Dump(i.IP, test.ip)
			t.Errorf("%s should be %s got %s", string(test.text), test.ip, i.IP)
		}
	}
}

func TestIpUnmarshalJSON(t *testing.T) {
	for _, test := range ipTests {
		var i IP
		if i.UnmarshalJSON(test.json); !i.Equal(test.ip) {
			t.Errorf("%s should be %d got %d", string(test.json), test.ip, i.IP)
		}
	}
}

func TestIpMarshalText(t *testing.T) {
	for _, test := range ipTests {
		i := IP{test.ip}
		expect := test.text
		if test.expText != nil {
			expect = test.expText
		}
		if text, err := i.MarshalText(); err != nil || !bytes.Equal(text, expect) {
			t.Errorf("%d should be %s got %s (error: %v)", test.ip, string(expect), string(text), err)
		}
	}
}

func TestIpMarshalJSON(t *testing.T) {
	for _, test := range ipTests {
		i := IP{test.ip}
		expect := test.json
		if test.expJSON != nil {
			expect = test.expJSON
		}
		if text, err := i.MarshalJSON(); err != nil || !bytes.Equal(text, expect) {
			t.Errorf("%d should be %s got %s (error: %v)", test.ip, string(expect), string(text), err)
		}
	}
}
