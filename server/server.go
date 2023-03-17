package server

import (
	"fmt"
	"net/http"
	"time"
)

func RunServer() {
	server := &http.Server{
		Addr: ":8080",
		//Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is main page"))
	})
	fmt.Println("Server is listening on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		panic("Server is down")
	}
}
