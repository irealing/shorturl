package main

import (
	"shorturl"
	"github.com/kataras/iris"
)

type ShortedAPI struct {
	handler *shorturl.ShortedHandler
}

func (sa *ShortedAPI) shorted(ctx iris.Context) {
	shorted := ctx.Params().Get("shorted")
	if shorted == "" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	if r, err := sa.handler.Find(shorted); err == nil {
		ctx.StatusCode(iris.StatusTemporaryRedirect)
		ctx.Header("Location", r.URL)
	} else {
		ctx.StatusCode(iris.StatusNotFound)
	}
}
func (sa *ShortedAPI) create(ctx iris.Context) {
}
func NewShortedAPI(fp string) (*ShortedAPI, error) {
	handler, err := shorturl.NewHandler(fp)
	if err != nil {
		return nil, err
	}
	return &ShortedAPI{handler: handler}, nil
}
