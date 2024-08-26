package validate

import "reflect"

func IsNumber(value any, min, max *int) bool {
	numberTypes := []reflect.Kind{
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
	}

	valueKind := reflect.ValueOf(value).Kind()
	for _, t := range numberTypes {
		if valueKind == t {
			v := reflect.ValueOf(value)
			switch valueKind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intVal := v.Int()
				if (min != nil && intVal < int64(*min)) || (max != nil && intVal > int64(*max)) {
					return false
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				uintVal := v.Uint()
				if (min != nil && uintVal < uint64(*min)) || (max != nil && uintVal > uint64(*max)) {
					return false
				}
			case reflect.Float32, reflect.Float64:
				floatVal := v.Float()
				if (min != nil && floatVal < float64(*min)) || (max != nil && floatVal > float64(*max)) {
					return false
				}
			}
			return true
		}
	}

	return false
}
