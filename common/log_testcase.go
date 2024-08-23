package common

import (
	"fmt"
	"log"
	"strings"
)

func LogValidationDetails(requestDict map[string]any, ActualErr, ExpcErr []error) {
	// Convert requestDict to a readable format
	var requestDetails []string
	for key, value := range requestDict {
		requestDetails = append(requestDetails, fmt.Sprintf("\n\t - %s:%v ", key, value))
	}
	requestDetailsStr := strings.Join(requestDetails, "")

	// Convert listErr to a readable format
	var actErrorDetails []string
	for _, err := range ActualErr {
		actErrorDetails = append(actErrorDetails, fmt.Sprintf("\n\t - %v ", err.Error()))
	}
	actErrorDetailsStr := strings.Join(actErrorDetails, "")

	// Convert ExpcErr to a readable format
	var expErrorDetails []string
	for _, err := range ExpcErr {
		expErrorDetails = append(expErrorDetails, fmt.Sprintf("\n\t - %v ", err.Error()))
	}
	expErrorDetailsStr := strings.Join(expErrorDetails, "")

	// Log the details
	log.Printf(" \n --- Validation Details ---\n + Request Data: %s\n + Actual Errors: %s\n + Expected Errors: %s \n\n",
		requestDetailsStr, actErrorDetailsStr, expErrorDetailsStr)
}
