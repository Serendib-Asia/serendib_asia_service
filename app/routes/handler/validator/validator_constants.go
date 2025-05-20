package validator

// Regexs
const (
	alphaNumericRegex                        = "[^A-Za-z0-9]"
	alphaNumericWithHyphenRegex              = "[^A-Za-z0-9-]"
	alphaNumericWithHyphenSpaceRegex         = "[^A-Za-z0-9- ]"
	alphaNumericWithHyphenDotAndAddressRegex = "[^A-Za-z0-9-.@]"
	alphaRegex                               = "[^A-Za-z]"
	positiveIntegerWithPlusRegex             = `[^+0-9]`
	yearRegex                                = `^(19\d\d|20\d\d)$`
	timestampRegex                           = `^(\d{4})-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])T([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])Z`
	dateRegex                                = `^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])$`
)

// Validator keys
const (
	alphaNumeric                = "alphaNumeric"
	alphaNumericWithHyphen      = "alphaNumericWithHyphen"
	alphaNumericWithHyphenSpace = "alphaNumericWithHyphenSpace"
	alpha                       = "alpha"
	required                    = "required"
	requiredWithout             = "required_without"
	requiredWith                = "required_with"
	minValue                    = "min"
	maxValue                    = "max"
	timestamp                   = "timestamp"
	intWithPlus                 = "int_with_plus"
	date                        = "date"
	oneof                       = "oneof"
	oneOfPriority               = "oneOfPriority"
	week                        = "week"
	year                        = "year"
	oneOfBool                   = "oneOfBool"
	oneOfStatus                 = "oneOfStatus"
	oneOfPeriod                 = "oneOfPeriod"
	omitEmpty                   = "omitEmpty"
	oneOfSortFields             = "oneOfSortFields"
	oneOfSortOrders             = "oneOfSortOrders"
)

// Methods
const (
	CustomValidators                           = "CustomValidators"
	GenericBaseValidatorMethod                 = "GenericBaseValidator"
	GenericValidatorMethod                     = "GenericValidator"
	BuildValidationErrorResponseMethod         = "BuildValidationErrorResponse"
	ValidateCommonRequestMethod                = "ValidateCommonRequest"
	validateStartAndEndDateMethod              = "validateStartAndEndDate"
	ValidateDataRangeMethod                    = "ValidateDataRange"
	ValidateWeekRangeMethod                    = "ValidateWeekRange"
	ValidateYearRangeMethod                    = "ValidateYearRange"
	ValidateAgentInfoRequestMethod             = "ValidateAgentInfoRequest"
	ValidateAgentRequestMethod                 = "ValidateAgentRequest"
	ValidateDateRangeRequestMethod             = "ValidateDateRangeRequest"
	ValidatePaginatedDateRangeRequestMethod    = "ValidatePaginatedDateRangeRequest"
	ValidateDateTrendRequestMethod             = "ValidateDateTrendRequest"
	StructCasterMethod                         = "StructCaster"
	ValidatePaginatedCommonFilterRequestMethod = "ValidatePaginatedCommonFilterRequest"
	ValidateSentimentCategoriesMethod          = "ValidateSentimentCategories"
	ValidateBacklogCategoriesMethod            = "ValidateBacklogCategories"
	ValidateClientCategoriesSelectionMethod    = "ValidateClientCategoriesSelection"
	ValidateUserCategoriesSelectionMethod      = "ValidateUserCategoriesSelection"
	ValidateBacklogCategoriesSelectionMethod   = "ValidateBacklogCategoriesSelection"
	WeekComparison                             = "WeekComparison"
	customValidators                           = "customValidators"
)

// Dto
const (
	InvalidRequestType = "Invalid request type: %v"
)

// Specific constants for validation
const (
	Plus           = "+"
	JSON           = "json"
	Underscore     = "_"
	EmptyString    = ""
	DatePattern    = "2006-01-02"
	StartDate      = "start_date"
	EndDate        = "end_date"
	StartWeekField = "start_week"
	EndWeekField   = "end_week"
	StartYearField = "start_year"
	EndYearField   = "end_year"
	maximumWeek    = 53
)

// error message
const (
	invalidWeekRangeMessage  = "The start week (%d) cannot be later than the end week (%d)"
	sameYearInvalidWeekRange = "start week cannot be later than the end week"
)
