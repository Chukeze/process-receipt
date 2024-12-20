package handlers

import (
	"encoding/json"
	"net/http"
)

func ProcessReceipt( writer http.ResponseWriter, request * http.Request) {

	//ToDo: Validating the request
	if {}

	writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(writer).Encode();//to Do: encode the data store response
}

