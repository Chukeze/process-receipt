package handlers

import (
	"encoding/json"
	"net/http"
	"fetch-process-receipt/models"
	"fetch-process-receipt/services"
	"fetch-process-receipt/utils"

	"github.com/gorilla/mux"
)

var ReceiptStore = make(map[string]models.Receipt);

func ProcessReceipt( writer http.ResponseWriter, request * http.Request) {

	var receipt models.Receipt;
	//ToDo: Validating the request
	if err := json.NewDecoder(request.Body).Decode(&receipt); err != nil{
		http.Error(writer, "The payload you're sending is causing an error", http.StatusBadRequest);
		return;
	}

	if err := receipt.Validate(); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest);
		return;
	}

	id := utils.GenerateID();
	ReceiptStore[id] = receipt;

	writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(writer).Encode(map[string]string{"id": id});
}

func GetPoints(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"];

	receipt, exists := ReceiptStore[id];
	if !exists {
		json.NewEncoder(w).Encode(map[string]string{"error": "Receipt not found"});
		writer.WriteHeader(http.StatusNotFound)
		return;
	}

	points := services.CalculatePoints(receipt);

	writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(writer).Encode(map[string]int{"points": points});
}
