package handler

import "github.com/chazool/serendib_asia_service/pkg/custom"

// HandleError handles the error and returns the status code and error result.
func HandleError(errRes *custom.ErrorResult) (statusCode int, errorResult custom.ErrorResult) {
	errorResult = *errRes
	errorResult.IsError = true

	return errRes.StatusCode, errorResult
}
