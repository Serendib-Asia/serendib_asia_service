package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils"

	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func checkRepoError(commonLogFields []zap.Field, errorOccurredFrom string, err error) *custom.ErrorResult {
	log.Logger.Debug(log.TraceMsgFuncStart(checkRepoErrorMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(checkRepoErrorMethod), commonLogFields...)
	log.Logger.Error(log.TraceMsgErrorOccurredFrom(errorOccurredFrom), log.TraceError(commonLogFields, err)...)

	// record not found error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errRes := custom.BuildBadReqErrResult(constant.ErrRecordNotFoundCode, constant.ErrRecordNotFoundMsg, constant.Empty)
		return &errRes
	}

	// DB errors
	errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode, constant.ErrOccurredWhileRetrieving, err.Error())
	return &errRes
}

// ExtractAccessToken extracts the access token from the authorization string.
// It returns the access token and any error that occurred during extraction.
func ExtractAccessToken(commonLogFields []zap.Field, authString string) (accessToken string, err *custom.ErrorResult) {
	log.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncStart, ExtractAccessTokenMethod), commonLogFields...)
	defer log.Logger.Debug(fmt.Sprintf(constant.TraceMsgFuncEnd, ExtractAccessTokenMethod), commonLogFields...)

	// Split the Authorization header
	if authString == constant.Empty {
		log.Logger.Error(constant.ErrEmptyAuthHeaderMsg, commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.ErrEmptyAuthHeaderCode, constant.ErrEmptyAuthHeaderMsg, constant.Empty)
		return constant.Empty, &errRes
	}

	splitTokenStr := strings.Split(authString, constant.Space)
	if len(splitTokenStr) != constant.IntTwo {
		log.Logger.Error(constant.ErrOccurredWhenAuthHeaderSplitToArrayMsg, commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.ErrTokenSplitToArrayCode, constant.ErrOccurredWhenAuthHeaderSplitToArrayMsg, constant.Empty)
		return constant.Empty, &errRes
	}

	return splitTokenStr[constant.IntOne], nil
}

func buildDBError(method string, err error) *custom.ErrorResult {
	errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode,
		fmt.Sprintf(constant.ErrorOccurredWhenSelecting, method),
		err.Error())
	return &errRes
}

func buildPanicErr(method string) *custom.ErrorResult {
	errRes := custom.BuildInternalServerErrResult(constant.UnexpectedErrorCode,
		fmt.Sprintf(constant.UnexpectedErrorMessage, method),
		constant.Empty)
	return &errRes
}

func buildSelectErrFromRepo(when string, err error) *custom.ErrorResult {
	errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode,
		fmt.Sprintf(constant.ErrorOccurredWhenSelecting, when),
		err.Error())
	return &errRes
}

func buildInsertErrFromRepo(when string, err error) *custom.ErrorResult {
	errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode,
		fmt.Sprintf(constant.ErrorOccurredWhenInserting, when),
		err.Error())
	return &errRes
}

func buildDeleteErrFromRepo(when string, err error) *custom.ErrorResult {
	errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode,
		fmt.Sprintf(constant.ErrorOccurredWhenDeleting, when),
		err.Error())
	return &errRes
}

func buildUpdateErrFromRepo(when string, err error) *custom.ErrorResult {
	errRes := custom.BuildInternalServerErrResult(constant.ErrDatabaseCode,
		fmt.Sprintf(constant.ErrorOccurredWhenUpdating, when),
		err.Error())
	return &errRes
}

// APICall is a generic function that makes an API call and handles the response
func APICall[T any, E any](commonLogFields []zap.Field, request utils.Request) (*T, *E, *custom.ErrorResult) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.ExternalAPICallMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.ExternalAPICallMethod), commonLogFields...)

	result, errResult := utils.CallHTTPEndpoint(commonLogFields, request, constant.ExternalAPICallMethod)
	log.Logger.Debug(constant.AuthResponse, append(commonLogFields, zap.Any(constant.Response, result.Body))...)
	if errResult != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhen(constant.ExternalAPICallMethod), log.TraceCustomError(commonLogFields, *errResult)...)
		return nil, nil, errResult
	}

	if result.StatusCode != 200 && result.StatusCode != 201 {
		er, err := utils.JSONUnmarshal[*E](commonLogFields, result.Body)
		if err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(constant.DataUnmarshalMethod), log.TraceCustomError(commonLogFields, *err)...)
			return nil, nil, err
		}
		return nil, er, nil
	}

	sr, err := utils.JSONUnmarshal[*T](commonLogFields, result.Body)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(constant.DataUnmarshalMethod), log.TraceCustomError(commonLogFields, *err)...)
		return nil, nil, err
	}

	return sr, nil, nil
}

// GeneralServiceCall is a generic function that makes a service call and handles the response
func GeneralServiceCall[T any](commonLogFields []zap.Field, request utils.Request) (*T, *custom.ErrorResult) {
	log.Logger.Debug(log.TraceMsgFuncStart(constant.GeneralServiceCallMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(constant.GeneralServiceCallMethod), commonLogFields...)

	res, errList, err := APICall[T, []custom.ErrorInfo](commonLogFields, request)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhen(constant.ExternalAPICallMethod), log.TraceCustomError(commonLogFields, *err)...)
		return nil, err
	}

	if errList != nil {
		errRes := custom.BuildBadReqErrResultWithList(*errList...)
		log.Logger.Error(constant.ErrOccurredFromServiceCall, log.TraceCustomError(commonLogFields, errRes)...)
		return nil, &errRes
	}

	return res, nil
}

func handleDBError(commonLogFields []zap.Field, methodName, repoName string, err error) *custom.ErrorResult {
	log.Logger.Error(log.TraceMsgErrorOccurredFrom(methodName), log.TraceError(commonLogFields, err)...)
	return buildDBError(repoName, err)
}
