package handler

import (
	"fmt"
	"math"
)

type Pagination struct {
	Page             int    `json:"page" form:"page"`
	NumberOfPages    int    `json:"numberOfPages"`
	PaginationString string `json:"paginationString"`
	MaxPageSize      int    `json:"maxPageSize" form:"maxPageSize"`
	TotalElements    int    `json:"totalElements"`
}

func (p *Pagination) New(totalElements int) *Pagination {
	var numberOfPages = calculateNumberOfPages(totalElements, p.MaxPageSize)
	return &Pagination{
		Page:             p.Page,
		NumberOfPages:    numberOfPages,
		PaginationString: fmt.Sprintf("%v / %v", p.Page, numberOfPages),
		MaxPageSize:      p.MaxPageSize,
		TotalElements:    totalElements,
	}
}

func calculateNumberOfPages(totalElements, maxPageSize int) int {
	if totalElements == 0 {
		return 0
	}

	if totalElements <= maxPageSize {
		return 1
	}

	return int(math.Ceil(float64(totalElements) / float64(maxPageSize)))
}
