package main

import (
	"log"

	"github.com/fabiodemo/distributed-systems-golang/teste2/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
