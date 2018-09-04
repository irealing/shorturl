package main

import (
	"log"
	"github.com/kataras/iris"
)

func main() {
	api, err := NewShortedAPI("data")
	if err != nil {
		log.Fatal(err)
	}
	app := iris.Default()
	app.Get("/{shorted:string regexp(^[a-zA-Z0-9]{2,6}$)}", api.shorted)
	app.Post("/", api.create)
	app.Run(iris.Addr("127.0.0.1:8001"))
}
