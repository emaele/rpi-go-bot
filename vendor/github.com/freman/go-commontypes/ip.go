// Copyright 2017 Shannon Wynter. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commontypes

import (
	"encoding/json"
	"net"
)

type IP struct {
	net.IP
}

func (i *IP) IsUnspecified() bool {
	return i == nil || i.IP == nil || i.IP.IsUnspecified()
}

func (i *IP) UnmarshalText(text []byte) error {
	return i.IP.UnmarshalText(text)
}

func (i *IP) UnmarshalTOML(text []byte) error {
	return i.IP.UnmarshalText(text[1 : len(text)-1])
}

func (i *IP) UnmarshalJSON(data []byte) error {
	var text string
	err := json.Unmarshal(data, &text)
	if err != nil {
		return err
	}
	return i.IP.UnmarshalText([]byte(text))
}

func (i *IP) MarshalText() ([]byte, error) {
	return []byte(i.IP.String()), nil
}

func (i *IP) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.IP.String())
}
