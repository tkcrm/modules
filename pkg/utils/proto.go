package utils

import (
	"github.com/goccy/go-json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type IProtoMessage interface {
	ToProtoMessage() (protoreflect.ProtoMessage, error)
}

func ToProtoSlice[TSrc IProtoMessage, TDst protoreflect.ProtoMessage](src []TSrc) ([]TDst, error) {
	res := make([]TDst, 0, len(src))
	for _, item := range src {
		pbitem, err := item.ToProtoMessage()
		if err != nil {
			return nil, err
		}
		res = append(res, pbitem.(TDst))
	}

	return res, nil
}

func ToProto(src any, dst protoreflect.ProtoMessage) error {
	result, err := json.Marshal(src)
	if err != nil {
		return err
	}

	u := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	if err := u.Unmarshal(result, dst); err != nil {
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
