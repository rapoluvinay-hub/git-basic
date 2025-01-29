package main

import (
	"log"
	"net/http"n

	"github.com/gorilla/mux"
	"receipt-processor/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handlers.ProcessReceipts).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
