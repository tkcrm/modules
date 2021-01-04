package utils

// GetDefaultString ...
func GetDefaultString(v, d string) string {
	if v == "" {
		v = d
	}
	return v
}

// GetDefaultInt ...
func GetDefaultInt(v, d int) int {
	if v == 0 {
		v = d
	}
	return v
}

// GetDefaultInt32 ...
func GetDefaultInt32(v, d int32) int32 {
	if v == 0 {
		v = d
	}
	return v
}
