package main

import (
	"os"
	"fmt"
	"log"
	"net/http"

	"github.com/rmitsubayashi/bodyweight-server/src/registry"

)

func main() {
	registry.NewRouter().Route()

	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
