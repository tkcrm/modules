package pgtypeutils

import "github.com/jackc/pgx/v5/pgtype"

func EncodeText(value *string) pgtype.Text {
	v := ""
	if value != nil {
		v = *value
	}
	return pgtype.Text{
		String: v,
		Valid:  value != nil,
	}
}

func DecodeText(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}
