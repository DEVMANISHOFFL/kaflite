package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("kaflite started"))
}

func main() {
	http.HandleFunc("/", root)

	addr := ":8080"
	log.Println("starting kaflite (dev) on:", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("server error:", err)
	}
}
