package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//router
	r := mux.NewRouter()

	//register a request handler
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	//Serving static assets
	fs := http.FileServer(http.Dir("static/"))

	//  strip away a part of the url path.
	http.Handle("static/", http.StripPrefix("static/", fs))

	// listen to an http connection
	fmt.Printf("server running on port 8000")
	http.ListenAndServe(":8000", r)

}
