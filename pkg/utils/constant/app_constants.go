package constant

// Security service tenant id for ICX application
const (
	SecSrvICXTenantName = "icx-dashboard"
)

// DatadogTracingSink is the sink name for Datadog tracing.
const DatadogTracingSink = "datadog"

// Special characters
const (
	Colon                      = ":"
	Basepath                   = "./"
	Empty                      = ""
	Hyphen                     = "-"
	Dot                        = "."
	Slash                      = "/"
	Space                      = " "
	OpenCurlyBrace             = "{"
	CloseCurlyBrace            = "}"
	Comma                      = ","
	EqualsQuestionMark         = " = ?"
	NewlineBytes               = '\n'
	Zero                       = 0
	True                       = true
	False                      = false
	TRUE                       = "TRUE"
	FALSE                      = "FALSE"
	TwoPlaceholdersWithHyphen  = "%v-%v"
	ThreePlaceholderWithHyphen = "%v-%v-%v"
	DatePattern                = "2006-01-02"
	BearerSpace                = "Bearer "
)

// IntOne, IntTwo, and IntThree are integer constants.
const (
	IntOne   int = 1
	IntTwo   int = 2
	IntThree int = 3
	IntFour  int = 4
)

// File constants
const (
	DocumentHTML  = "document.html"
	StaticHTML    = "static.html"
	Doc           = "doc"
	Static        = "static"
	HTML          = "html"
	DotJSON       = ".json"
	DotGoldenJSON = ".golden.json"
)

// Token constants
const (
	Access  = "access"
	Refresh = "refresh"
	Code    = "code"
)

// Incident Type constants
const (
	IncidentTypeLikelyToEscalate    = "escalation_likely"
	IncidentTypeEscalationRequested = "escalation_requested"
	IncidentTypeEscalated           = "escalation_done"
)

// Logger Messages
const (
	TraceMsgBeforeFetching     = "Before Fetching %v"
	TraceMsgAfterFetching      = "After Fetching %v"
	TraceMsgBeforeInsertion    = "Before Creating %v"
	TraceMsgAfterInsertion     = "After Creating %v"
	TraceMsgBeforeUpdate       = "Before Update %v"
	TraceMsgAfterUpdate        = "After Update %v"
	TraceMsgFuncEnd            = "%v End here"
	TraceMsgFuncStart          = "%v Start here"
	TraceMsgRequestInitiated   = "%v request initiated"
	TraceMsgReqID              = "Request Id"
	TraceMsgReqBody            = "Request Body"
	BuildedQuery               = "Builded query"
	TraceMsgRequestHeader      = "Request Header"
	TranceMsgResponse          = "Response"
	TraceMsgBeforeInvoke       = "Before Call %v"
	TraceMsgAfterInvoke        = "After Call %v"
	TraceMsgBeforeRollback     = "Before rollback from %s"
	TraceMsgAfterRollback      = "After rollback from %s"
	TraceMsgBeforeCommit       = "Before commit from %s"
	TraceMsgAfterCommit        = "After commit from %s"
	TraceMsgAPIResponse        = "Build API Response"
	TraceMsgResponseDetails    = "Response Details"
	TraceMsgAPISuccess         = "Success Response"
	TraceMsgAPIErrorResponse   = "Error Response"
	TraceMsgErrorDetails       = "Error Details"
	MethodInput                = "Method Input"
	SQLQuery                   = "SQL Query: "
	Result                     = "Result"
	DebugNote                  = "Debug workflow"
	ErrorNote                  = "Error Note"
	Error                      = "Error"
	HTMLPassErr                = "HTML Template pass Error"
	Response                   = "Response"
	AuthResponse               = "Auth response"
	ConvertingStrToUint        = "Converting string to uint"
	ConvertingIntToUint        = "Converting int to uint"
	EmptyHeaderDetails         = "%s does not exist in the request header"
	InactiveChangePwdRecord    = "Inactive change password record"
	InvalidKeyForChangePwd     = "Invalid key for change password"
	InvalidTokenForChangePwd   = "Invalid token for change password"
	TimeExpiredForChangePwd    = "Time expired for change password"
	DataFetchedSuccess         = "Data fetched successfully from %s"
	MethodOutput               = "Method Output"
	MethodError                = "Method Error"
	TraceMsgRespBody           = "Response Body"
	TraceMsgBeforeParse        = "Trace before parse in %s"
	TraceMsgAfterParse         = "Trace after parse in %s"
	TraceRequestType           = "Request Type %T"
	Request                    = "Request"
	Successfully               = "Successfully %s"
	AlreadyExistWithFieldValue = "%s already exists with provided %s %s"
	DefaultClientDetectedAs    = "Default client detected as: %s"
)

