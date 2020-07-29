package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	http.ListenAndServe(":9090", nil)
}
