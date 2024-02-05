package dbutils

type PaginatorOption func(p *Paginator)

func WithMaxVisibleItems(v int) PaginatorOption {
	return func(p *Paginator) {
		p.maxPages = v
	}
}
