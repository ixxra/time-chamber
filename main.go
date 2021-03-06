package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "HELLO WORLD")
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/data/olympiads/", http.StripPrefix("/data/olympiads/", http.FileServer(http.Dir("data/olympiads"))))
	http.HandleFunc("/olympiad", olympiadHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
