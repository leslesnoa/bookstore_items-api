package main

import (
	"github.com/gorilla/mux"
	"github.com/leslesnoa/bookstore_items-api/app"
)

var (
	router = mux.NewRouter()
)

func main() {
	app.StartApplication()
}
