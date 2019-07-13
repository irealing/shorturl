package main

import (
	"shorturl"

	"github.com/kataras/iris"
	"fmt"
)

type BaseRet struct {
	ErrNo int         `json:"err_no"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

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
	form := new(NewShorted)
	if err := ctx.ReadJSON(form); err != nil || form.Validate() != nil {
		ctx.JSON(&BaseRet{ErrNo: 400, Msg: "failed"})
		return
	}
	var ret *BaseRet
	if su, err := sa.handler.Create(form.URL); err != nil {
		ret = &BaseRet{ErrNo: 500, Msg: "failed"}
	} else {
		ret = &BaseRet{ErrNo: 0, Msg: "success", Data: su.Shorted}
	}
	ctx.JSON(ret)
}
func (sa *ShortedAPI) query(ctx iris.Context) {
	shorted := ctx.Params().Get("shorted")
	if shorted == "" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	var ret *BaseRet
	if r, err := sa.handler.Find(shorted); err == nil {
		ret = &BaseRet{ErrNo: 0, Msg: "success", Data: r.URL}
	} else {
		ret = &BaseRet{ErrNo: iris.StatusNotFound, Msg: fmt.Sprintf("%v", err)}
	}
	ctx.JSON(ret)
}
func NewShortedAPI(fp string) (*ShortedAPI, error) {
	handler, err := shorturl.NewHandler(fp)
	if err != nil {
		return nil, err
	}
	return &ShortedAPI{handler: handler}, nil
}
