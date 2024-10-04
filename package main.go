package main

import "net/http"

func Ping(rw http.ResponseWriter, r *http.Request) {
	result := "pong"
	rw.Write([]byte(result))
}

func Hello(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	rw.Write([]byte(name))
}

func main() {
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":8080", nil)
}
