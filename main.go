package main

import (
	"net/http"
	"os"

	"log"
)

func main() {
	log.Println(os.Getenv("GREETING"))

	log.Println("starting server...")

	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		w.Write([]byte("HEY!"))
	})))
}
