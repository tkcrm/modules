package utils

func GetDefaultString(v, d string) string {
	if v == "" {
		v = d
	}
	return v
}

func GetDefaultNumber[T Number](value, defaultValue T) T {
	if value == 0 {
		return defaultValue
	}
	return value
}

func ExistInArray[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
