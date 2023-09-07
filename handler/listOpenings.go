package handler

import (
	"net/http"

	"github.com/agmguerra/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Sumary List Opening
// @Description List openings opened
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("list operation error: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
