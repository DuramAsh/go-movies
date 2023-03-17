package main

import "net/http"

func RunServer() {
	server := &http.Server{}
	server.ListenAndServe()
}
