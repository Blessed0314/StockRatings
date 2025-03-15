package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertPrice(priceStr string) (float64, error) {
	cleanStr := strings.ReplaceAll(priceStr, "$", "")
	cleanStr = strings.ReplaceAll(cleanStr, ",", "")

	// Convertir a float64
	return strconv.ParseFloat(cleanStr, 64)
}

func ConvertToString(price float64) string {
	return fmt.Sprintf("$%.2f", price)
}