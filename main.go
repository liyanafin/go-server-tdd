package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "public/form.html")

}
func formHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")
	if name == "" {
		http.Redirect(w, r, "/?error=Missing+name", http.StatusSeeOther)
		return
	}
	if email == "" {
		http.Redirect(w, r, "/?error=Missing+email", http.StatusSeeOther)
		return
	}
	if subject == "" {
		http.Redirect(w, r, "/?error=Missing+subject", http.StatusSeeOther)
		return
	}
	if message == "" {
		http.Redirect(w, r, "/?error=Missing+message", http.StatusSeeOther)
		return
	}
	// Redirect to home with status message
	http.Redirect(w, r, "/?status=success", http.StatusSeeOther)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		if status == "success" {
			// Read index.html and inject status message
			http.ServeFile(w, r, "public/index.html")
		} else {
			http.ServeFile(w, r, "public/index.html")
		}
	})
	mux.HandleFunc("/form", formHandle)
	mux.HandleFunc("/contact", defaultHandle)

	// Logging middleware for all requests
	logRequests := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
			h.ServeHTTP(w, r)
		})
	}

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", logRequests(mux)); err != nil {
		log.Fatal(err)
	}
}
