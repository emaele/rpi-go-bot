// Copyright 2017 Shannon Wynter. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commontypes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type KeyFile struct {
	src string
	key []byte
}

func (f *KeyFile) UnmarshalText(text []byte) error {
	return f.Unmarshal(string(text))
}

func (f *KeyFile) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return f.Unmarshal(s)
}

func (f *KeyFile) Unmarshal(s string) error {
	if s[0:7] == "file://" {
		var err error
		if f.key, err = ioutil.ReadFile(s[7:]); err != nil {
			return fmt.Errorf("Unable to read keyfile: %s", err)
		}
		f.src = s
	} else {
		f.key = []byte(s)
	}
	return nil
}

func (f *KeyFile) MarshalText() ([]byte, error) {
	if f.src == "" {
		return f.key, nil
	}
	return []byte(f.src), nil
}

func (f *KeyFile) MarshalJSON() ([]byte, error) {
	bytes, err := f.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(bytes))
}

func (f *KeyFile) Bytes() []byte {
	return f.key
}
