package main

import (
	"web/server"
	"log"
	"net/http"
)

const (
	infoMsg = `This is simple web server written in Go
	
Maintainer: Zdravko Chorlev`
)

func main() {
	log.Println("Starting server ...")

	http.HandleFunc("/", info)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func info(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(infoMsg))
	if err != nil {
		log.Fatal(err)
	}
}
