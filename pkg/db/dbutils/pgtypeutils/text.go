package pgtypeutils

import "github.com/jackc/pgx/v5/pgtype"

func EncodeText(value *string) pgtype.Text {
	v := ""
	var valid bool
	if value != nil {
		v = *value
		valid = true
	}
	return pgtype.Text{
		String: v,
		Valid:  valid,
	}
}

func DecodeText(value pgtype.Text) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}
