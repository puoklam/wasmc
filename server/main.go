package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../assets"))
	if err := http.ListenAndServe(":9000", fs); err != nil {
		log.Fatalln(err)
	}
}
