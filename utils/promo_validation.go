package utils

import (
	"compress/gzip"
	"io"
	"os"
	"strings"
)

// IsValidPromoCode checks if a promo code is valid
func IsValidPromoCode(code string, files []string) bool {
	if len(code) < 8 || len(code) > 10 {
		return false
	}

	matchCount := 0
	for _, file := range files {
		if containsCode(file, code) {
			matchCount++
		}
		if matchCount >= 2 {
			return true
		}
	}

	return false
}

func containsCode(fileName, code string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		return false
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return false
	}
	defer gzipReader.Close()

	var builder strings.Builder
	// Use io.Copy to read from gzipReader into the strings.Builder
	_, err = io.Copy(&builder, gzipReader)
	if err != nil {
		return false
	}

	return strings.Contains(builder.String(), code)
}
