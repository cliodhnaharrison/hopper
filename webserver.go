package main

import (
  "fmt"
  "net/http"
)

func ContainerRequestHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query()["image"][0]
  fmt.Println(imageName)
  switch {
  case imageName == "ubuntu":
      CreateContainer("ubuntu-wetty")
  default:
      fmt.Println("Image not found")
  }
}

func main() {
  http.HandleFunc("/create", ContainerRequestHandler)
  http.Handle("/", http.FileServer(http.Dir("./static")))
  http.ListenAndServe(":3000", nil)
}
