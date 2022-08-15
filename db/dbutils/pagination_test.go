package dbutils_test

import (
	"testing"

	"github.com/tkcrm/modules/db/dbutils"
)

func Test_Pagination(t *testing.T) {
	page := uint64(1)
	records := uint64(1)

	_, _, err := dbutils.Pagination(&page, &records, dbutils.WithMaxLimit(101))
	if err != nil {
		t.Fatal(err)
	}
}
