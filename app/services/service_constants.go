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
	GetEscalationsMethod       = "GetEscalations"
	GetEscalationSummaryMethod = "GetEscalationSummary"
	// Escalation service methods -private
	mapEscalationsMethod = "mapEscalations"

	// Backlog service methods
	GetBacklogIncidentsMethod = "GetBacklogIncidents"
	// Backlog service methods -private
	mapBacklogIncidentsMethod = "mapBacklogIncidents"

	// Customer insights service methods
	GetCustomerScoreIndicatorMethod         = "GetCustomerScoreIndicator"
	GetInternalCustomerScoreIndicatorMethod = "GetInternalCustomerScoreIndicator"
	GetNumberOfUserAccountsForClientMethod  = "GetNumberOfUserAccountsForClient"
	GetNumberOfReportersForClientMethod     = "GetNumberOfReportersForClient"
	GetUserIncidentDetailsForClientMethod   = "GetUserIncidentDetailsForClient"
	GetResponseTimeMetricsMethod            = "GetResponseTimeMetrics"
	GetCasesSummaryMethod                   = "GetCasesSummary"
	GetSentimentHistoryMethod               = "GetSentimentHistory"
	CreateNoteMethod                        = "CreateNote"
	GetNotesMethod                          = "GetNotes"
	GetTotalIncidentsCountMethod            = "GetTotalIncidentsCount"
	GetTotalInactiveIncidentsCountMethod    = "GetTotalInactiveIncidentsCount"
	GetNotesCountMethod                     = "GetNotesCount"
	GetCasesCountMethod                     = "GetCasesCount"
	GetIncidentsCountMethod                 = "GetIncidentsCount"
	// private
	getPeriodFieldMethod = "getPeriodField"

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

	// General service methods
	GetUserDetailsMethod                                                  = "GetUserDetails"
	ValidateTokenMethod                                                   = "ValidateToken"
	GetAllUsersMethod                                                     = "GetAllUsers"
	GetUserByIdMethod                                                     = "GetUserById"
	GetUserByEmailMethod                                                  = "GetUserByEmail"
	GetUserByUsernameMethod                                               = "GetUserByUsername"
	GetUserByPhoneMethod                                                  = "GetUserByPhone"
	GetUserByRoleMethod                                                   = "GetUserByRole"
	GetUserByStatusMethod                                                 = "GetUserByStatus"
	GetUserByCreatedDateMethod                                            = "GetUserByCreatedDate"
	GetUserByUpdatedDateMethod                                            = "GetUserByUpdatedDate"
	GetUserByLastLoginDateMethod                                          = "GetUserByLastLoginDate"
	GetUserByLastLoginIpMethod                                            = "GetUserByLastLoginIp"
	GetUserByLastLoginUserAgentMethod                                     = "GetUserByLastLoginUserAgent"
	GetUserByLastLoginLocationMethod                                      = "GetUserByLastLoginLocation"
	GetUserByLastLoginDeviceMethod                                        = "GetUserByLastLoginDevice"
	GetUserByLastLoginBrowserMethod                                       = "GetUserByLastLoginBrowser"
	GetUserByLastLoginOsMethod                                            = "GetUserByLastLoginOs"
	GetUserByLastLoginPlatformMethod                                      = "GetUserByLastLoginPlatform"
	GetUserByLastLoginCountryMethod                                       = "GetUserByLastLoginCountry"
	GetUserByLastLoginCityMethod                                          = "GetUserByLastLoginCity"
	GetUserByLastLoginRegionMethod                                        = "GetUserByLastLoginRegion"
	GetUserByLastLoginTimezoneMethod                                      = "GetUserByLastLoginTimezone"
	GetUserByLastLoginLanguageMethod                                      = "GetUserByLastLoginLanguage"
	GetUserByLastLoginReferrerMethod                                      = "GetUserByLastLoginReferrer"
	GetUserByLastLoginReferrerDomainMethod                                = "GetUserByLastLoginReferrerDomain"
	GetUserByLastLoginReferrerPathMethod                                  = "GetUserByLastLoginReferrerPath"
	GetUserByLastLoginReferrerQueryMethod                                 = "GetUserByLastLoginReferrerQuery"
	GetUserByLastLoginReferrerFragmentMethod                              = "GetUserByLastLoginReferrerFragment"
	GetUserByLastLoginReferrerProtocolMethod                              = "GetUserByLastLoginReferrerProtocol"
	GetUserByLastLoginReferrerHostMethod                                  = "GetUserByLastLoginReferrerHost"
	GetUserByLastLoginReferrerPortMethod                                  = "GetUserByLastLoginReferrerPort"
	GetUserByLastLoginReferrerUsernameMethod                              = "GetUserByLastLoginReferrerUsername"
	GetUserByLastLoginReferrerPasswordMethod                              = "GetUserByLastLoginReferrerPassword"
	GetUserByLastLoginReferrerHashMethod                                  = "GetUserByLastLoginReferrerHash"
	GetUserByLastLoginReferrerSearchMethod                                = "GetUserByLastLoginReferrerSearch"
	GetUserByLastLoginReferrerOriginMethod                                = "GetUserByLastLoginReferrerOrigin"
	GetUserByLastLoginReferrerHrefMethod                                  = "GetUserByLastLoginReferrerHref"
	GetUserByLastLoginReferrerProtocolRelativeMethod                      = "GetUserByLastLoginReferrerProtocolRelative"
	GetUserByLastLoginReferrerHostRelativeMethod                          = "GetUserByLastLoginReferrerHostRelative"
	GetUserByLastLoginReferrerPathRelativeMethod                          = "GetUserByLastLoginReferrerPathRelative"
	GetUserByLastLoginReferrerQueryRelativeMethod                         = "GetUserByLastLoginReferrerQueryRelative"
	GetUserByLastLoginReferrerFragmentRelativeMethod                      = "GetUserByLastLoginReferrerFragmentRelative"
	GetUserByLastLoginReferrerProtocolRelativeHostMethod                  = "GetUserByLastLoginReferrerProtocolRelativeHost"
	GetUserByLastLoginReferrerProtocolRelativePathMethod                  = "GetUserByLastLoginReferrerProtocolRelativePath"
	GetUserByLastLoginReferrerProtocolRelativeQueryMethod                 = "GetUserByLastLoginReferrerProtocolRelativeQuery"
	GetUserByLastLoginReferrerProtocolRelativeFragmentMethod              = "GetUserByLastLoginReferrerProtocolRelativeFragment"
	GetUserByLastLoginReferrerProtocolRelativeHostPathMethod              = "GetUserByLastLoginReferrerProtocolRelativeHostPath"
	GetUserByLastLoginReferrerProtocolRelativeHostQueryMethod             = "GetUserByLastLoginReferrerProtocolRelativeHostQuery"
	GetUserByLastLoginReferrerProtocolRelativeHostFragmentMethod          = "GetUserByLastLoginReferrerProtocolRelativeHostFragment"
	GetUserByLastLoginReferrerProtocolRelativePathQueryMethod             = "GetUserByLastLoginReferrerProtocolRelativePathQuery"
	GetUserByLastLoginReferrerProtocolRelativePathFragmentMethod          = "GetUserByLastLoginReferrerProtocolRelativePathFragment"
	GetUserByLastLoginReferrerProtocolRelativeQueryFragmentMethod         = "GetUserByLastLoginReferrerProtocolRelativeQueryFragment"
	GetUserByLastLoginReferrerProtocolRelativeHostPathQueryMethod         = "GetUserByLastLoginReferrerProtocolRelativeHostPathQuery"
	GetUserByLastLoginReferrerProtocolRelativeHostPathFragmentMethod      = "GetUserByLastLoginReferrerProtocolRelativeHostPathFragment"
	GetUserByLastLoginReferrerProtocolRelativeHostQueryFragmentMethod     = "GetUserByLastLoginReferrerProtocolRelativeHostQueryFragment"
	GetUserByLastLoginReferrerProtocolRelativePathQueryFragmentMethod     = "GetUserByLastLoginReferrerProtocolRelativePathQueryFragment"
	GetUserByLastLoginReferrerProtocolRelativeHostPathQueryFragmentMethod = "GetUserByLastLoginReferrerProtocolRelativeHostPathQueryFragment"
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
