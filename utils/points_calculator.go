package utils

import (
	"math"
	"receipt-processor/models"
	"strconv"
	"strings"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: 1 point per alphanumeric character in retailer name
	for _, char := range receipt.Retailer {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
			points++
		}
	}

	// Rule 2: 50 points if total is a round dollar amount
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	// Rule 3: 25 points if total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for item descriptions being multiples of 3
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		if len(trimmed)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day is odd
	day := receipt.PurchaseDate[len(receipt.PurchaseDate)-2:]
	if dayInt, _ := strconv.Atoi(day); dayInt%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if time is between 2:00 PM and 4:00 PM
	hour, _ := strconv.Atoi(receipt.PurchaseTime[:2])
	if hour == 14 {
		points += 10
	}

	return points
}
