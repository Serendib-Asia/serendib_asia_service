package validator

import (
	"github.com/chazool/serendib_asia_service/pkg/log"

	"github.com/chazool/serendib_asia_service/pkg/config"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
	"go.uber.org/zap"
)

var (
	validate         *validator.Validate
	trans            ut.Translator
	generalErrorCode map[string]string
	err              error
	// AllowedPriorities holds the list of allowed priorities.
	AllowedPriorities []string
	boolValues        = []string{"true", "false", "TRUE", "FALSE", "1", "0", "t", "f"}
	statusValues      = []string{"true", "false", ""}
	periodValues      = []string{
		string(constant.TimePeriodYearly),
		string(constant.TimePeriodQuarterly),
		string(constant.TimePeriodMonthly),
		string(constant.TimePeriodWeekly),
		string(constant.TimePeriodDaily),
	}
	AllowedSentimentCategories = []string{"overall", "sla", "customer_experience", "efficiency", "churn_risk"}
	validSortFields            = []string{"incident_id", "start_date", "last_response", "predicted_escalation_date", "escalation_date", "escalation_request_date"}
	validSortOrders            = []string{"asc", "desc", "ASC", "DESC"}
)

// InitValidator used to initiate go playground validator
func InitValidator() {
	validate = validator.New()
	RegisterTagName()
	AllowedPriorities = config.GetConfig().AllowedPriorities
	trans, err = SetTransLatorForStructError(validate)
	if err != nil {
		log.Logger.Debug(log.TraceMsgErrorOccurredFrom(constant.ErrInitValidatorMethod), zap.Error(err))
	}

	RegisterCustomValidation(validate)
	RegisterCustomTranslation(validate, trans)

	generalErrorCode = BuildGeneralErrorCode()
}

// BuildValidationErrorResponse used to build go playground validator error responses
func BuildValidationErrorResponse(requestID string, validationError error) *custom.ErrorResult {
	commonLogFields := []zap.Field{zap.String(constant.TraceMsgReqID, requestID)}
	log.Logger.Debug(log.TraceMsgFuncStart(BuildValidationErrorResponseMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(BuildValidationErrorResponseMethod), commonLogFields...)

	if validationError != nil {
		errorList := []custom.ErrorInfo{}

		for _, validationErrorsTranslation := range validationError.(validator.ValidationErrors) {
			errorList = append(errorList, custom.BuildErrorInfo(generalErrorCode[validationErrorsTranslation.Tag()], validationErrorsTranslation.Translate(trans), ""))
		}

		err := custom.BuildBadReqErrResultWithList(errorList...)
		return &err
	}

	return nil
}

// SetTransLatorForStructError used to set the translator for the struct error
func SetTransLatorForStructError(validate *validator.Validate) (ut.Translator, error) {
	uni := ut.New(en_US.New())
	translator, _ := uni.GetTranslator("en_US")
	validationErr := translations.RegisterDefaultTranslations(validate, translator)

	return translator, validationErr
}

// BuildGeneralErrorCode used to validate general error code
func BuildGeneralErrorCode() map[string]string {
	return map[string]string{
		required:                    constant.MissingRequiredFieldErrorCode,
		requiredWithout:             constant.MissingRequireWithoutFieldCode,
		requiredWith:                constant.MissingRequireWithFieldCode,
		minValue:                    constant.MinLengthFieldCode,
		maxValue:                    constant.MaxLengthFieldCode,
		alpha:                       constant.PatternErrorCode,
		alphaNumeric:                constant.PatternErrorCode,
		timestamp:                   constant.PatternErrorCode,
		intWithPlus:                 constant.PatternErrorCode,
		alphaNumericWithHyphen:      constant.PatternErrorCode,
		alphaNumericWithHyphenSpace: constant.PatternErrorCode,
		date:                        constant.PatternErrorCode,
		oneof:                       constant.PatternErrorCode,
		oneOfPriority:               constant.PatternErrorCode,
		year:                        constant.PatternErrorCode,
		oneOfBool:                   constant.PatternErrorCode,
		oneOfStatus:                 constant.PatternErrorCode,
		oneOfPeriod:                 constant.PatternErrorCode,
		omitEmpty:                   constant.PatternErrorCode,
		oneOfSortFields:             constant.PatternErrorCode,
		oneOfSortOrders:             constant.PatternErrorCode,
	}
}
