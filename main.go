package main

import (
	"fmt"
	"log"
	"net/http"
)

/// Parse form
///
/// if error is nil print out the [name] and [address] supplied in the form

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

/// Handle request from user
///
/// if path supplied by user is not /hello then throw 404
/// throw Method not supported error if Method is not GET

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	
	// Return a handler that serves HTTP requests with the contents of the file system rooted at root 
	fileServer := http.FileServer(http.Dir("./static"))
	
	// register the handler
	http.Handle("/", fileServer)

	// Register the handler function
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err:= http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}