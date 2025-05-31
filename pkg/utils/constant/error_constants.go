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
	// General error codes
	ErrCodeInvalidRequest            = "INVALID_REQUEST"
	ErrCodeUnauthorized              = "UNAUTHORIZED"
	ErrCodeForbidden                 = "FORBIDDEN"
	ErrCodeNotFound                  = "NOT_FOUND"
	ErrCodeInternalServerError       = "INTERNAL_SERVER_ERROR"
	ErrCodeBadGateway                = "BAD_GATEWAY"
	ErrCodeServiceUnavailable        = "SERVICE_UNAVAILABLE"
	ErrCodeGatewayTimeout            = "GATEWAY_TIMEOUT"
	ErrCodeValidationError           = "VALIDATION_ERROR"
	ErrCodeDatabaseError             = "DATABASE_ERROR"
	ErrCodeExternalServiceError      = "EXTERNAL_SERVICE_ERROR"
	ErrCodeAuthenticationError       = "AUTHENTICATION_ERROR"
	ErrCodeAuthorizationError        = "AUTHORIZATION_ERROR"
	ErrCodeResourceNotFound          = "RESOURCE_NOT_FOUND"
	ErrCodeDuplicateResource         = "DUPLICATE_RESOURCE"
	ErrCodeInvalidInput              = "INVALID_INPUT"
	ErrCodeTimeout                   = "TIMEOUT"
	ErrCodeRateLimitExceeded         = "RATE_LIMIT_EXCEEDED"
	ErrCodeInvalidToken              = "INVALID_TOKEN"
	ErrCodeTokenExpired              = "TOKEN_EXPIRED"
	ErrCodeInvalidCredentials        = "INVALID_CREDENTIALS"
	ErrCodeAccountLocked             = "ACCOUNT_LOCKED"
	ErrCodeAccountDisabled           = "ACCOUNT_DISABLED"
	ErrCodeAccountExpired            = "ACCOUNT_EXPIRED"
	ErrCodeAccountInactive           = "ACCOUNT_INACTIVE"
	ErrCodeAccountPending            = "ACCOUNT_PENDING"
	ErrCodeAccountSuspended          = "ACCOUNT_SUSPENDED"
	ErrCodeAccountTerminated         = "ACCOUNT_TERMINATED"
	ErrCodeAccountDeleted            = "ACCOUNT_DELETED"
	ErrCodeAccountNotFound           = "ACCOUNT_NOT_FOUND"
	ErrCodeAccountAlreadyExists      = "ACCOUNT_ALREADY_EXISTS"
	ErrCodeAccountInvalid            = "ACCOUNT_INVALID"
	ErrCodeAccountInvalidState       = "ACCOUNT_INVALID_STATE"
	ErrCodeAccountInvalidType        = "ACCOUNT_INVALID_TYPE"
	ErrCodeAccountInvalidRole        = "ACCOUNT_INVALID_ROLE"
	ErrCodeAccountInvalidPermission  = "ACCOUNT_INVALID_PERMISSION"
	ErrCodeAccountInvalidAccess      = "ACCOUNT_INVALID_ACCESS"
	ErrCodeAccountInvalidOperation   = "ACCOUNT_INVALID_OPERATION"
	ErrCodeAccountInvalidRequest     = "ACCOUNT_INVALID_REQUEST"
	ErrCodeAccountInvalidResponse    = "ACCOUNT_INVALID_RESPONSE"
	ErrCodeAccountInvalidData        = "ACCOUNT_INVALID_DATA"
	ErrCodeAccountInvalidFormat      = "ACCOUNT_INVALID_FORMAT"
	ErrCodeAccountInvalidValue       = "ACCOUNT_INVALID_VALUE"
	ErrCodeAccountInvalidParameter   = "ACCOUNT_INVALID_PARAMETER"
	ErrCodeAccountInvalidArgument    = "ACCOUNT_INVALID_ARGUMENT"
	ErrCodeAccountInvalidOption      = "ACCOUNT_INVALID_OPTION"
	ErrCodeAccountInvalidConfig      = "ACCOUNT_INVALID_CONFIG"
	ErrCodeAccountInvalidSetting     = "ACCOUNT_INVALID_SETTING"
	ErrCodeAccountInvalidProperty    = "ACCOUNT_INVALID_PROPERTY"
	ErrCodeAccountInvalidAttribute   = "ACCOUNT_INVALID_ATTRIBUTE"
	ErrCodeAccountInvalidField       = "ACCOUNT_INVALID_FIELD"
	ErrCodeAccountInvalidColumn      = "ACCOUNT_INVALID_COLUMN"
	ErrCodeAccountInvalidTable       = "ACCOUNT_INVALID_TABLE"
	ErrCodeAccountInvalidDatabase    = "ACCOUNT_INVALID_DATABASE"
	ErrCodeAccountInvalidSchema      = "ACCOUNT_INVALID_SCHEMA"
	ErrCodeAccountInvalidQuery       = "ACCOUNT_INVALID_QUERY"
	ErrCodeAccountInvalidStatement   = "ACCOUNT_INVALID_STATEMENT"
	ErrCodeAccountInvalidTransaction = "ACCOUNT_INVALID_TRANSACTION"
	ErrCodeAccountInvalidConnection  = "ACCOUNT_INVALID_CONNECTION"
	ErrCodeAccountInvalidPool        = "ACCOUNT_INVALID_POOL"
	ErrCodeAccountInvalidDriver      = "ACCOUNT_INVALID_DRIVER"
	ErrCodeAccountInvalidProtocol    = "ACCOUNT_INVALID_PROTOCOL"
	ErrCodeAccountInvalidVersion     = "ACCOUNT_INVALID_VERSION"
	ErrCodeAccountInvalidEncoding    = "ACCOUNT_INVALID_ENCODING"
	ErrCodeAccountInvalidCompression = "ACCOUNT_INVALID_COMPRESSION"
	ErrCodeAccountInvalidEncryption  = "ACCOUNT_INVALID_ENCRYPTION"
	ErrCodeAccountInvalidHash        = "ACCOUNT_INVALID_HASH"
	ErrCodeAccountInvalidSignature   = "ACCOUNT_INVALID_SIGNATURE"
	ErrCodeAccountInvalidCertificate = "ACCOUNT_INVALID_CERTIFICATE"
	ErrCodeAccountInvalidKey         = "ACCOUNT_INVALID_KEY"
	ErrCodeAccountInvalidToken       = "ACCOUNT_INVALID_TOKEN"
	ErrCodeAccountInvalidSession     = "ACCOUNT_INVALID_SESSION"
	ErrCodeAccountInvalidCookie      = "ACCOUNT_INVALID_COOKIE"
	ErrCodeAccountInvalidHeader      = "ACCOUNT_INVALID_HEADER"

	// Legacy error codes (for backward compatibility)
	ErrAccessTokenCode  = "AUTH_001"
	BindingErrorCode    = "REQ_001"
	UnexpectedErrorCode = "SYS_001"

	// Add missing error codes for utils and handler usage
	ErrHTTPRequestCode       = "HTTP_REQUEST_ERROR"
	ErrStringToUintParseCode = "STRING_TO_UINT_PARSE_ERROR"
	ErrDateParseCode         = "DATE_PARSE_ERROR"
	ErrDateCompareCode       = "DATE_COMPARE_ERROR"
	ErrDataUnmarshalCode     = "DATA_UNMARSHAL_ERROR"
	ErrDataMarshalCode       = "DATA_MARSHAL_ERROR"
	ErrWeekValidateCode      = "WEEK_VALIDATE_ERROR"
	ErrIntToUintParseCode    = "INT_TO_UINT_PARSE_ERROR"
	ErrIntToStringParseCode  = "INT_TO_STRING_PARSE_ERROR"

	// Validation error codes
	MissingRequiredFieldErrorCode  = "MISSING_REQUIRED_FIELD"
	MissingRequireWithoutFieldCode = "MISSING_REQUIRED_WITHOUT_FIELD"
	MissingRequireWithFieldCode    = "MISSING_REQUIRED_WITH_FIELD"
	MinLengthFieldCode             = "MIN_LENGTH_FIELD"
	MaxLengthFieldCode             = "MAX_LENGTH_FIELD"
	PatternErrorCode               = "PATTERN_ERROR"

	// Service and repository error codes
	ErrRecordNotFoundCode    = "RECORD_NOT_FOUND"
	ErrDatabaseCode          = "DATABASE_ERROR"
	ErrEmptyAuthHeaderCode   = "EMPTY_AUTH_HEADER"
	ErrTokenSplitToArrayCode = "TOKEN_SPLIT_TO_ARRAY_ERROR"

	// Additional validation error codes
	ErrYearValidationCode  = "YEAR_VALIDATION_ERROR"
	ErrEmptyCategoriesCode = "EMPTY_CATEGORIES_ERROR"
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
	AppendToArrayIfNotEmpty = "AppendToArrayIfNotEmpty"
	SplitAndDecodeMethod    = "SplitAndDecode"
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
