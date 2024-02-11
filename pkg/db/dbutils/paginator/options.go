package paginator

type Option func(p *Paginator)

func WithMaxVisibleItems(v int) Option {
	return func(p *Paginator) {
		p.maxPages = v
	}
}
