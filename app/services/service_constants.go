package services

const (
	// Common service methods
	StructCastarMethod = "StructCastar"

	// Health service methods
	ReadyzServiceMethod = "ReadyzService"

	// Agent service methods
	GetAgentsMethod                 = "GetAgents"
	GetCompletedCaseCountMethod     = "GetCompletedCaseCount"
	GetAgentSkillScoresMethod       = "GetAgentSkillScores"
	GetAgentStatusMethod            = "GetAgentStatus"
	buildAgentStatusResponseMethod  = "buildAgentStatusResponse"
	buildOutOfDatesMapMethod        = "buildOutOfDatesMap"
	getAuthUserDetailsMethod        = "getAuthUserDetails"
	buildUsersMapMethod             = "buildUsersMap"
	buildAgentIdentityListMethod    = "buildAgentIdentityList"
	updateNewAgentMethod            = "updateNewAgent"
	ConvertWeekToStartDateMethod    = "ConvertWeekToStartDate"
	GetAgentInfoFromIDMethod        = "GetAgentInfoFromID"
	ValidateAgentsMethod            = "ValidateAgents"
	FetchAgentExistenceMethod       = "FetchAgentExistence"
	updateTenantMethod              = "updateTenant"
	validateAndUpdateICXAgentMethod = "validateAndUpdateICXAgent"
	updateICXAgentMethod            = "updateICXAgent"
	ValidateExistingAgentMethod     = "ValidateExistingAgent"
	ValidateAndUpdateTenantMethod   = "ValidateAndUpdateTenant"
	GetExistingICXAgentMethod       = "GetExistingICXAgent"

	// Client service methods
	GetOverallScoreMethod          = "GetOverallScore"
	FindSentimentsByClientIDMethod = "FindSentimentsByClientID"
	GetUsersWithClientMethod       = "GetUsersWithClient"
	ValidateClientMethod           = "ValidateClient"
	getClientMethod                = "getClient"
	GetClientListMethod            = "GetClientList"

	// Incident service methods
	GetTopIncidentsByPriorityMethod      = "GetTopIncidentsByPriority"
	GetRestoreIncidentsMethod            = "GetRestoreIncidents"
	GetTopIncidentsByServiceMethod       = "GetTopIncidentsByService"
	GetTopIncidentsByCostMethod          = "GetTopIncidentsByCost"
	getManualIncidentsDataMethod         = "getManualIncidentsData"
	getRestorationMethodData             = "getRestorationMethodData"
	getIncidentRestorationDataMethod     = "getIncidentRestorationData"
	GetIncidentsTrendMethod              = "GetIncidentsTrend"
	determineFrequencyMethod             = "determineFrequency"
	getHandlerMethod                     = "getHandler"
	getFilteredIncidentsByPriorityMethod = "getFilteredIncidentsByPriority"

	// Service bootstrap methods
	BeginNewTransactionMethod    = "BeginNewTransaction"
	HandleTransactionMethod      = "HandleTransaction"
	checkRepoErrorMethod         = "checkRepoError"
	ExtractAccessTokenMethod     = "ExtractAccessToken"
	DecodeJWTTokenMethod         = "DecodeJWTToken" // #nosec G101
	ExtractUserIDFromTokenMethod = "ExtractUserIDFromToken"
	TransactionDoesNotExist      = "Transaction does not exist"

	// dashboard service methods
	GetMasterDashboardDataMethod = "GetMasterDashboardData"
	getPrioritiesMethod          = "getPriorities"
	getClientsMethod             = "getClients"
	getServicesMethod            = "getServices"
	getCategoriesMethod          = "getCategories"
	getInstanceMethod            = "getInstance"

	// Alerts service methods
	GetAlertsMethod = "GetAlerts"

	// Metrics service methods
	GetSLAMetricsMethod         = "GetSLAMetrics"
	getAllSLAMetricsMethod      = "getAllSLAMetrics"
	getFilteredSLAMetricsMethod = "getFilteredSLAMetrics"

	// Tickets service methods
	GetAssignedTicketsMethod = "GetAssignedTickets"
	mapTicketsMethod         = "mapTickets"
	mapCasesSummaryMethod    = "mapCasesSummary"
	mapIncidentSummaryMethod = "mapIncidentSummary"

	// Assets service methods
	GetCurrentlyDownMethod = "GetCurrentlyDown"

	// Escalation service methods
	GetEscalationsMethod              = "GetEscalations"
	mapEscalationTicketsWithURLMethod = "mapEscalationTicketsWithURL"
	isLikelyToEscalateRequestMethod   = "isLikelyToEscalateRequest"
	GetEscalationSummaryMethod        = "GetEscalationSummary"

	// Runbook service methods
	GetRunbooksMethod            = "GetRunbooks"
	CreateRunbookMethod          = "CreateRunbook"
	SetActiveRunbookMethod       = "SetActiveRunbook"
	GetAgentIDfromTenantMethod   = "GetAgentIDfromTenant"
	usrByTenantsMethod           = "usrByTenants"
	getUserDetailsByUserIDMethod = "getUserDetailsByUserID"
	createNewRunbookMethod       = "createNewRunbook"
	findRunbookByIDMethod        = "findRunbookByID"
	RunbookCodeGenServiceMethod  = "RunbookCodeGenService"

	// Backlog service methods
	GetBacklogDataMethod                 = "GetBacklogData"
	GetIncidentSummaryMethod             = "GetIncidentSummary"
	GetInactiveIncidentsMethod           = "GetInactiveIncidents"
	getTotalIncidentsCountMethod         = "getTotalIncidentsCount"
	getTotalInactiveIncidentsCountMethod = "getTotalInactiveIncidentsCount"
	mapBacklogTicketsMethod              = "mapBacklogTickets"

	// Sentiment service methods
	GetSentimentMetricsMethod     = "GetSentimentMetrics"
	GetClientSentimentScoreMethod = "GetClientSentimentScore"
	GetUserSentimentScoreMethod   = "GetUserSentimentScore"

	// Customer insights service methods
	GetTicketSummaryMethod           = "GetTicketSummary"
	GetCustomerScoreIndicatorMethod  = "GetCustomerScoreIndicator"
	GetAccountReporterInsightsMethod = "GetAccountReporterInsights"
	GetResponseTimeMetricsMethod     = "GetResponseTimeMetrics"
	GetCasesSummaryMethod            = "GetCaseSummary"
	GetSentimentHistoryMethod        = "GetSentimentHistory"
	CreateNoteMethod                 = "CreateNotes"
	processNewNoteMethod             = "processNewNote"
	GetNotesMethod                   = "GetNotes"

	// Create tenant service methods
	CreateTenantMethod        = "CreateTenant"
	processNewTenantMethod    = "processNewTenant"
	ErrInvalidEmailDomainCode = "INVALID_EMAIL_DOMAIN"

	// Security service Methods
	SecSrvUpdateTenantMethod = "SecSrvUpdateTenant"
	//nolint:gosec // This is not sensitive information
	SecSrvTokenValidationMethod = "SecSrvTokenValidation"
	updateSecSrvTenantMethod    = "updateSecSrvTenant"

	// Instance service methods
	GetInstanceDetailsMethod = "GetInstanceDetails"
)

