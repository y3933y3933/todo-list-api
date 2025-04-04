package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	port int
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.Parse()

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: mux,
	}

	log.Printf("starting server :%d\n", cfg.port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
