package responsebuilder

import (
	"net/http"

	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// APIResponse use to define api response
type APIResponse struct {
	_             struct{}
	Ctx           *fiber.Ctx
	Response      any
	RequestID     string
	HTTPStatus    int
	ErrorResponse custom.ErrorResult
}

// CommonSuccessResponse used to return common success response
type CommonSuccessResponse struct {
	Data any `json:"data"`
}

// BuildAPIResponse is used to build the API response
func (response *APIResponse) BuildAPIResponse() {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, response.RequestID)}
	log.Logger.Debug(constant.TraceMsgAPIResponse, commonLogFields...)

	if response.ErrorResponse.IsError {
		log.Logger.Debug(constant.TraceMsgAPIErrorResponse, append(commonLogFields, zap.Any(constant.TraceMsgReqBody, response.ErrorResponse))...)

		if response.HTTPStatus == 0 {
			response.HTTPStatus = http.StatusInternalServerError
		}

		err := response.Ctx.Status(response.HTTPStatus).JSON(response.ErrorResponse.ErrorList)
		if err != nil {
			log.Logger.Error(constant.TraceMsgAPIErrorResponse, append(commonLogFields, zap.Error(err))...)
		}
	} else {
		log.Logger.Debug(constant.TraceMsgAPISuccess, commonLogFields...)
		successResponse := CommonSuccessResponse{
			Data: response.Response,
		}

		err := response.Ctx.Status(response.HTTPStatus).JSON(successResponse)
		if err != nil {
			log.Logger.Error(constant.TraceMsgAPIErrorResponse, append(commonLogFields, zap.Error(err))...)
		}
	}
}
