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
	address := r.FormValue("address")
	if name == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}
	if address == "" {
		http.Error(w, "Missing address", http.StatusBadRequest)
		return
	}
	// Show a simple thank you message, then redirect to home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func main() {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileserver)
	mux.HandleFunc("/form", formHandle)
	mux.HandleFunc("/hello", defaultHandle)

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
