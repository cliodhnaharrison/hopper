package main

import (
  "fmt"
  "net/http"
)

func ContainerRequestHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query()["image"][0]
  username := r.URL.Query()["username"][0]
  switch {
  case imageName == "ubuntu":
      CreateContainer("ubuntu-wetty", username)
      http.Redirect(w, r, "/"+ username, http.StatusSeeOther)
  default:
      fmt.Println("Image not found")
  }
}

func main() {
  http.HandleFunc("/create", ContainerRequestHandler)
  http.Handle("/", http.FileServer(http.Dir("./static")))
  http.ListenAndServe(":3000", nil)
}
