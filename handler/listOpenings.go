package handler

import (
	"net/http"

	"github.com/agmguerra/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("list operation error: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
