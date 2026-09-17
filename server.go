package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home page")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to know more of us")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact us: htps://apiblogmy/blog")
}

func apiPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "active", "version": "1.0"}`)
}

func healthPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func calls() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/api/", apiPage)
	http.HandleFunc("/health", healthPage)

	fmt.Println("Server starting on port: 8080")

	http.ListenAndServe(":8080", nil)
}
