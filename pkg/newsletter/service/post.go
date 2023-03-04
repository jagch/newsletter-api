package service

import "git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"

func (s *service) Post(subscription *newsletter.Subscription) error {
	return s.repo.Create(subscription)
}
