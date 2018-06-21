// Copyright 2017 Shannon Wynter. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commontypes

import (
	"encoding/json"
	"regexp"
)

type Regexp struct {
	*regexp.Regexp
}

type Regexps []Regexp

func (r *Regexp) UnmarshalText(text []byte) error {
	return r.Unmarshal(string(text))
}

func (r *Regexp) UnmarshalTOML(text []byte) error {
	return r.Unmarshal(string(text[1 : len(text)-1]))
}

func (r *Regexp) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	return r.Unmarshal(s)
}

func (r *Regexp) Unmarshal(s string) (err error) {
	r.Regexp, err = regexp.Compile(s)
	return
}

func (r *Regexp) MarshalText() ([]byte, error) {
	return []byte(r.Regexp.String()), nil
}

func (r *Regexp) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Regexp.String())
}

func (r *Regexps) MatchString(s string) bool {
	for _, e := range *r {
		if e.MatchString(s) {
			return true
		}
	}
	return false
}

func (r *Regexps) Match(b []byte) bool {
	for _, e := range *r {
		if e.Match(b) {
			return true
		}
	}
	return false
}
