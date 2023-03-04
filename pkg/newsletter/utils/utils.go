package utils

import "git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"

func SubscriptionInterestsToArrayOfString(subscription *newsletter.Subscription) []string {
	interests := make([]string, len(subscription.Interests))
	for i, v := range subscription.Interests {
		interests[i] = string(v)
	}

	return interests
}

func ArraysOfStringToNewsletterInterest(filterInterests []string) []newsletter.Interest {
	newsLetterInterests := make([]newsletter.Interest, len(filterInterests))
	for i, v := range filterInterests {
		newsLetterInterests[i] = newsletter.Interest(v)
	}

	return newsLetterInterests
}