// Utils func
const (
	Dial = "Dial"
)

// HTTPMethod represents the type for HTTP methods.
type HTTPMethod string

// Get is a constant for the HTTP GET method.
const Get HTTPMethod = "GET"

// Post is a constant for the HTTP POST method.
const Post HTTPMethod = "POST"

// Patch is a constant for the HTTP PATCH method.
const Patch HTTPMethod = "PATCH"

// Delete is a constant for the HTTP DELETE method.
const Delete HTTPMethod = "DELETE"

// Header fields
const (
	XForwardedForHeader    = "X-Forwarded-For"
	UserAgentHeader        = "User-Agent"
	PlatformHeader         = "Sec-Ch-Ua-Platform"
	BrowserHeader          = "Sec-Ch-Ua"
	Authorization          = "Authorization"
	UserID                 = "UserID"
	Domain                 = "Domain"
	InstanceID             = "InstanceId"
	TenantID               = "TenantID"
	XAccessTokenFromHeader = "X-Access-Token"
	UserKey                = "User"
)

// Method names
const (
	DateParserMethod                  = "DateParser"
	DateCompareMethod                 = "DateCompare"
	CallHTTPEndpointMethod            = "CallHTTPEndpoint"
	HTTPRequestMethod                 = "HTTPRequest"
	SetResponseToAgentMethod          = "setResponseToAgent"
	GetResponseHeadersMethod          = "getResponseHeaders"
	BuildAndParseRequestMethod        = "buildAndParseRequest"
	DataUnmarshalMethod               = "DataUnmarshal"
	ValidateResponseMethod            = "validateResponse"
	ExternalAPICallMethod             = "ExternalAPICall"
	GeneralServiceCallMethod          = "GeneralServiceCall"
	AiMlServiceAPICallMethod          = "AiMlServiceAPICall"
	AppendToArrayIfNotEmpty           = "AppendToArrayIfNotEmpty"
	ValidateAndGetWeekStartDateMethod = "ValidateWeekAndGetWeekStartDate"
	ParseBaseURLMethod                = "ParseBaseURL"
	GenerateTicketURLMethod           = "GenerateTicketURL"
	IsStrValueTrueMethod              = "IsStrValueTrue"
)

// Service flags
const (
	SDesk = "sdesk"
	ICX   = "icx"
)

// External Endpoints security service
const (
	SecSrvValidateTokenURI = "/icx/auth/v1/validate/token/icx-dashboard"
	SecSrvUpdateTenantURI  = "/icx/auth/v1/tenant/user"
)

// goconst
const (
	AuditEventsTable         = "`audit.Events`"
	RequestSubjectTable      = "`request.Subject`"
	AuditEventTypesTable     = "`audit.EventTypes`"
	InstanceRequestTypeTable = "`instance.Request.Type`"
	RequestBody              = "`request.Body`"
	RequestTable             = "`request`"
	User                     = "`user`"
	RequestSLA               = "`request.Sla`"
	RequestSubject           = "`request.Subject`"
	LowerTrueStr             = "true"
	UpperTrueStr             = "TRUE"
	UpperFalseStr            = "FALSE"
)

type TimePeriod string

const (
	TimePeriodYearly    TimePeriod = "yearly"
	TimePeriodQuarterly TimePeriod = "quarterly"
	TimePeriodMonthly   TimePeriod = "monthly"
	TimePeriodWeekly    TimePeriod = "weekly"
	TimePeriodDaily     TimePeriod = "daily"
)
