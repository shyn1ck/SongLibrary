package handlers

import (
	"SongLibrary/logger"
	"SongLibrary/utils/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Message: message}
}

func handleError(c *gin.Context, err error) {
	var statusCode int
	var errorResponse ErrorResponse

	switch {
	case errors.Is(err, errs.ErrInvalidToken):
		statusCode = http.StatusForbidden
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrNoReportsFound):
		statusCode = http.StatusNotFound
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrUserNotFound):
		statusCode = http.StatusNotFound
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrSongNotFound):
		statusCode = http.StatusNotFound
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrAlbumNotFound):
		statusCode = http.StatusNotFound
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrArtistNotFound):
		statusCode = http.StatusNotFound
		errorResponse = NewErrorResponse(err.Error())

	case errors.Is(err, errs.ErrPermissionDenied):
		statusCode = http.StatusForbidden
		errorResponse = NewErrorResponse(err.Error())

	default:
		logger.Error.Printf("Standard error occurred: %v", err)
		statusCode = http.StatusInternalServerError
		errorResponse = NewErrorResponse(errs.ErrSomethingWentWrong.Error())
	}

	c.JSON(statusCode, errorResponse)
}
