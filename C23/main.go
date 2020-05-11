package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const indexHTML = `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Hello World</title>
				<script src="/static/app.js"></script>
				<link rel="stylesheet" href="/static/app.css"">
			</head>
			<body>
			Hello, gopher!<br>
			<img src="https://blog.golang.org/go-brand/logos.jpg" height="100">
			</body>
		</html>
	`

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if pusher, ok := w.(http.Pusher); ok {
			if err := pusher.Push("/static/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app.cs", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}

		fmt.Fprintf(w, indexHTML)
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("."))))

	log.Println("Server started at :7000")
	err := http.ListenAndServeTLS(":7000", "server.crt", "server.key", nil)
	if err != nil {
		panic(err)
	}
}
