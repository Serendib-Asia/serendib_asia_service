package validator

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/gofiber/fiber/v2"
)

func ValidateAgentRequest(requestID string, ctx *fiber.Ctx) (request dto.ValidateAgentsReq, errRes *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ValidateAgentRequestMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ValidateAgentRequestMethod), log.TraceMethodOutputs(commonLogFields, request, errRes)...)

	request, errRes = GenericValidator[dto.ValidateAgentsReq](requestID, ctx)
	if errRes != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(GenericValidatorMethod), log.TraceCustomError(commonLogFields, *errRes)...)
		return request, errRes
	}

	request.Authorization = ctx.Locals(constant.Authorization).(string)
	request.AuthServerID = ctx.Locals(constant.UserID).(string)
	request.User = ctx.Locals(constant.UserKey).(dto.UserResponse)
	return request, nil
}
