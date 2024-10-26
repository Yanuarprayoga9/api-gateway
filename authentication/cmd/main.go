package main

import "net/http"

func main() {
	var mux = http.NewServeMux();
	mux.HandleFunc("/login")
}