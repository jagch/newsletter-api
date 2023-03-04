package repository

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/utils"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	limit int,
	offset int,
) ([]*newsletter.Subscription, error) {
	data := r.data

	if blogID.String() == uuid.Nil.String() && userID.String() == uuid.Nil.String() && len(interests) == 0 {
		//return all data
		return mapperArrayOfSubscriptionDBModelToArrayOfNewsletterSubscription(data), nil
	}

	return nil, nil
}

func mapperArrayOfSubscriptionDBModelToArrayOfNewsletterSubscription(arrayOfSubscriptionDBModel []*subscriptionDBModel) []*newsletter.Subscription {
	var newsletterSubscriptions []*newsletter.Subscription

	i := 0
	for i < len(arrayOfSubscriptionDBModel) {
		newsletterSubscription := new(newsletter.Subscription)
		newsletterSubscription.UserID, _ = uuid.Parse(arrayOfSubscriptionDBModel[i].UserID)
		newsletterSubscription.BlogID, _ = uuid.Parse(arrayOfSubscriptionDBModel[i].BlogID)
		newsletterSubscription.Interests = utils.ArraysOfStringToNewsletterInterest(arrayOfSubscriptionDBModel[i].Interests)

		newsletterSubscriptions = append(newsletterSubscriptions, newsletterSubscription)

		i++
	}

	return newsletterSubscriptions
}
