package repository

// Enum types
const (
	EnumLabel = "enumlabel"
)

// Table names
const (
	ClientSentiments = "client_sentiments"
	Services         = "Services"
	Agents           = "agents"
	Clients          = "clients"
	Categories       = "Categories"
	Incidents        = "incidents"
	Alerts           = "Alerts"
	Assets           = "Assets"
	AgentPerformance = "agent_performance"
	Runbooks         = "run_books"
	UserSentiments   = "user_sentiments"
	Users            = "users"
	Notes            = "notes"
	Instance         = "instance"
	Request          = "request"
	Priorities       = "Priorities"
	Customer         = "customer"
)

// Fields
const (
	Email = "Email"
)

// Sdesk table names
const (
	RequestAnalysis = "analysis"
)

const (
	SDesk = "sdesk"
)

const (
	// Bootstrap repository methods
	GetEnumValuesMethod = "GetEnumValues"

	// MasterData repository methods
	MasterDataRepositoryFindOverallScoreMethod  = "MasterDataRepositoryFindOverallScoreMethod"
	MasterDataRepositoryFindAllServicesMethod   = "MasterDataRepositoryFindAllServices"
	MasterDataRepositoryFindAllClientsMethod    = "MasterDataRepositoryFindAllClients"
	MasterDataRepositoryFindAllCategoriesMethod = "MasterDataRepositoryFindAllCategories"
	MasterDataRepositoryFindAllPrioritiesMethod = "MasterDataRepositoryFindAllPriorities"

	// Service repository methods
	ServiceRepositoryFindAllMethod = "ServiceRepositoryFindAll"

	// Agent repository methods
	AgentFindAllMethod               = "AgentRepositoryFindAll"
	AgentGetCompletedCaseCountMethod = "AgentGetCompletedCaseCount"
	AgentGetSkillsMethod             = "AgentGetSkills"
	AgentGetStatusDetailsMethod      = "AgentGetStatusDetails"
	FindAgentInfoByIDMethod          = "FindAgentInfoByID"
	CreateAgentMethod                = "CreateAgent"
	GetExistingAgentMethod           = "GetExistingAgent"
	AddAgentMethod                   = "AddAgent"

	// Client repository methods
	ClientRepositoryFindAllMethod          = "ClientRepositoryFindAll"
	ClientRepositoryFindOverallScoreMethod = "ClientRepositoryFindOverallScoreMethod"
	FindClientListMethod                   = "FindClientList"
	FindClientsWithUsersMethod             = "FindClientsWithUsers"
	GetClientByClientIDMethod              = "GetClientByClientID"

	// Category repository methods
	CategoryRepositoryFindAllMethod = "CategoryRepositoryFindAll"

	// Incident repository methods
	FindFilteredIncidentsByServiceMethod       = "FindFilteredIncidentsByService"
	FindIncidentsByCostMethod                  = "FindIncidentsByCost"
	GetSQLQueryForIncidentByServiceMethod      = "getSQLQueryForIncidentByService"
	GetSQLQueryForTopIncidentsByPriorityMethod = "getSQLQueryForTopIncidentsByPriority"
	IncidentRepositoryFindByCostMethod         = "IncidentRepositoryFindByCost"
	FindManualIncidentStatusesMethod           = "FindManualIncidentStatuses"
	FindRestorationStatusMethod                = "FindRestorationStatus"
	FindIncidentsTrendDailyMethod              = "FindIncidentsTrendDaily"
	FindIncidentsTrendWeeklyMethod             = "FindIncidentsTrendWeekly"
	FindIncidentsTrendMonthlyMethod            = "FindIncidentsTrendMonthly"
	FindIncidentsByPriorityMethod              = "FindIncidentsByPriority"
	FindAllIncidentsByPriorityMethod           = "FindAllIncidentsByPriority"
	FindAllIncidentsByServiceMethod            = "FindAllIncidentsByService"
	FindAllIncidentsTrendDailyMethod           = "FindAllIncidentsTrendDaily"
	FindAllIncidentsTrendWeeklyMethod          = "FindAllIncidentsTrendWeekly"
	FindAllIncidentsTrendMonthlyMethod         = "FindAllIncidentsTrendMonthly"
	FetchAllManualIncidentsStatusMethod        = "FetchAllManualIncidentsStatus"
	FetchAllRestorationStatusMethod            = "FetchAllRestorationStatus"

	// Alert repository methods
	FindAlertsMethod     = "FindAlerts"
	GetAlertsCountMethod = "GetAlertsCount"

	// asset repository methods
	FindCurrentlyDownMethod           = "FindCurrentlyDown"
	GetSQLQueryForCurrentlyDownMethod = "getSQLQueryForCurrentlyDown"
	GetCurrentlyDownCountMethod       = "GetCurrentlyDownCount"

	// Tickets repository methods
	FindFilteredAssignedTicketsMethod = "FindFilteredAssignedTickets"
	GetAssignedTicketsCountMethod     = "GetAssignedTicketsCount"
	FindAllAssignedTicketsMethod      = "FindAllAssignedTickets"
	GetAllAssignedTicketsCountMethod  = "GetAllAssignedTicketsCount"

	// Metrics repository methods
	FindFilteredSLAMetricsMethod = "FindFilteredSLAMetrics"
	FindAllSLAMetricsMethod      = "FindAllSLAMetrics"

	// Escalation repository methods
	FindEscalationsMethod       = "FindEscalations"
	GetEscalationsCountMethod   = "GetEscalationsCount"
	FindEscalationSummaryMethod = "FindEscalationSummary"
	// Escalation repository methods -private
	GenerateEscalatedScenarioQueryMethod         = "generateEscalatedScenarioQuery"
	GenerateEscalationRequestScenarioQueryMethod = "generateEscalationRequestScenarioQuery"
	GenerateLikelyToEscalateScenarioQueryMethod  = "generateLikelyToEscalateScenarioQuery"
	GenerateWhereClauseListMethod                = "generateWhereClauseList"

	// Runbook repository methods
	RunbookFindByIDMethod                      = "RunbookFindByID"
	RunbookFindByFilterMethod                  = "RunbookFind"
	RunbookCreateMethod                        = "RunbookCreate"
	RunbookSetActiveMethod                     = "RunbookSetActive"
	GetRunbookRecordCountByFilterRequestMethod = "RunbookGetRunbookRecordCount"

	// Backlog repository methods
	FindBacklogIncidentsMethod             = "FindBacklogIncidents"
	FindIncidentSummaryMethod              = "FindClientIncidentSummary"
	FindIncidentsByViewCountMethod         = "FindIncidentsByViewCount"
	FindInactiveIncidentsMethod            = "FindInactiveIncidents"
	FindUniqueResponderIncidentsMethod     = "FindUniqueResponderIncidents"
	FindIncidentsByLastOutboundMethod      = "FindIncidentsByLastOutbound"
	FindIncidentsByUniqueRespondersMethod  = "FindIncidentsByUniqueResponders"
	FindIncidentsByAttentionScoreMethod    = "FindIncidentsByAttentionScore"
	FindIncidentsByConversationCountMethod = "FindIncidentsByConversationCount"
	FindIncidentsBySentimentScoreMethod    = "FindIncidentsBySentimentScore"
	FindAgingIncidentsByLastOutboundMethod = "FindAgingIncidentsByLastOutbound"

	// Sentiment repository methods
	FindSentimentMetricsMethod = "FindSentimentMetrics"
	// Client sentiment score methods
	FindClientOverallSentimentScoreMethod            = "FindClientOverallSentimentScore"
	FindClientSLASentimentScoreMethod                = "FindClientSLASentimentScore"
	FindClientEfficiencySentimentScoreMethod         = "FindClientEfficiencySentimentScore"
	FindClientCustomerExperienceSentimentScoreMethod = "FindClientCustomerExperienceSentimentScore"
	FindClientChurnRiskSentimentScoreMethod          = "FindClientChurnRiskSentimentScore"
	// User sentiment score methods
	FindUserOverallSentimentScoreMethod            = "FindUserOverallSentimentScore"
	FindUserSLASentimentScoreMethod                = "FindUserSLASentimentScore"
	FindUserEfficiencySentimentScoreMethod         = "FindUserEfficiencySentimentScore"
	FindUserCustomerExperienceSentimentScoreMethod = "FindUserCustomerExperienceSentimentScore"
	FindUserChurnRiskSentimentScoreMethod          = "FindUserChurnRiskSentimentScore"

	// Customer insights repository methods
	FindTicketSummaryMethod                  = "FindTicketSummary"
	FindCustomerScoreIndicatorMethod         = "FindCustomerScoreIndicator"
	FindInternalCustomerScoreIndicatorMethod = "FindInternalCustomerScoreIndicator"
	FindNumberOfUserAccountsForClientMethod  = "FindNumberOfUserAccountsForClient"
	FindNumberOfReportersForClientMethod     = "FindNumberOfReportersForClient"
	FindUserIncidentDetailsForClientMethod   = "FindUserIncidentDetailsForClient"
	FindResponseTimeMetricsMethod            = "FindResponseTimeMetrics"
	FindCasesSummaryMethod                   = "FindCaseSummary"
	FindSentimentHistoryMethod               = "FindSentimentHistory"
	CreateNoteRepoMethod                     = "CreateNote"
	FindNotesMethod                          = "FindNotes"
	FindTotalIncidentsCountMethod            = "FindTotalIncidentsCount"
	FindTotalInactiveIncidentsCountMethod    = "FindTotalInactiveIncidentsCount"
	GetNotesCountMethod                      = "GetNotesCount"
	GetCasesCountMethod                      = "GetCasesCount"
	GetIncidentsCountMethod                  = "GetIncidentsCount"
	FindEscalatedTicketsCountMethod          = "FindEscalatedTicketsCount"
	// private
	GetPeriodFieldMethod = "getPeriodField"

	// Instance repository methods
	GetInstanceDetailsMethod = "GetInstanceDetails"
)
