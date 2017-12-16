package main

import (
  "fmt"
  "net/http"
  "time"
)

const (
  Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
  response := "The time is now " + time.Now().String()
  fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "static.html")
}

func main() {
  http.HandleFunc("/static", serveStatic)
  http.HandleFunc("/", serveDynamic)
  fmt.Printf("Server starts at port localhost:8080\n")
  fmt.Printf("Static files are at localhost:/static\n")
  http.ListenAndServe(Port, nil)
}

