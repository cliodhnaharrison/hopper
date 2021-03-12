package main

import (
  "fmt"
  "net/http"
)

func ContainerRequestHandler(w http.ResponseWriter, r *http.Request) {
  imageName := r.URL.Query()["image"][0]
  userCookie, err := r.Cookie("username")
  if err != nil {
      fmt.Println("Cant find cookie")
      return
      // Some kind of handling pls
  }
  switch {
  case imageName == "ubuntu":
      username := userCookie.Value
      CreateContainer("ubuntu-wetty", username)
  default:
      fmt.Println("Image not found")
  }
}

func ContainerDestructionHandler(w http.ResponseWriter, r *http.Request) {
  userCookie, err := r.Cookie("username")
  if err != nil {
      fmt.Println("Cant find cookie")
      return
      // Some kind of handling pls
  }
  username := userCookie.Value
  KillContainer(username)
}

func main() {
  http.HandleFunc("/create", ContainerRequestHandler)
  http.HandleFunc("/destroy", ContainerDestructionHandler)
  http.Handle("/", http.FileServer(http.Dir("./static")))
  http.ListenAndServe(":3000", nil)
}
