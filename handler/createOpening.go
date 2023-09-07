package handler

import (
	"net/http"

	"github.com/agmguerra/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Sumary Create Opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	// bind json request
	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("json bind error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// validate fields
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// create opening
	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	sendSuccess(ctx, "create-opening", opening)
}
