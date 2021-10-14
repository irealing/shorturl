package main

import (
	"log"

	iris "github.com/kataras/iris/v12"
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
	log.Fatal(app.Run(iris.Addr(ap.Address())))
}
