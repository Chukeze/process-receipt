package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processor/handlers"
	"receipt-processor/models"
	"testing"
	"fmt"
)


func TestProcessReceipts(t *testing.T) {
	validReceipt := `{
		"retailer": "Walmart",
		"purchaseDate": "2021-01-01",
		"purchaseTime": "14:00",
		"items": [
			{ "shortDescription": "Coca Cola", "price": "10.00" }
		], 
		"total": "100.00"
	},`

	req := httptest.NewRequest("POST", "/receipts/process", bytes.NewBufferString(validReceipt));
	req.Header.Set("Content-Type", "application/json");
	rec := httptest.NewRecorder();

	handlers.ProcessReceipts(rec, req);

	fmt.Println("Returned Payload:", rec.Body.String()) 

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rec.Code);
	}

	var response map[string]string;
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshalling response, because parse response: %v", err);
	}

	if _, ok := response["id"]; !ok {
		t.Errorf("Expected response to have an ID");
	}
}

func TestProcessReceiptInvalidPayload(t *testing.T) {
	invalidReceipt := `{
		"retailer": "Target",
	}`

	req := httptest.NewRequest("POST", "/receipts/process", bytes.NewBufferString(invalidReceipt));
	req.Header.Set("Content-Type", "application/json");
	rec := httptest.NewRecorder();

	handlers.ProcessReceipts(rec, req);

	fmt.Println("Returned Payload:", rec.Body.String()) 

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %d", rec.Code);
	}
}

func TestGetPoints(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		},
		Total:  "35.35",
	}

	body, err := json.Marshal(receipt)
	
	id := "test01";
	handlers.ReceiptStore[id] = receipt;

	req := httptest.NewRequest("GET", "/receipts/"+id+"/points", bytes.NewReader(body));
	rec := httptest.NewRecorder();

	handlers.GetPoints(rec, req);

	fmt.Println("Returned Payload:", rec.Body.String()) 

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rec.Code);
	}

	var response map[string]int;
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshalling response, because parse response: %v", err);
	}

	if _, ok := response["points"]; !ok {
		t.Errorf("Expected response to have points");
	}
}

func TestGetPointsNonExistentID(t *testing.T) {
	req := httptest.NewRequest("GET", "/receipts/non-existent-id/points", nil);
	rec := httptest.NewRecorder();

	handlers.GetPoints(rec, req);

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %d", rec.Code);
	}
}

func TestReceiptValidation(t *testing.T) {
	// Valid receipt
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		},
	}

	if err := receipt.Validate(); err != nil {
		t.Errorf("expected receipt to be valid, got error: %v", err)
	}

	// Invalid receipt (invalid purchaseDate)
	invalidReceipt := receipt
	invalidReceipt.PurchaseDate = "01/01/2022"
	if err := invalidReceipt.Validate(); err == nil {
		t.Errorf("expected validation error for invalid purchaseDate")
	}

	// Invalid receipt (empty items)
	invalidReceipt.Items = []models.Item{}
	if err := invalidReceipt.Validate(); err == nil {
		t.Errorf("expected validation error for empty items")
	}
}

