package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortnener/handler"
	"url-shortnener/store"

	"go-micro.dev/v5/web"
)

func main() {
	// Create a new web service
	service := web.NewService(
		web.Name("url-shortnener"),
		web.Address(":8080"),
	)

	// Initialize the service
	service.Init()

	// Set up a route and handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the URL Shortener API")
	})
	http.HandleFunc("/create-short-url", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateShortUrl(w, r)
	})
	http.HandleFunc("/shortUrl", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleShortUrlRedirect(w, r)
	})
	store.InitializeStore()

	// Assign the handler to the service
	service.Handle("/", http.DefaultServeMux)

	// Start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
