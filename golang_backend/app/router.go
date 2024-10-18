package app

import (
	"github.com/julienschmidt/httprouter"
	"golang_backend/api"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/ask", api.AskHandler)

	return router
}
