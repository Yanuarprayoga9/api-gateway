package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// middleware merchant
func merchantMidlle (next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		if author != "merchant" {
			w.Write([]byte("anda tidak memiliki akses"))
		}
		next.ServeHTTP(w,r)
	}
}

// middleware admin
func superAdmin (next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		if author != "su-admin" {
			w.Write([]byte("anda tidak memiliki akses"))
		}
		next.ServeHTTP(w,r)
	}
}

func getMerchant(w http.ResponseWriter,r *http.Request) {
	resp , err := http.Get("http://localhost:8000/get-merchant")
	if err != nil {
		w.Write([]byte("err at get merchant"))
	}
	data,_ := ioutil.ReadAll(resp.Body)
	json.NewEncoder(w).Encode(data);
}

func getAllToko(w http.ResponseWriter,r *http.Request) {
	resp , err := http.Get("http://localhost:9000/get-all-toko")
	if err != nil {
		w.Write([]byte("err at get merchant"))
	}
	data,_ := ioutil.ReadAll(resp.Body)
	json.NewEncoder(w).Encode(data);
}
func main () {
	mux := http.NewServeMux()

	mux.HandleFunc("/merchants",merchantMidlle(getMerchant))
	mux.HandleFunc("/toko",superAdmin(getMerchant))
	fmt.Print("server running")
	http.ListenAndServe(":5000",mux)
}