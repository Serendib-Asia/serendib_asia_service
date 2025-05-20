package dto

import (
	"time"
)

// UserResponse represents the response for a user.
type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AgentIdentityDetails represents the response for a agent identity details response.
type AgentIdentityDetails struct {
	AgentID    uint      `json:"agent_id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	JoinedDate time.Time `json:"joined_date"`
}

// AgentIdentityDetailsList represents the list of agent identity details responses.
type AgentIdentityDetailsList []AgentIdentityDetails

// AgentCompletedCasesCount represents the response for the completed incidents count which completed by agent.
type AgentCompletedCasesCount struct {
	AgentID    uint  `json:"agent_id"`
	CasesCount int64 `json:"cases_count"`
}

type AgentSkillScore struct {
	ManagementSkillScore float64 `json:"management_skill_score"`
	SoftSkillScore       float64 `json:"soft_skill_score"`
	Week                 int64   `json:"week"`
	Year                 int64   `json:"year"`
	WeekStartDate        string  `json:"week_start_date"`
}

type AgentSkillScoreList []AgentSkillScore

type AgentStatusResponse struct {
	ProblemOpenTimeDays float64             `json:"problem_open_time_days"`
	ResponseTimes       float64             `json:"response_times"`
	ConversationCount   float64             `json:"conversation_count"`
	StartTime           string              `json:"start_time"`
	EndTime             string              `json:"end_time"`
	OutOfOffice         map[string][]string `json:"out_of_office"`
}

type AgentStatusResult struct {
	ProblemOpenTimeDays float64 `json:"problem_open_time_days"`
	ResponseTimes       float64 `json:"response_times"`
	ConversationCount   float64 `json:"conversation_count"`
	StartTime           string  `json:"start_time"`
	EndTime             string  `json:"end_time"`
	OutOfOffice         string  `json:"out_of_office"`
}

// Tenant represents the response for tenant validation.
type Tenant struct {
	TenantID         string `json:"tenant_id"`
	TenantName       string `json:"tenant_name"`
	TenantUserID     string `json:"tenant_user_id"`
	ServiceCompanyID string `json:"service_company_id"`
}
