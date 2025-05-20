package repository

// Column names
const (
	clientIDColumn                     = "client_id"
	priorityIDColumn                   = "priority_id"
	startDateColumn                    = "start_date"
	escalatedColumn                    = "escalated"
	predictedColumn                    = "predicted"
	slaMissedColumn                    = "sla_missed"
	escalationRequestedColumn          = "escalation_requested"
	agentIDColumn                      = "agent_id"
	statusColumn                       = "status"
	runbookIDColumn                    = "runbook_id"
	titleColumn                        = "title"
	operatingSystemColumn              = "operating_system"
	lastOutboundDateColumn             = "last_outbound_date"
	attentionScoreColumn               = "attention_score"
	conversationCountColumn            = "conversation_count"
	sentimentScoreColumn               = "sentiment_score"
	agentCountColumn                   = "agent_count"
	dateColumn                         = "date"
	overallScoreColumn                 = "overall_score"
	slaScoreColumn                     = "sla_score"
	efficiencyScoreColumn              = "efficiency_score"
	customerExperienceScoreColumn      = "customer_experience_score"
	churnRiskColumn                    = "churn_risk"
	userIDColumn                       = "user_id"
	overallPercentageColumn            = "overall_percentage"
	slaPercentageColumn                = "sla_percentage"
	efficiencyPercentageColumn         = "efficiency_percentage"
	customerExperiencePercentageColumn = "customer_experience_percentage"
	churnRiskPercentage                = "churn_risk_percentage"
	viewCountColumn                    = "view_count"
	categoryIDColumn                   = "category_id"
	serviceIDColumn                    = "service_id"
)

const (
	ifEqualRunbookID = "runbook_id=?"
)

// SQL Operators
const (
	andOperator         = "AND"
	orOperator          = "OR"
	iLIKEOperator       = "%s ILIKE '%%%s%%'"
	equalValOperator    = "%s = '%s'"
	equalValOperatorInt = "%s = %d"
)

const (
	GenerateInQueryMethod           = "GenerateInQuery"
	CombineQueriesMethod            = "CombineQueries"
	GenerateBetweenQueryMethod      = "GenerateBetweenQuery"
	buildSQLWhereClauseMethod       = "buildSQLWhereClause"
	BuildSQLQueryMethod             = "buildSQLQuery"
	getSQLQueryMethod               = "getSQLQueryMethod"
	generateDateIntervalQueryMethod = "generateDateIntervalQuery"
	generateOrderByClauseMethod     = "generateOrderByClause"
	GetExistingUserMethod           = "GetExistingUserMethod"
)

// log const
const (
	ClientID             = "ClientID"
	GeneratedQuery       = "Generated query"
	GeneratedWhereClause = "Generated where clause"
)

// query clauses
const (
	CommaWithSpace              = ", "
	StringPlaceholder           = "'%s'"
	StringPlaceholderWithSpaces = " %s "
	ValuePlaceholderWithSpaces  = " %v "
	ValuePlaceholder            = "%v"
	CloseBracket                = ")"
	closed                      = "closed"
)

const (
	daysInterval = 30
	ascending    = "ASC"
	descending   = "DESC"
)
