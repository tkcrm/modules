package dbutils

import (
	"fmt"
	"strconv"

	"github.com/huandu/go-sqlbuilder"
)

type SearchFieldType int

const (
	SearchFieldTypeLikeCaseInsensitive SearchFieldType = iota
	SearchFieldTypeLikeCaseSensitive
	SearchFieldTypeExactMatch
	SearchFieldTypeIntExactMatch
)

type SearchField struct {
	Field string
	Type  SearchFieldType
}

func BuildSearch(sb *sqlbuilder.SelectBuilder, search string, fields ...SearchField) {
	var whereExpr []string
	for _, item := range fields {
		switch item.Type {
		case SearchFieldTypeLikeCaseInsensitive:
			whereExpr = append(whereExpr, fmt.Sprintf("LOWER(%s) LIKE LOWER(%s)", item.Field, sb.Var(GetLikeVal(search))))
		case SearchFieldTypeLikeCaseSensitive:
			whereExpr = append(whereExpr, sb.Like(item.Field, GetLikeVal(search)))
		case SearchFieldTypeExactMatch:
			whereExpr = append(whereExpr, sb.Equal(item.Field, search))
		case SearchFieldTypeIntExactMatch:
			if _, err := strconv.ParseInt(search, 10, 64); err == nil {
				whereExpr = append(whereExpr, sb.Equal(item.Field, search))
			}
		}
	}

	if len(whereExpr) > 0 {
		sb.Where(sb.Or(whereExpr...))
	}
}
