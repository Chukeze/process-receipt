package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter();

	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST");
	router.HandleFunc("receipts/{id}/points", handlers.GetPoints).Methods("GET");

	log.Println("Server started on port 8080");
	log.Fatal(http.ListenAndServe(":8080", router));
}