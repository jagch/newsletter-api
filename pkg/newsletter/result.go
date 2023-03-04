package newsletter

import (
	"context"
	"math"
)

type Result[T any] struct {
	Total int
	Pages int
	Page  Page[T]
}

func (r *Result[T]) Get(ctx context.Context, data []T) {
	pageRequest, maxPageSize := ctx.Value("page").(int), ctx.Value("maxPageSize").(int)
	offset, limit := maxPageSize*(pageRequest-1), maxPageSize

	var page = new(Page[T])

	r.Pages = countPages(maxPageSize, len(data))
	r.Total = len(data)
	r.Page = page.New(data, limit, offset)
}

func countPages(maxPageSize int, totalElements int) int {
	if totalElements == 0 {
		return 0
	}

	if totalElements <= maxPageSize {
		return 1
	}
	return int(
		math.Round(
			float64(totalElements) / float64(maxPageSize),
		),
	)
}
