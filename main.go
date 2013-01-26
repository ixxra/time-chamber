package main
import (
    "net/http"
    "io"
    "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "HELLO WORLD")
}

func main() {
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe", err)
    }
}

