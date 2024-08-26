package verifystruct

import (
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
	verifyModelMap   map[string]map[string]string
	StandardFieldMap map[string]bool
}

// VerifyStruct validates the fields in request_dict against the standardModel.
// It returns an error if any field in request_dict is invalid or if any other validation fails.
func VerifyStruct(request_dict map[string]any, standardModel any) []error {
	listErr := []error{} // Initialize the list of error messages.
	// Parse the standardModel to extract validation rules and field names.
	verifier, err_parseVerify := parseVerify(standardModel)
	if err_parseVerify != nil {
		listErr = append(listErr, err_parseVerify)
		return listErr
	}

	// check fields of request_dict that do not exist in standardmodel
	if err_fieldInvalid := utils.CheckFieldNotExistInStandardModel(request_dict, verifier.StandardFieldMap); err_fieldInvalid != nil {
		listErr = append(listErr, err_fieldInvalid...)
	}

	// Check fields of request_dict that are not allowed to be nil(null) as required by standardmodel
	if err_Requirefield := utils.CheckRequirementField(request_dict, verifier.verifyModelMap); err_Requirefield != nil {
		listErr = append(listErr, err_Requirefield...)
	}

	// Check for field values ​​in request_dict that do not match the specified type of the corresponding fields in standardmodel
	if err_InvalidFieldValue := utils.CheckValidType(request_dict, verifier.verifyModelMap); err_InvalidFieldValue != nil {
		listErr = append(listErr, err_InvalidFieldValue...)
	}

	if len(listErr) == 0 {
		return nil
	}
	return listErr
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
		verifyModelMap:   make(map[string]map[string]string),
		StandardFieldMap: make(map[string]bool),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// Check if field has pointer data type
		if field.Type.Kind() == reflect.Ptr {
			return nil, fmt.Errorf("field '%s' (tag:'%s') must not have pointer type", field.Name, field.Tag.Get("json"))
		}
		fieldName := field.Tag.Get("json")
		FieldCheckStr := field.Tag.Get("verify")
		checkField.verifyModelMap[fieldName] = parseProperties(FieldCheckStr)
		checkField.StandardFieldMap[fieldName] = true
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
