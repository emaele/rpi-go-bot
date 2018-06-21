package commontypes

import (
	"bytes"
	"testing"
	"time"
)

var durationTests = []struct {
	text     []byte
	json     []byte
	duration time.Duration
	expText  []byte
	expJSON  []byte
}{{
	[]byte("1h"), []byte("\"1h\""),
	1 * time.Hour,
	[]byte("1h0m0s"), []byte("\"1h0m0s\""),
}, {
	[]byte("3m2s"), []byte("\"3m2s\""),
	3*time.Minute + 2*time.Second,
	nil, nil,
}}

func TestDurationUnmarshalText(t *testing.T) {
	for _, test := range durationTests {
		var d Duration
		if d.UnmarshalText(test.text); d.Duration != test.duration {
			t.Errorf("%s should be %d got %d", string(test.text), test.duration, d)
		}
	}
}

func TestDurationUnmarshalJSON(t *testing.T) {
	for _, test := range durationTests {
		var d Duration
		if d.UnmarshalJSON(test.json); d.Duration != test.duration {
			t.Errorf("%s should be %d got %d", string(test.json), test.duration, d)
		}
	}
}

func TestDurationMarshalText(t *testing.T) {
	for _, test := range durationTests {
		d := Duration{test.duration}
		expect := test.text
		if test.expText != nil {
			expect = test.expText
		}
		if text, err := d.MarshalText(); err != nil || !bytes.Equal(text, expect) {
			t.Errorf("%d should be %s got %s (error: %v)", test.duration, string(expect), string(text), err)
		}
	}
}

func TestDurationMarshalJSON(t *testing.T) {
	for _, test := range durationTests {
		d := Duration{test.duration}
		expect := test.json
		if test.expJSON != nil {
			expect = test.expJSON
		}
		if json, err := d.MarshalJSON(); err != nil || !bytes.Equal(json, expect) {
			t.Errorf("%d should be %s got %s (error: %v)", test.duration, string(expect), string(json), err)
		}
	}
}
