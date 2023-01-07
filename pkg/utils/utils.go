package utils

import (
	"math"
	"reflect"
	"runtime"
	"strings"

	"github.com/goccy/go-json"
)

func Pointer[T any](v T) *T {
	return &v
}

func Round(v float64, n float64) float64 {
	return math.Round(v*n) / n
}

func GetDefaultBool(v, d bool) bool {
	if !v {
		v = d
	}
	return v
}

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

func JsonToStruct(src interface{}, dst interface{}) error {
	result, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(result, dst)
}

// GetFunctionName return caller function name
func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	funcName := strs[len(strs)-1]
	strs = strings.Split(funcName, "-")
	return strs[0]
}
