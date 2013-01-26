package main
import (
    "net/http"
    "io"
    "log"
    "os"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "HELLO WORLD")
}

func main() {
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        log.Fatal("ListenAndServe", err)
    }
}

