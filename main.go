package main

import (
    "fmt"
    "net/http"
	"os"
)

func main() {
    http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
        fmt.Fprint(rw, "This is mycodesmells/golang-examples server from Heroku!")
    })
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)}