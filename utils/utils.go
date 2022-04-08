package utils

import (
	"math"

	"github.com/goccy/go-json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Pointer[T any](v T) *T {
	return &v
}

func Round(v float64, n float64) float64 {
	return math.Round(v*n) / n
}

func GetDefaultBool(v, d bool) bool {
	if v == false {
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

func ToProto(src interface{}, dst protoreflect.ProtoMessage) error {

	result, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := protojson.Unmarshal(result, dst); err != nil {
		return err
	}

	return nil
}

func FromProto[TDst any](src protoreflect.ProtoMessage) (TDst, error) {

	m := protojson.MarshalOptions{
		UseProtoNames: true,
	}

	var dst TDst
	bs, err := m.Marshal(src)
	if err != nil {
		return dst, err
	}

	if err := json.Unmarshal(bs, &dst); err != nil {
		return dst, err
	}

	return dst, nil
}

func ToModel(src interface{}, dst interface{}) error {

	result, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(result, dst); err != nil {
		return err
	}

	return nil
}
