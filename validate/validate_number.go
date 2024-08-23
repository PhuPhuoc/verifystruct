package validate

import "reflect"

func IsNumber(value any) bool {
	numberTypes := []reflect.Kind{
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
	}

	valueKind := reflect.ValueOf(value).Kind()
	for _, t := range numberTypes {
		if valueKind == t {
			return true
		}
	}
	return false
}
