package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("This is from TestHandler!")
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/test", TestHandler)
  http.Handle("/", router)
  fmt.Println("Everything is OK!")
}


