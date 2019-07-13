package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/irealing/argsparser"
)

var ap = &AppConfig{}

type AppConfig struct {
	Host string `param:"host" usage:"bind address"`
	Port int    `param:"port" usage:"bind port"`
	Data string `param:"data" usage:"data source path"`
}

func (ap *AppConfig) Address() string {
	return fmt.Sprintf("%s:%d", ap.Host, ap.Port)
}
func (ap *AppConfig) Validate() error {
	if ap.Host == "" || ap.Port < 0 || ap.Data == "" {
		return errors.New("error arguments")
	}
	return nil
}
func init() {
	parser := argsparser.New(ap)
	err := parser.Init()
	if err != nil {
		log.Fatal(err)
	}
	err = parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
}
