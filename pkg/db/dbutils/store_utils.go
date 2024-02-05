package dbutils

type FindResponse[T any] struct {
	Items []T `json:"items"`
}

type FindResponseWithPaginator[T any] struct {
	Items     []T       `json:"items"`
	Pagniator Pagniator `json:"pagination"`
}

type PageParams struct {
	Page              *uint64 `query:"page" params:"page" json:"page" validate:"omitempty,gte=1"`
	PageSize          *uint64 `query:"page_size" params:"page_size" json:"page_size" validate:"omitempty,gte=1"`
	DisablePagination bool    `query:"disable_pagination" params:"disbale_pagination" json:"disable_pagination" default:"false"`
}

type OrderByParams[TOrderBy ValidableString] struct {
	OrderBy            TOrderBy `query:"order_by" params:"order_by" json:"order_by"`
	QueryIsAscOrdering bool     `query:"is_asc_ordering" params:"is_asc_ordering" json:"is_asc_ordering"`
}

type SearchParams struct {
	Search string `query:"search" params:"search" json:"search"`
}
