package dbutils

import (
	"math"
)

var maxVisiblePages = 5

type Pagniator struct {
	page       int
	pageSize   int
	totalItems int

	totalPages int
	maxPages   int
}

func New(page, pageSize, totalItems int) Pagniator {
	total := totalItems / pageSize
	if totalItems%pageSize > 0 {
		total++
	}
	return Pagniator{
		page:       page,
		pageSize:   pageSize,
		totalItems: totalItems,
		totalPages: total,
		maxPages:   maxVisiblePages,
	}
}

func (p Pagniator) CurrentPage() int {
	return p.page
}

func (p Pagniator) PageSize() int {
	return p.pageSize
}

func (p Pagniator) TotalItems() int {
	return p.totalItems
}

func (p Pagniator) HasPrevious() bool {
	return p.page > 1
}

func (p Pagniator) HasNext() bool {
	return p.page < p.totalPages
}

func (p Pagniator) HasPages() bool {
	return p.totalPages > 1
}

func (p Pagniator) Previous() *int {
	if !p.HasPrevious() {
		return nil
	}
	res := p.page - 1
	return &res
}

func (p Pagniator) Next() *int {
	if !p.HasNext() {
		return nil
	}
	res := p.page + 1
	return &res
}

func (p Pagniator) First() int {
	return 1
}

func (p Pagniator) Last() int {
	return p.TotalPages()
}

func (p Pagniator) TotalPages() int {
	return p.totalPages
}

func (p Pagniator) CurrentPageFromItem() int {
	if p.page == 1 {
		return 1
	}
	return (p.page-1)*p.pageSize + 1
}

func (p Pagniator) CurrentPageToItem() int {
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

func (p Pagniator) DrawPages() []*DrawPagesItem {
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
