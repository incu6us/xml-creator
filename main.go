package main

import (
	"log"
	"net/http"
	"xml-creator/internal/api"
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
