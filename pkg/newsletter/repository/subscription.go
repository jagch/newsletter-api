package repository

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/utils"
)

type subscriptionDBModel struct {
	UserID    string
	BlogID    string
	Interests []string
}

func mapperSubscriptionToSubscriptionDBModel(subscription *newsletter.Subscription) *subscriptionDBModel {

	return &subscriptionDBModel{
		UserID:    subscription.UserID.String(),
		BlogID:    subscription.BlogID.String(),
		Interests: utils.SubscriptionInterestsToArrayOfString(subscription),
	}
}
