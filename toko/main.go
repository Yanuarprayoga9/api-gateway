package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getAllToko(w http.ResponseWriter,R *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode([]map[string]string{
		{
			"nama":"yanuar shop",
			"jumlah_product":"10",
		},
		{
			"nama":"syifa shop",
			"jumlah_product":"10",
		},
		{
			"nama":"ahmad shop",
			"jumlah_product":"10",
		},
	})
}
func getDetailToko(w http.ResponseWriter,r *http.Request) {
	// we already get data from db to get all toko
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"nama":"yanuar shop",
		"jumlah_product":"10",
	})
}
func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/get-detail-toko",getDetailToko)
	mux.HandleFunc("/get-all-toko",getAllToko)
	fmt.Println("server running")
	log.Fatal(http.ListenAndServe(":9000",mux))
}