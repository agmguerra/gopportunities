package handler

import (
	"fmt"
	"net/http"

	"github.com/agmguerra/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "applicaation/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "applicaation/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type BaseOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type CreateOpeningResponse struct {
	BaseOpeningResponse
}

type DeleteOpeningResponse struct {
	BaseOpeningResponse
}

type ListOpeningResponse struct {
	BaseOpeningResponse
}

type ShowOpeningResponse struct {
	BaseOpeningResponse
}

type UpdateOpeningResponse struct {
	BaseOpeningResponse
}
