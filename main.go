package main

import (
	"log"
	"net/http"
	"github.com/incu6us/xml-creator/internal/api"
)

func main() {
	var err error

	srv := &http.Server{
		Addr:    ":8080",
		Handler: api.NewHandler(),

	}

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
