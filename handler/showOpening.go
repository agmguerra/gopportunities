package handler

import (
	"fmt"
	"net/http"

	"github.com/agmguerra/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Errorf("show operation error: %v", errParamIsRequired("id", "queryParameter").Error())
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	// FInd Opening
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("show operation error: %v", err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id [%s] not found", id))
		return
	}
	sendSuccess(ctx, "show-opening", opening)
}
