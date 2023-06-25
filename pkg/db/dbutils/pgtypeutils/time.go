package pgtypeutils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func EncodeTime(v time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  v,
		Valid: true,
	}
}
