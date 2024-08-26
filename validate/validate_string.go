package validate

func IsString(value any, min, max *int) bool {
	if v, ok := value.(string); ok {
		if (min != nil && len(v) < *min) || (max != nil && len(v) > *max) {
			return false
		}
		return true
	}
	return false
}
