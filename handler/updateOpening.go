package handler

import (
	"fmt"
	"net/http"

	"github.com/agmguerra/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func populateOpeningForUpdate(o *schemas.Opening, r *UpdateOpeningRequest) {
	if r.Role != "" {
		o.Role = r.Role
	}
	if r.Company != "" {
		o.Company = r.Company
	}
	if r.Location != "" {
		o.Location = r.Location
	}
	if r.Link != "" {
		o.Link = r.Link
	}
	if r.Remote != nil {
		o.Remote = *r.Remote
	}
	if r.Salary >= 0 {
		o.Salary = r.Salary
	}
}

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	// bind json request
	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("json bind error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// validate fields, If any field is provided, validation is truthy
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		logger.Errorf("update operation error: %v", errParamIsRequired("id", "queryParameter").Error())
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// FInd Opening
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("update operation error: %v", err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id [%s] not found", id))
		return
	}
	populateOpeningForUpdate(&opening, &request)
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("update operation error: %v", err)
		sendError(ctx, http.StatusNotModified, fmt.Sprintf("opening with id [%s] not found", id))
		return
	}
	sendSuccess(ctx, "update-opening", opening)
}
