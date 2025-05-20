package dto

// CommonFilterRequest is a struct that represents common filters for requests.
type CommonFilterRequest struct {
	StartDate string   `json:"start_date" validate:"required,max=10,date"`
	EndDate   string   `json:"end_date" validate:"required,max=10,date"`
	Priority  []string `json:"priority" validate:"required,min=1"`
	Client    []int    `json:"client"`
	Category  []int    `json:"category" validate:"required,min=1"`
	Service   []int    `json:"service"`
}

type DateRangeRequest struct {
	StartDate string `json:"start_date" validate:"required,max=10,date"`
	EndDate   string `json:"end_date" validate:"required,max=10,date"`
}

type PaginatedDateRangeRequest struct {
	DateRangeRequest
	PaginationRequest
}

// TopIncidentsRequest is a struct that represents a request for top incidents.
type TopIncidentsRequest struct {
	CommonFilterRequest
	Limit int `json:"limit" validate:"required,min=1"`
}

// PaginatedCommonFilterRequest is a struct that represents a request for paginated common filters.
type PaginatedCommonFilterRequest struct {
	CommonFilterRequest
	PaginationRequest
}

type IncidentsTrendRequest struct {
	CommonFilterRequest
	Frequency string `json:"frequency" validate:"omitempty,oneof=daily weekly monthly"`
}

type DateTrendRequest struct {
	DateRangeRequest
	Frequency string `json:"frequency" validate:"omitempty,oneof=daily weekly monthly"`
}

// EscalationRequest is a struct that represents a request for escalation.
type EscalationRequest struct {
	PaginationRequest
	IsPredicted         string     `json:"is_predicted" validate:"oneOfStatus"`
	IsEscalated         string     `json:"is_escalated" validate:"oneOfStatus"`
	IsEscalationRequest string     `json:"is_escalation_request" validate:"oneOfStatus"`
	Sorting             SortParams `json:"sorting" validate:"required"`
}

type SortParams struct {
	SortField string `json:"sort_field" validate:"required,oneOfSortFields" enums:"incident_id,start_date,last_response,predicted_escalation_date,escalation_date,escalation_request_date"`
	SortOrder string `json:"sort_order" validate:"required,oneOfSortOrders" enums:"asc,desc,ASC,DESC"`
}

type AgentCompletedCasesCountRequest struct {
	AgentID   uint   `json:"agent_id" validate:"required"`
	StartDate string `json:"start_date" validate:"required,max=10,date"`
	EndDate   string `json:"end_date" validate:"required,max=10,date"`
}

type AgentInfoRequest struct {
	AgentID   uint `json:"agent_id" validate:"required"`
	StartYear int  `json:"start_year" validate:"required,year"`
	StartWeek int  `json:"start_week" validate:"required"`
	EndYear   int  `json:"end_year" validate:"required,year"`
	EndWeek   int  `json:"end_week" validate:"required"`
}

type AuthUserRequest struct {
	AuthToken string
}

// PaginationRequest is a struct that represents a request for pagination.
type PaginationRequest struct {
	Limit  int `json:"limit" validate:"required,min=1"`
	Offset int `json:"offset" validate:"min=0"`
}

// RunbookFilterRequest is a struct that represents a request for runbook filters.
type RunbookFilterRequest struct {
	PaginationRequest
	Search          string `json:"search"`
	OperatingSystem string `json:"operating_system"`
	Status          string `json:"status" validate:"oneOfStatus"`
	AgentID         uint   `json:"agent_id" validate:"required"`
}

// RunbookCreateRequest is a struct that represents a request for creating a runbook.
type RunbookCreateRequest struct {
	Title           string `json:"title" validate:"required,omitEmpty,max=255"`
	Description     string `json:"description" validate:"required,omitEmpty"`
	OperatingSystem string `json:"operating_system" validate:"required,omitEmpty,max=20"`
	AuthToken       string `json:"token" swaggerignore:"true"`
	UserID          string `json:"user_id" swaggerignore:"true"`
	AgentID         uint   `json:"agent_id" validate:"required"`
	Language        string `json:"language" validate:"required,omitEmpty,max=20"`
	Script          string `json:"script" validate:"required,omitEmpty"`
}

