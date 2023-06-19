package main

import (
	W "groupie-tracker/cmd/web"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", W.Start)
	mux.HandleFunc("/info/", W.ArtistInfo)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Server launched on http://127.0.0.1:5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
