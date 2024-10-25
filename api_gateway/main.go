package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// middleware merchant
func merchantMidlle (next http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		if author != "su-admin" {
			w.WriteHeader(http.StatusForbidden) 
			w.Write([]byte("Anda tidak memiliki akses"))
			return 
		}

		next.ServeHTTP(w,r)
	}
}

func superAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		if author != "su-admin" {
			w.WriteHeader(http.StatusForbidden) 
			w.Write([]byte("Anda tidak memiliki akses"))
			return 
		}

		next.ServeHTTP(w, r)
	}
}

func getMerchant(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8000/get-merchant")
	if err != nil {
		http.Error(w, "Error fetching merchant", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close() 

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error fetching merchant: "+resp.Status, resp.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set header content type
	w.WriteHeader(http.StatusOK) // Mengatur status code 200
	w.Write(data)
}

func getAllToko(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:9000/get-all-toko")
	if err != nil {
		http.Error(w, "Error fetching all toko", http.StatusInternalServerError) // Mengatur status code 500
		return
	}
	defer resp.Body.Close() // Pastikan untuk menutup body setelah dibaca

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error fetching all toko: "+resp.Status, resp.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set header content type
	w.WriteHeader(http.StatusOK) // Mengatur status code 200
	w.Write(data)
}

func main () {
	mux := http.NewServeMux()

	mux.HandleFunc("/merchants",merchantMidlle(getMerchant))
	mux.HandleFunc("/toko",superAdmin(getMerchant))
	fmt.Print("server running")
	http.ListenAndServe(":5000",mux)
}