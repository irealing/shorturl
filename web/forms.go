package main

import (
	"regexp"
	"github.com/kataras/iris/core/errors"
)

const (
	regexURLStr = "^https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)$"
)

var regex = regexp.MustCompile(regexURLStr)

type NewShorted struct {
	URL string `json:"url"`
}

func (ns *NewShorted) Validate() error {
	if regex.MatchString(ns.URL) {
		return nil
	}
	return errors.New("error format")
}
