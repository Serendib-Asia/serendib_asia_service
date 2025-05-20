package constant

// Log error messages
const (
	PanicOccurred                       = "error or panic occurred with following stacktrace"
	StackTrace                          = "error stacktrace"
	ErrorRequestBody                    = "error Request"
	InvalidInputAndPassErr              = "error input provided is invalid & unable to parse"
	MissingRequiredField                = "missing Required Field"
	ErrorOccurredFromService            = "missing Occurred from  %s"
	ErrorOccurredFromMethod             = "error Occurred from  Method %s"
	BindingErrorMessage                 = "error Occurred when bind the request context"
	ErrorOccurredWhen                   = "error Occurred When %s"
	ErrorOccurredWhenSelecting          = "error occurred when selecting %s"
	ErrorOccurredWhenInserting          = "error occurred when inserting %s"
	ErrorOccurredWhenTemplateParse      = "error occourred when template parse"
	ErrorOccurredWhenDeleting           = "error occurred when deleting %s"
	ErrorOccurredWhenUpdating           = "error occurred when updating %s"
	ErrorOccurredWhenCounting           = "error occurred when counting %s"
	ErrorOccurredWhenHashing            = "error occurred when hashing"
	ErrorOccurredWhenHashCompare        = "error occurred when hash compare"
	ErrOccurredWhenSendingEmail         = "error occurred when sending email"
	ErrOccouredWhenParseDate            = "error occurred when parse %s"
	ErrorOccurredWhenCompareGreaterThan = "%s should be greater than %s"
	ErrOccurredWhenCompGreaterThanOrEql = "%s should be greater than or equal %s"
	ErrInvalidIncidentType              = "invalid incident type"
	ErrorOccurredWhenWeekValidation     = "%s should be a valid week number"
	ErrInitValidatorMethod              = "initiate validator"
	ErrOccurredFromServiceCall          = "error occurred from service call"
	ErrOccurredWhenValidateWeek         = "invalid week when get the week start date"
	ErrClientNotFound                   = "client with ID %d not found"
	ErrOccurredDuringServiceCall        = "error occurred during %s"
	ErrEmptyDateMsg                     = "Start date or end date is empty"
	ErrEmptyDateCode                    = "EMPTY_DATE"
)

// "No client found with ID: %d"

// Data state constants
const (
	Encoded = "Encoded"
	Decoded = "Decoded"
)

// Error codes
const (
	UnexpectedErrorCode            = "ICX0000"
	BindingErrorCode               = "ICX0001"
	MissingRequiredFieldErrorCode  = "ICX0002"
	MissingRequireWithoutFieldCode = "ICX0003"
	MissingRequireWithFieldCode    = "ICX0004"
	MinLengthFieldCode             = "ICX0005"
	MaxLengthFieldCode             = "ICX0006"
	GreaterValueFieldCode          = "ICX0007"
	PatternErrorCode               = "ICX0008"
	TestConnectionFilCode          = "ICX0009"
	// Auth error codes
	ErrEmptyAuthHeaderCode   = "ICX0010"
	ErrInvalidAuthHeaderCode = "ICX0011"
	// DB error codes
	ErrRecordNotFoundCode   = "ICX0012"
	ErrDatabaseCode         = "ICX0013"
	ErrBeginTransactionCode = "ICX0014"
	// Parse error codes
	ErrStringToUintParseCode = "ICX0015"
	ErrDateParseCode         = "ICX0016"
	ErrDateCompareCode       = "ICX0017"
	ErrRequestParseCode      = "ICX0018"
	ErrHTTPRequestCode       = "ICX0019"
	// bootstrap error
	ErrTokenSplitToArrayCode = "ICX0020"
	ErrDataUnmarshalCode     = "ICX0021"
	ErrDataMarshalCode       = "ICX0022"
	// access token error codes
	ErrAccessTokenCode = "ICX0023"
	ErrDataCastingCode = "ICX0024"
	// Escalation error codes
	ErrFieldsMustBeTrueOrFalseCode = "ICX0022"
	ErrFieldsCannotBeNegativeCode  = "ICX0023"
	// Agent error codes
	ErrWeekValidateCode   = "ICX0024"
	ErrYearValidationCode = "ICX0025"
	ErrTenantNotFoundCode = "ICX0026"
	// Sentiment score request error code
	ErrEmptyCategoriesCode = "ICX0027"
	// Service error codes
	ErrServiceCode = "ICX0028"
	// Parse error codes (cont.)
	ErrIntToUintParseCode   = "ICX0029"
	ErrIntToStringParseCode = "ICX0030"
	ErrStringToIntParseCode = "ICX0031"
)

