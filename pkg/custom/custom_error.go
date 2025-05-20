package custom

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chazool/serendib_asia_service/pkg/utils/constant"
)

// ErrorResult used to define error result of the response
type ErrorResult struct {
	ErrorList  []ErrorInfo
	StatusCode int  `json:"StatusCode" example:"400"`
	IsError    bool `json:"IsError" example:"true"`
}

// Error implements error.
func (e *ErrorResult) Error() string {
	panic("unimplemented")
}

// ErrorInfo use to define error information of the ErrorResult
type ErrorInfo struct {
	ErrorCode    string `json:"ErrorCode" example:"ER0001"`
	ErrorMessage string `json:"ErrorMessage" example:"Records not found"`
	ErrorDetail  string `json:"ErrorDetail" example:"XYZ data not available in db"`
}

// BuildErrorInfo used to build error information
func BuildErrorInfo(errCode, errMessage, errDetail string) ErrorInfo {
	return ErrorInfo{
		ErrorCode:    errCode,
		ErrorMessage: errMessage,
		ErrorDetail:  errDetail,
	}
}

// BuildErrResultWithSuccessStatus used to build ErrorResult with success code
func BuildErrResultWithSuccessStatus(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusOK,
	}
}

// BuildBadReqErrResultWithList used to build ErrorResult with ErrorInfo list and bad request code
func BuildBadReqErrResultWithList(errInfo ...ErrorInfo) ErrorResult {
	return ErrorResult{
		ErrorList:  errInfo,
		IsError:    false,
		StatusCode: http.StatusBadRequest,
	}
}

// BuildBadReqErrResult used to build ErrorResult with bad request code
func BuildBadReqErrResult(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusBadRequest,
	}
}

// BuildNotFoundErrResult used to build ErrorResult with bad request code
func BuildNotFoundErrResult(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusNotFound,
	}
}

// BuildInternalServerErrResult used to build ErrorResult with internal server error code
func BuildInternalServerErrResult(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusInternalServerError,
	}
}

// GetErrorMessage use to retun error message from ErrorList
func GetErrorMessage(errorResult *ErrorResult) string {
	var errMessages []string
	for _, err := range errorResult.ErrorList {
		errMessages = append(errMessages, err.ErrorMessage)
	}

	return strings.Join(errMessages, ",")
}

// BuildForbiddenErrResult used to build ErrorResult with forbidden code
func BuildForbiddenErrResult(errCode, errMessage, errDetail string) ErrorResult {
	errList := []ErrorInfo{BuildErrorInfo(errCode, errMessage, errDetail)}

	return ErrorResult{
		ErrorList:  errList,
		IsError:    false,
		StatusCode: http.StatusForbidden,
	}
}

// BuildPanicErrResult used to build ErrorResult with internal server error code
func BuildPanicErrResult(panicMethod string) *ErrorResult {
	errRes := BuildInternalServerErrResult(constant.UnexpectedErrorCode, fmt.Sprintf(constant.UnexpectedErrorMessage, panicMethod), "")
	return &errRes
}

// CombineErrors used to combine multiple errors into one
func CombineErrors(errors []error, errorCode string) *ErrorResult {
	if len(errors) == constant.Zero {
		return nil
	}

	var errorList []ErrorInfo
	for _, err := range errors {
		errInfo := BuildErrorInfo(errorCode, err.Error(), constant.Empty)
		errorList = append(errorList, errInfo)
	}

	combinedError := ErrorResult{
		ErrorList:  errorList,
		StatusCode: http.StatusBadRequest,
		IsError:    true,
	}

	return &combinedError
}
