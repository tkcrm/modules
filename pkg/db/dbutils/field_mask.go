package dbutils

import "slices"

type FieldMask []string

func (s *FieldMask) Add(items ...string) []string {
	if s == nil {
		*s = make([]string, 0)
	}

	if len(items) == 0 {
		return *s
	}

	for _, item := range items {
		if !slices.Contains(*s, item) {
			*s = append(*s, item)
		}
	}

	return *s
}

func (s FieldMask) Items() []string {
	if s == nil {
		return FieldMask{}
	}
	return s
}

func (s FieldMask) Len() int {
	return len(s)
}