// Tables & column
const (
	ClientSentiments  = "client_sentiments"
	UserSentiments    = "user_sentiments"
	outOfOfficeColumn = "out_of_office"
)

// Constants
const (
	Customer           = "Customer"
	AllAgents          = "all agents"
	CompletedIncidents = "completed incidents"
	AgentSkills        = "agent skills"
	AgentStatus        = "agent status"
	Incidents          = "incidents"
	Clients            = "clients"
	codeGenerationType = "code-generation"
	Agent              = "agent"
	Instances          = "Instances"
)

// Backlog categories
const (
	MostViewed          = "MostViewed"
	InactiveCases       = "InactiveCases"
	LastOutbound        = "LastOutbound"
	UniqueResponders    = "UniqueResponders"
	PriorityByAttnScore = "PriorityByAttnScore"
	MostConversations   = "MostConversations"
	SentimentScore      = "SentimentScore"
	AgingCases          = "AgingCases"
)

// Sentiment categories
const (
	Overall            = "Overall"
	SLA                = "SLA"
	CustomerExperience = "CustomerExperience"
	Efficiency         = "Efficiency"
	ChurnRisk          = "ChurnRisk"
)

// Account reporter insights response categories
const (
	Accounts        = "Accounts"
	Reporters       = "Reporters"
	IncidentDetails = "IncidentDetails"
)

// Incidents trend constants
const (
	dateFormat       = "2006-01-02"
	daily            = "daily"
	weekly           = "weekly"
	monthly          = "monthly"
	weeklyThreshold  = 7
	monthlyThreshold = 60
)
