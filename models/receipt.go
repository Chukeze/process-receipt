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
		return errors.New("Retailer is Required");
	}

	if len(receipt.PurchaseDate) == 0 {
		return errors.New("Purchase Date is Required");
	}

	if len(receipt.PurchaseTime) == 0 {
		return errors.New("Purchase Time is Required");
	}

	if len(receipt.Items) == 0 {
		return errors.New("Items are Required");
	}

	if len(receipt.Items) < 1 {
		return errors.New("Receipt must have at least 1 item");
	}

	for _, item := range receipt.Items {
		if err := item.Validate(); err != nil {
			return err;
		}
	}

	for _, item := range receipt.Items {
		if len(item.ShortDescription) == 0 {
			return errors.New(" A ShortDescription is Required");
		}

		if len(item.Price) == 0 {
			return errors.New("Price is Required");
		}
	}

	if len(receipt.Total) == 0 {
		return errors.New("Total is Required");
	}


	// Format Validation
	if match, _ := regexp.MatchString(`^[\w\s\-&]+$`, receipt.Retailer); !match {
		return errors.New("Retailer Does not match our Expected Format");
	}

	if _, err := timeParse("2006-01-02", receipt.PurchaseDate); err != nil {
		return errors.New("Purchase Date Does not match our Expected Format, must be YYYY-MM-DD");
	}

	if _, err := timeParse("15:04", receipt.PurchaseTime); err != nil {
		return errors.New("Purchase Time Does not match our Expected Format, must be HH:MM");
	}

	if match, _ := regexp.MatchString(`^\d+\.\d{2}$`, receipt.Total); !match {
		return errors.New("Total Does not match our Expected Format, must be a decimal number with 2 decimal places");
	}

	return nil;
}

func (item *Item) Validate() error {

	// Format Validation
	if match, _ := regexp.MatchString( `^[\w\s\-]+$`, item.ShortDescription); !match {
		return errors.New("ShortDescription Does not match our Expected Format");
	}

	if match _ := regexp.MatchString(`^\d+\.\d{2}$`, item.Price); !match {
		return errors.New("Price Does not match our Expected Format, must be a decimal number with 2 decimal places");
	}

	return nil;
}
