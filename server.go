package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        http.ServeFile(w, r, r.URL.Path[1:])
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}