package main

import (
	"log"
	"fmt"
	"net/http"
)

func StartServer(port int, path string, handleFunc http.HandlerFunc) error {
	handler := new(http.ServeMux)
	handler.HandleFunc(path, handleFunc)
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: handler,
	}
	log.Print("Server started on\n127.0.0.1:", port)
	return server.ListenAndServe()
}