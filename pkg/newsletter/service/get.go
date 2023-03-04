package service

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (s *service) Get(
	ctx context.Context,
	UserID uuid.UUID,
	BlogID uuid.UUID,
	Interests []newsletter.Interest,
) (*newsletter.Result[*newsletter.Subscription], error) {
	result := &newsletter.Result[*newsletter.Subscription]{}

	var err error

	page, maxPageSize := ctx.Value("page").(int), ctx.Value("maxPageSize").(int)
	offset, limit := maxPageSize*(page-1), maxPageSize

	subscriptions, err := s.repo.Search(ctx, UserID, BlogID, Interests, limit, offset)
	if err != nil {
		return nil, err
	}

	result.Get(ctx, subscriptions)

	return result, nil
}
