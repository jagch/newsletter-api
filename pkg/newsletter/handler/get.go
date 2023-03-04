package handler

import (
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/repository"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/service"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// nolint:lll // godoc
// Get godoc
// @Summary      Read subscriptions
// @Tags         subscriptions
// @Router       /subscriptions [get]
// @Param        page	        query  int		 true   "Result page"                                   example(1)
// @Param        maxPageSize	query  int		 true   "Max page size"                                  example(10)
// @Param        userId	        query  string	 false  "User ID"                                        example(1)
// @Param        blogId	        query  string	 false  "Blog ID"                                        example(1)
// @Param        interests	    query  []string  false  "Interests"                                      example(["tech","sports"])
// @Produce      json
// @Success      200  {array}  handler.ResponseDoc
// nolint:gocyclo //error checking branches
func (h *handler) Get(ctx *gin.Context) {
	userID, blogID := uuid.UUID{}, uuid.UUID{}
	filter, pagination := &Filter{}, &Pagination{}

	var err error

	if err := ctx.BindQuery(filter); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if filter.UserID != "" {
		userID, err = uuid.Parse(filter.UserID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}
	}

	if filter.BlogID != "" {
		blogID, err = uuid.Parse(filter.BlogID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}
	}

	filter.Interests = ctx.Request.URL.Query()["interests"]

	if err := ctx.BindQuery(pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Set("page", pagination.Page)
	ctx.Set("maxPageSize", pagination.MaxPageSize)

	newsletterService := service.Must(repository.Must())

	results, err := newsletterService.Get(ctx, userID, blogID, utils.ArraysOfStringToNewsletterInterest(filter.Interests))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Filter:     *filter,
		Pagination: pagination.New(results.Total),
		Results:    results,
	})

}
