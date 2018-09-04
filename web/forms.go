package main

import (
	"net/url"
)

type NewShorted struct {
	URL string `json:"url"`
}

func (ns *NewShorted) Validate() error {
	_, err := url.Parse(ns.URL)
	return err
}
