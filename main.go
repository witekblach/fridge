package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			http.NotFound(writer, request)
			return
		}
		fmt.Fprintf(writer, "Welcome AAAAAAA to the home page!")
	})

	mux.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Welcome to the A page!")
	})

	slog.Info("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
