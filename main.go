package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter();

	router.HandleFunc("/receipts/process")
	router.HandleFunc("receipts/{id}/points")
}