package server

import (
	"fmt"
	"net/http"
)

const (
	HomePath = "/"
	Port     = ":8080"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Saint-Hubert doggy is coming for you !\n")
}

func Start() {
	http.HandleFunc(HomePath, home)
	http.ListenAndServe(Port, nil)
}