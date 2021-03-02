package main

import (
  "fmt"
  "net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query()["image"][0]
  fmt.Println(imageName)
	//CreateContainer(imageName)
}


func main() {
  http.HandleFunc("/create", PostHandler)
  http.Handle("/", http.FileServer(http.Dir("./static")))
  http.ListenAndServe(":3000", nil)
}
