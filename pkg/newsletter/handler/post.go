package handler

import (
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/repository"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/service"
	"github.com/gin-gonic/gin"
)

// nolint:lll // godoc
// Post godoc
// @Summary Create subscriptions
// @Tags subscriptions
// @Description Create subscriptions
// @Accept json
// @Router       /subscriptions [post]
// @param Subcription body newsletter.Subscription true  "Create subscriptions"
// @Produce      json
// @Success      200  {string} Success
// nolint:gocyclo //error checking branches
func (h *handler) Post(ctx *gin.Context) {
	subscription := newsletter.Subscription{}

	newsletterService := service.Must(repository.Must())

	if err := ctx.BindJSON(&subscription); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := newsletterService.Post(&subscription); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
