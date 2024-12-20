package services

import (
	"math"
	"time"
	"strings"
	"strconv"
	"fetch-process-receipt/models"
	"fetch-process-receipt/utils"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0;

	points += len(strings.ReplaceAl(receipt.Retailer, " ", ""));

	total, _ := strconv.ParseFloat(receipt.Total, 64);

	if utils.Epsilon(total - float64(int(total)),0) {
		points += 50;
	}

	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5;

	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription);
		if len(description)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64);
			points += int(math.Ceil(price * 0.2));
		}
	}

	purchaseDate, _ := time.parse("2006-01-02", receipt.PurchaseDate);
	if purchaseDate.Day()%2  != 0 {
		points += 6;
	}

	purchaseTime, _ := time.Parse("15:04:05", receipt.PurchaseTime);
	IsItTwoOClock := purchaseTime.Hour() == 14 && purchaseTime.Minute() == 0 && purchaseTime.Second() == 0;
	IsItPastTwoOClock := purchaseTime.Hour() >= 14 && purchaseTime.Minute() == 0 && purchaseTime.Second() > 0;
	IsItBeforeFourOClock := purchaseTime.Hour() < 16
	if !IsItTwoOClock && IsItPastTwoOClock && IsItBeforeFourOClock {
		points += 10;
	}

	return points
}