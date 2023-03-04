package repository

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

func (r *repository) Create(subscription *newsletter.Subscription) error {
	r.data = append(r.data, mapperSubscriptionToSubscriptionDBModel(subscription))

	return nil
}
