package main

import (
	"golang_backend/app"
	"golang_backend/helper"
	"net/http"
)

func main() {
	router := app.NewRouter()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicErr(err)
}