// RunbookActivationRequest is a struct that represents a request for activating a runbook.
type RunbookActivationRequest struct {
	RunbookID int  `json:"runbook_id" validate:"required,min=1"`
	Status    bool `json:"status"`
}

type BacklogCommonRequestFields struct {
	PaginationRequest
	Client    []int  `json:"client"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type BacklogRequest struct {
	BacklogCommonRequestFields
	Categories BacklogCategories `json:"categories"`
}

type InactiveIncidentsRequest struct {
	BacklogCommonRequestFields
}

type IncidentSummaryRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Client    []int  `json:"client"`
}

type BacklogCategories struct {
	MostViewed          bool `json:"most_viewed"`
	InactiveCases       bool `json:"inactive_cases"`
	LastOutbound        bool `json:"last_outbound"`
	UniqueResponders    bool `json:"unique_responders"`
	PriorityByAttnScore bool `json:"priority_by_attn_score"`
	MostConversations   bool `json:"most_conversations"`
	SentimentScore      bool `json:"sentiment_score"`
	AgingCases          bool `json:"aging_cases"`
}

type SentimentScoreCategories struct {
	Overall            bool `json:"overall"`
	SLA                bool `json:"sla"`
	CustomerExperience bool `json:"customer_experience"`
	Efficiency         bool `json:"efficiency"`
	ChurnRisk          bool `json:"churn_risk"`
}

type ClientSentimentScoreRequest struct {
	Client     []int                    `json:"client" validate:"required,min=1"`
	Categories SentimentScoreCategories `json:"categories"`
	StartDate  string                   `json:"start_date" validate:"required,max=10,date"`
	EndDate    string                   `json:"end_date" validate:"required,max=10,date"`
}

type UserSentimentScoreRequest struct {
	User       []int                    `json:"user" validate:"required,min=1"`
	Categories SentimentScoreCategories `json:"categories"`
	StartDate  string                   `json:"start_date" validate:"required,max=10,date"`
	EndDate    string                   `json:"end_date" validate:"required,max=10,date"`
}

type ClientRequest struct {
	Client []int `json:"client"`
}

type PaginatedClientRequest struct {
	Client []int `json:"client"`
	PaginationRequest
}

type ClientDateRangeRequest struct {
	Client    []int  `json:"client"`
	Period    string `json:"period" validate:"required,oneOfPeriod"`
	StartDate string `json:"start_date" validate:"required,date"`
	EndDate   string `json:"end_date" validate:"required,date"`
}

type CreateNoteRequest struct {
	Client    []int  `json:"client" validate:"required,min=1"`
	AgentID   uint   `json:"agent_id" validate:"required"`
	AuthToken string `json:"token" swaggerignore:"true"`
	UserID    string `json:"user_id" swaggerignore:"true"`
	Content   string `json:"content" validate:"required,omitEmpty,min=3,max=255"`
}

type RunbookCodeGenRequest struct {
	Type           string `json:"type"`
	CodePrompt     string `json:"code_prompt" validate:"required"`
	SourceLanguage string `json:"source_language" validate:"required"`
}

type ClientValidationRequest struct {
	Client []int `json:"client" validate:"required,min=1"`
}

type SecSrvUpdateTenant struct {
	UserID       uint   `json:"userID"`
	TenantName   string `json:"tenantName"`
	TenantUserID string `json:"tenantUserID"`
}

type ValidateAgentsReq struct {
	AuthServerID     string `json:"authServerID" validate:"required,alphaNumeric"`
	FirstName        string `json:"firstName" validate:"alpha"`
	LastName         string `json:"lastName" validate:"alpha"`
	TenantID         string `json:"tenantID"`
	Email            string `json:"email" validate:"required,email"`
	Authorization    string
	ServiceCompanyID string `json:"serviceCompanyID" validate:"required,alphaNumeric"`
	User             UserResponse
}

type DomainInstance struct {
	Domain     string `json:"domain"`
	InstanceID int32  `json:"instance_id"`
}
