package main

import (
	"log"
	"github.com/kataras/iris"
)

func main() {
	api, err := NewShortedAPI(ap.Data)
	if err != nil {
		log.Fatal(err)
	}
	app := iris.Default()
	app.Get("/{shorted:string regexp(^[a-zA-Z0-9]{2,6}$)}", api.shorted)
	app.Post("/s", api.create)
	app.Get("/s/{shorted:string regexp(^[a-zA-Z0-9]{2,6}$)}", api.query)
	app.Run(iris.Addr(ap.Address()))
}
