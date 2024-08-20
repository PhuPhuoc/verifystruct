package verifystruct

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/PhuPhuoc/verifystruct/utils"
)

// verify struct holds the validation requirements, the map of expected fields, and error messages.
// - verifyRequirementMap: maps field names to their validation rules.
// - listFieldMap: a map that tracks the existence of fields in the standardModel (for fast lookup).
// - listMessError: a slice of error messages accumulated during validation.
type verify struct {
	verifyRequirementMap map[string]map[string]string
	listFieldMap         map[string]bool
	listMessError        []string
}

// returnErrorMessage consolidates all error messages into a single string and logs them.
// If there are errors, it logs the concatenated error message string and returns an error.
func (v *verify) returnErrorMessage() error {
	if len(v.listMessError) > 0 {
		var message string
		for _, err_mess := range v.listMessError {
			if len(message) > 0 {
				// Concatenate all error messages into one string, separated by "; ".
				message += "; "
			}
			message += err_mess
		}
		//log.Println(message)
		//return nil // Return an error here if needed (currently returning nil).
		return errors.New(message)
	}
	// Return nil if there are no errors.
	return nil
}

// VerifyStruct validates the fields in request_dict against the standardModel.
// It returns an error if any field in request_dict is invalid or if any other validation fails.
func VerifyStruct(request_dict map[string]any, standardModel any) error {
	// Parse the standardModel to extract validation rules and field names.
	verifier, err_parseVerify := parseVerify(standardModel)
	if err_parseVerify != nil {
		return err_parseVerify
	}
	verifier.listMessError = []string{} // Initialize the list of error messages.

	// Check for fields in request_dict that do not exist in the standardModel.
	if err_fieldInvalid := utils.CheckFieldNotExistInStandardModel(request_dict, verifier.listFieldMap); err_fieldInvalid != nil {
		verifier.listMessError = append(verifier.listMessError, err_fieldInvalid...)
	}

	// other validation func

	return verifier.returnErrorMessage() // Return any accumulated error messages.
}

// parseVerify reflects on the standardModel to extract field names and their associated
// validation rules (from struct tags). It returns a pointer to a verify struct containing this metadata.
func parseVerify(standardModel any) (*verify, error) {
	t := reflect.TypeOf(standardModel)
	// Ensure that the standardModel is a struct; return an error if it's not.
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, but got %s", t.Kind())
	}
	checkField := verify{
		verifyRequirementMap: make(map[string]map[string]string),
		listFieldMap:         make(map[string]bool),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := strings.ToLower(field.Name)
		FieldCheckStr := field.Tag.Get("verify")
		checkField.verifyRequirementMap[fieldName] = parseProperties(FieldCheckStr)
		checkField.listFieldMap[fieldName] = true

	}
	return &checkField, nil
}

// parseProperties parses the "verify" tag of a struct field to extract validation rules.
// It returns a map where each key is a validation rule name and the value is the corresponding rule parameter.
func parseProperties(tagVerify string) map[string]string {
	result := make(map[string]string)
	for _, pair := range strings.Split(tagVerify, ",") {
		parts := strings.Split(pair, "=")
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}
