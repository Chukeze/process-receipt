package models

import (
	"errors"
	"regexp"
	"time"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

type Receipt struct {
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items []Item `json:"items"`
	Total string `json:"total"`
}

func (receipt *Receipt) Validate() error {

	// Validate Required Fields
	if len(receipt.Retailer) == 0 {
		return errors.New("retailer is required");
	}

	if len(receipt.PurchaseDate) == 0 {
		return errors.New("purchase date is required");
	}

	if len(receipt.PurchaseTime) == 0 {
		return errors.New("purchase time is required");
	}

	if len(receipt.Items) == 0 {
		return errors.New("items are required");
	}

	if len(receipt.Items) < 1 {
		return errors.New("receipt must have at least 1 item");
	}

	for _, item := range receipt.Items {
		if err := item.Validate(); err != nil {
			return err;
		}
	}

	for _, item := range receipt.Items {
		if len(item.ShortDescription) == 0 {
			return errors.New("a short description is required");
		}

		if len(item.Price) == 0 {
			return errors.New("price is required");
		}
	}

	if len(receipt.Total) == 0 {
		return errors.New("total is required");
	}


	// Format Validation
	if match, _ := regexp.MatchString(`^[\w\s\-&]+$`, receipt.Retailer); !match {
		return errors.New("retailer does not match our expected format");
	}

	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return errors.New("purchase date does not match our expected format, must be YYYY-MM-DD");
	}

	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return errors.New("purchase time does not match our expected format, must be HH:MM");
	}

	if match, _ := regexp.MatchString(`^\d+\.\d{2}$`, receipt.Total); !match {
		return errors.New("total does not match our expected format, must be a decimal number with 2 decimal places");
	}

	return nil;
}

func (item *Item) Validate() error {

	// Format Validation
	if match, _ := regexp.MatchString( `^[\w\s\-]+$`, item.ShortDescription); !match {
		return errors.New("short description does not match our expected format");
	}

	if match, _ := regexp.MatchString(`^\d+\.\d{2}$`, item.Price); !match {
		return errors.New("price does not match our expected format, must be a decimal number with 2 decimal places");
	}

	return nil;
}
