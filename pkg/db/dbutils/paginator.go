package dbutils

import (
	"math"
)

var maxVisiblePages = 5

type Paginator struct {
	page       int
	pageSize   int
	totalItems int

	totalPages int
	maxPages   int
}

func New(page, pageSize, totalItems int, opts ...PaginatorOption) Paginator {
	total := totalItems / pageSize
	if totalItems%pageSize > 0 {
		total++
	}

	p := Paginator{
		page:       page,
		pageSize:   pageSize,
		totalItems: totalItems,
		totalPages: total,
		maxPages:   maxVisiblePages,
	}

	for _, opt := range opts {
		opt(&p)
	}

	return p
}

func (p Paginator) CurrentPage() int {
	return p.page
}

func (p Paginator) PageSize() int {
	return p.pageSize
}

func (p Paginator) TotalItems() int {
	return p.totalItems
}

func (p Paginator) HasPrevious() bool {
	return p.page > 1
}

func (p Paginator) HasNext() bool {
	return p.page < p.totalPages
}

func (p Paginator) HasPages() bool {
	return p.totalPages > 1
}

func (p Paginator) Previous() *int {
	if !p.HasPrevious() {
		return nil
	}
	res := p.page - 1
	return &res
}

func (p Paginator) Next() *int {
	if !p.HasNext() {
		return nil
	}
	res := p.page + 1
	return &res
}

func (p Paginator) First() int {
	return 1
}

func (p Paginator) Last() int {
	return p.TotalPages()
}

func (p Paginator) TotalPages() int {
	return p.totalPages
}

func (p Paginator) CurrentPageFromItem() int {
	if p.page == 1 {
		return 1
	}
	return (p.page-1)*p.pageSize + 1
}

func (p Paginator) CurrentPageToItem() int {
	if p.page == 1 {
		res := p.pageSize
		if res > p.totalItems {
			res = p.totalItems
		}
		return res
	}

	res := p.page * p.pageSize
	if res > p.totalItems {
		res = p.totalItems
	}

	return res
}

type DrawPagesItem struct {
	IsActive   bool
	PageNumber int
	ClassName  string
}

func (p Paginator) DrawPages() []*DrawPagesItem {
	half := int(math.Floor(float64(p.maxPages) / 2))
	lastPage := p.maxPages

	if p.page+half > p.totalPages {
		lastPage = p.totalPages
	} else if p.page > half {
		lastPage = p.page + half
	}

	from := int(math.Max(0, float64(lastPage-p.maxPages)))
	arrLength := int(math.Min(float64(p.totalPages), float64(p.maxPages)))

	arr := make([]*DrawPagesItem, arrLength)
	for i := 1; i <= arrLength; i++ {
		arr[i-1] = &DrawPagesItem{
			PageNumber: i + from,
			IsActive:   i+from == p.page,
		}
	}

	return arr
}
