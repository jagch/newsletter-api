package handler

import "git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"

type Response struct {
	Filter     Filter                                       `json:"filter"`
	Pagination *Pagination                                  `json:"pagination"`
	Results    *newsletter.Result[*newsletter.Subscription] `json:"results"`
}

type Result struct {
	UserID    string   `json:"userId"`
	BlogID    string   `json:"blogId"`
	Interests []string `json:"interests"`
}
