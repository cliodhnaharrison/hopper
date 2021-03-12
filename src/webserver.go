package main

import (
  "fmt"
  "net/http"
  "strings"
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
      username := strings.Split(userCookie, "=")[1]
      CreateContainer("ubuntu-wetty", username)
  default:
      fmt.Println("Image not found")
  }
}

func ContainerDestructionhandler(w http.ResponseWriter, r *http.Request) {
  userCookie, err := r.Cookie("username")
  if err != nil {
      fmt.Println("Cant find cookie")
      return
      // Some kind of handling pls
  }
  username := strings.Split(userCookie, "=")[1]
  KillContainer(username)
}

func main() {
  http.HandleFunc("/create", ContainerRequestHandler)
  http.HandleFunc("/destroy", ContainerDestructionhandler)
  http.Handle("/", http.FileServer(http.Dir("./static")))
  http.ListenAndServe(":3000", nil)
}