// Error messages
const (
	InvalidRequestErrorMessage   = "Invalid Request validation Error Occurred"
	UnexpectedErrorMessage       = "Unexpected Error occurred at %s"
	UnexpectedWhenMarshalError   = "Unexpected Error occurred when Marshal the data"
	UnexpectedWhenUnmarshalError = "Unexpected Error occurred when Unmarshal the data"
	UnexpectedFileCreateError    = "Unexpected Error occurred when Create the %s file"
	TestConnectionFilMessage     = "Error Occurred when TestConnection"
	InvalidDataOrFile            = "Invalid Data or File "
	InvalidJSONTestData          = "Invalid Json test data"
	FileReadError                = "Unable to read file %v"
	HTMLTempPassError            = "HTML Template parse error"
	// auth error message
	ErrOccurredWhenGenAccessTokenMsg         = "error occurred when Generate Access Token"
	ErrOccurredWhenAuthHeaderSplitToArrayMsg = "error occurred When Authorization header split to array"
	ErrorOccurredWhenDecodeStringMsg         = "error occurred When Decode String"
	ErrorOccurredWhenPemBlockDecodeMsg       = "error occurred When decode PEM block containing private key"
	ErrorOccurredWhenParsingKeyMsg           = "error occurred when parsing private key"
	ErrorOccurredWhenDecryptStringMsg        = "error occurred When decrypt String"
	ErrorOccurredWhenHashingPasswordMsg      = "error occurred When hashing password"
	ErrorOccurredWhenPassingEmptyKeyMsg      = "error occurred when passing empty private key"
	ErrInvalidClientCredentialsMsg           = "invalid Client Credentials"
	ErrInvalidTokenSignatureMsg              = "invalid Token Signature"
	ErrOccurredWhenAccessingTokenMsg         = "error Occurred when accessing token"
	ErrEmptyAuthHeaderMsg                    = "authorization header is required"
	ErrInvalidAuthHeaderMsg                  = "invalid authorization header"
	ErrInParsingTokenMsg                     = "error Occurred when parsing token"
	ErrTokenDoesNotHaveRequiredRoleMsg       = "token does not have required role"
	ErrInvalidUserCredentialsMsg             = "Invalid User Credentials"
	ErrInvalidRoleMsg                        = "token should have (at least one of) the following role(s): %s, but it has following role(s): %s"
	ErrStringToUintParseMsg                  = "error occurred when converting string to uint"
	ErrIntToUintParseMsg                     = "error occurred when converting int to uint"
	ErrIntToStringParseMsg                   = "error occurred when converting int to string"
	ErrStringToIntParseMsg                   = "error occurred when converting string to int"
	ErrInvalidRefreshTokenMsg                = "invalid refresh token"
	ErrInvalidPasswordChangeMsg              = "invalid password change request"
	ErrPasswordChangeRequestExpiredMsg       = "password change request expired"
	ErrInvalidAuthCodeMsg                    = "invalid authorization code"
	ErrConvertStringToBoolMsg                = "error occurred when convert string:%s to bool"
	ErrConvertStringToUint                   = "error occurred when convert string:%s to uint"
	// password error message
	ErrPasswordChangeRequestDecodeMsg = "error Occurred when decode html template"
	// Email error messages
	ErrEmailUniqueConstraintViolationMsg = "Email already exists"
	// Timeonly error message
	TimeFormatErrorMsg     = "Invalid time format. Time should be in the format HH:MM:SS."
	InvalidTypeForTimeOnly = "invalid type %T for TimeOnly"
	WorkTimeErrorMsg       = "Work start time must be less than work end time"
	// fiber aqua messages
	CallingFromMethod          = "calling from"
	ErrRequestParseMsg         = "error occurred when parse the request"
	APIErrorResponse           = "API error response"
	ErrHTTPRequestMsg          = "error occurred when send HTTP request"
	ErrInvalidInputWhenParse   = "Invalid inputs when parse request"
	ErrInvalidInput            = "Invalid inputs for request"
	ErrUnableToReadHeader      = "Unable to read data from header"
	ErrInvalidRequestFormatMsg = "Invalid request format"
	// Escalation error messages
	ErrFieldsMustBeTrueOrFalseMsg = "Fields must be true or false"
	ErrFieldsCannotBeNegativeMsg  = "Fields cannot be negative"
	// Access token error message
	ErrAccessTokenMsg    = "Error found in the given access token"
	ErrTenantNotFoundMsg = "Tenant not found"
	// Sentiment score request error messages
	ErrEmptyCategoriesMsg = "At least one category must be selected"
	// User error messages
	DuplicateEmailErrorMessage = "Email already exists"
	InvalidCredentialsMessage  = "Invalid email or password"
	UserNotFoundMessage        = "User not found"

	// User error codes
	DuplicateEmailErrorCode = "EMAIL_EXISTS"
	InvalidCredentialsCode  = "INVALID_CREDENTIALS"
	UserNotFoundCode        = "USER_NOT_FOUND"
)

// "Client validation failed"

// Util methods
const (
	SplitAndDecodeMethod = "SplitAndDecode"
)

// DB errors for database operations
const (
	DBConnectionOpenError          = "Error occurred when opening the database connection"
	DBConnectionCloseError         = "Error occurred when closing the database connection"
	DBInitFailError                = "Failed to initialize database"
	DBConnectionIsNotEstablished   = "database connection is not established"
	DBErrorOccurredWhenAutoMigrate = "error occurred when auto migrate"
)

// DB errors returned by database operations
const (
	ErrRecordNotFoundMsg          = "Record not found"
	ErrDatabaseMsg                = "Database error"
	ErrOccurredWhileRetrieving    = "An error occurred while retrieving data from the database"
	ErrorOccurredWhenRetrieveData = "error occurred when retrieving data from %s: %s"
	ErrNotFound                   = "%s not found for %v"
	ErrBeginTransactionMsg        = "Error occurred when begin transaction"
)

// Server errors for the ICX dashboard
const (
	ServerInitFailError = "Failed to initialize ICX dashboard server"
)

// JWT errors
const (
	ErrInvalidSigningMethod                = "invalid signing method"
	ErrOccurredWhenGettingJwtSigningMethod = "error occurred when getting jwt signing method"
	ErrOccurredWhenSigningJWTToken         = "error occurred when signing jwt token"
)
