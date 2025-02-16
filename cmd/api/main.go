package main

import (
    "log"
    "net/http"
)

func main() {
    log.Println("Starting Docker AI Platform API...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
