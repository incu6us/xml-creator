package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const (
	PATH_PREFIX  = "/api/"
	HTTP_TIMEOUT = 120 * time.Second
)

type API struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type APIs []API

func NewHandler() http.Handler {

	router := mux.NewRouter().StrictSlash(true)
	middlewareHandler := http.TimeoutHandler(router, HTTP_TIMEOUT, "Server timedout!")

	for _, api := range apis {
		router.
		PathPrefix(PATH_PREFIX).
			Methods(api.Method).
			Path(api.Pattern).
			Name(api.Name).
			Handler(api.HandlerFunc)
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return middlewareHandler
}

var apis = APIs{
	API{
		"createXML",
		"POST",
		"/create",
		GetHandler().CreateObject,
	},
}
