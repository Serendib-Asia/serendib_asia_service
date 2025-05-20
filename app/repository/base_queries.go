package repository

// common queries
const (
	// BetweenQuery
	queryBetween = "%s.%s BETWEEN '%s' AND '%s'"
	// IN Query
	inQueryPrefix = "%s.%s IN ("
	// EQUALS Query
	equalsQueryPrefix = "%s.%s ="
	// EQUALS Query int
	equalsQueryPrefixInt = "%s.%s = %v"
	// Year week low limit Query
	queryYearAndWeekLowLimit = "(year > %v OR (year = %v AND week >= %v))"
	// Year week upper limit Query
	queryYearAndWeekUpperLimit = "(year < %v OR (year = %v AND week <= %v))"
	// Less than query
	lessThanQueryPrefix = "%s.%s <"
	// Order by query prefix
	orderByQueryPrefix = "%s %s"

	getTotalRecordCountQuery = `SELECT
		COUNT(*) AS total_count
		FROM
		 %s
		WHERE
		 %s;`

	// incidents queries
	// GetIncidentsByServiceQuery returns all incidents by service
	getIncidentsByServiceBaseQuery = `SELECT 
		incidents.service_id, 
		services.service_name,
		COUNT(*) AS total_count,
		ROUND((COUNT(*) * 100.0 / (
			SELECT COUNT(*) FROM incidents 
				WHERE 
					%s
			)), 2) AS percentage
		FROM 
			incidents
		JOIN 
			services ON services.service_id = incidents.service_id
		WHERE
			%s
		GROUP BY 
			incidents.service_id, services.service_name
		ORDER BY 
			total_count DESC
		LIMIT %v;`

	// getTopIncidentsByPriorityBaseQuery is a base query to get top incidents by priority
	getTopIncidentsByPriorityBaseQuery = `SELECT 
		priorities.priority_name AS priority,
		COUNT(*) AS total_count,
		ROUND((COUNT(*) * 100.0 / 
			(SELECT COUNT(*) FROM incidents
			WHERE
				%s
			)),2) AS percentage
		FROM
			incidents
		JOIN
			priorities
		ON incidents.priority_id = priorities.priority_id
		WHERE
			%s
		GROUP BY
			priorities.priority_name
		ORDER BY
			total_count DESC;`

	// getRestorationMethodBaseQuery is a base query to get restoration method
	getRestorationMethodBaseQuery = `SELECT
			incidents.resolution_method,
			ROUND((COUNT(*) * 100.0 / 
			(SELECT COUNT(*) FROM incidents
			WHERE
				%s
			)),2) AS percentage
		FROM
			incidents
		WHERE
			%s
		GROUP BY
			incidents.resolution_method
		ORDER BY
			percentage DESC;`

	// getManualIncidentsBaseQuery is a base query to get manual incidents
	getManualIncidentsBaseQuery = `SELECT 
			incidents.service_id,
			services.service_name,
			COUNT(*) AS total_count							
		FROM
			incidents
		JOIN
			services ON services.service_id = incidents.service_id
		WHERE
			%s AND incidents.resolution_method = 'manual'
		GROUP BY
			incidents.service_id, services.service_name
		ORDER BY
			total_count DESC;`

	// GetTopIncidentsByCostBaseQuery is a base query to get top incidents by cost
	getTopIncidentsByCostBaseQuery = `SELECT 
			incidents.service_id, 
			services.service_name,
			SUM(incidents.cost) AS total_cost
		FROM 
			incidents
		JOIN 
			services ON services.service_id = incidents.service_id
		WHERE
			%s
		GROUP BY 
			incidents.service_id, services.service_name
		ORDER BY 
			total_cost DESC
		LIMIT %v;`

	// Asset queries
	// getCurrentlyDownBaseQuery is used to get currently down assets
	getCurrentlyDownBaseQuery = `SELECT
			assets.asset_id, 
			assets.asset_type,
			assets.tag,
			assets.hostname as host_name,
			assets.ip_address,
			assets.issue_desc
		FROM
			assets
		WHERE
			%s
		ORDER BY
			assets.asset_id
		LIMIT
			%v
		OFFSET
			%v;`

	// getCurrentlyDownCountBaseQuery is a base query to get currently down count
	getCurrentlyDownCountBaseQuery = `SELECT COUNT(*) FROM assets WHERE %s;`

	// getAssignedTicketsBaseQuery is a base query to get assigned tickets
	getAssignedTicketsBaseQuery = `SELECT 
    	incidents.case_number,
    	priorities.priority_name AS priority,
    	incidents.status,
    	incidents.description,
    	incidents.start_date,
    	agents.agent_name
	FROM 
		incidents
	JOIN 
		agents ON incidents.agent_id = agents.agent_id
	JOIN
		priorities ON incidents.priority_id = priorities.priority_id
	WHERE 
		%s
	ORDER BY 
		start_date DESC
	LIMIT 
		%v
	OFFSET 
		%v;`

	getAssignedTicketsCountBaseQuery = `SELECT COUNT(*) FROM incidents WHERE %s`

	// getSLAMetricsBaseQuery is a base query to get SLA metrics
	getSLAMetricsBaseQuery = `SELECT
			SUM(CASE WHEN first_call_resolved = 'true' THEN 1 ELSE 0 END) AS first_call_resolved,
			ROUND(AVG(DATE(incidents.end_date) - DATE(incidents.start_date))) AS avg_handle_time,
			SUM(CASE WHEN EXTRACT(EPOCH FROM(CURRENT_DATE - incidents.start_date))/3600 > sla_metrics.warning_threshold THEN 1 ELSE 0 END) AS active_sla_warning,
			SUM(CASE WHEN EXTRACT(EPOCH FROM(CURRENT_DATE - incidents.start_date))/3600 > sla_metrics.breach_threshold THEN 1 ELSE 0 END) AS active_sla_breached
		FROM 
			incidents
		JOIN
			clients
		ON
			incidents.client_id = clients.client_id
		JOIN
			sla_metrics
		ON
			sla_metrics.client_id = clients.client_id
		WHERE
			%s`

	// getAlertsQuery is used to get the alerts
	getAlertsQuery = `SELECT 
			alerts.alert_id,
			alerts.start_date,
			alerts.alert_status,
			alerts.message,
			alerts.description,
			alerts.team
		FROM
			alerts
		WHERE
			%s
		ORDER BY
			alerts.start_date DESC
		LIMIT
			%v
		OFFSET
			%v;`

	// getAlertsCountBaseQuery is a base query to get alerts count
	getAlertsCountBaseQuery = `SELECT COUNT(*) FROM alerts WHERE %s`

	// Escalation queries
	// getEscalationsBaseQuery is a base query to get escalation incidents based on the incident type.
	getEscalationsBaseQuery = `SELECT
				incidents.incident_id,
			    incidents.case_number,
				priorities.priority_name AS priority,
				incidents.description,
				incidents.escalated,
				incidents.escalation_date,
				incidents.last_response,
				(CURRENT_DATE - DATE(incidents.last_response)) AS days_since_last_response,
				incidents.predicted,
				incidents.prediction_score,
				incidents.days_open,
				incidents.start_date,
				incidents.churn_risk,
				incidents.agent_id,
				incidents.escalation_requested,
				incidents.predicted_escalation_date,
				incidents.escalation_request_date,
				clients.client_name,
				clients.open_cases_count,
				agents.agent_name
			FROM
				incidents
			JOIN
				clients
			ON
				clients.client_id = incidents.client_id
			JOIN
				agents
			ON
				agents.agent_id = incidents.agent_id
			JOIN
				priorities
			ON
				priorities.priority_id = incidents.priority_id
			WHERE 
				%s
			ORDER BY
				%s
			LIMIT 
				%v 
			OFFSET 
				%v;`

	getEscalationsCountBaseQuery = `SELECT COUNT(*) FROM incidents WHERE %s`

	getEscalationSummaryQuery = `SELECT 
    COUNT(*) AS total_count,
    COUNT(CASE WHEN escalated = FALSE AND predicted = TRUE AND escalation_requested = FALSE THEN 1 END) AS likely_to_escalate_count,
    ROUND(
        (COUNT(CASE WHEN escalated = FALSE AND predicted = TRUE AND escalation_requested = FALSE THEN 1 END) * 100.0) / 
        NULLIF(COUNT(*), 0),
    2) AS percentage
FROM 
    incidents 
WHERE 
    status = 'open'`

	// Agent Queries
	// getAgentCompletedCaseCount is used to get the completed cases count by agent
	getAgentCompletedCaseCount = `SELECT
	        agents.agent_id,
			COUNT(*) AS cases_count
		FROM
			incidents
		JOIN
			clients
		ON
			incidents.client_id = clients.client_id 
		JOIN
			service_company as sc
		ON
			clients.service_company_id = sc.service_company_id
		JOIN
			agents
		ON 
			sc.service_company_id = agents.service_company_id
		WHERE 
			%s
		GROUP BY
			agents.agent_id;`

	// GetAgentSkillsScore
	getAgentSkillScores = `SELECT
			management_skill_score,
			soft_skill_score,
			year,
			week
		FROM 
			agent_performance
		WHERE
			%s
		ORDER BY year,week ASC;`

	// GetAgentStatus
	getAgentStatusDetails = `SELECT 
			percentile_cont(0.5) WITHIN GROUP (ORDER BY agent_performance.problem_open_time_days) AS problem_open_time_days,
			percentile_cont(0.5) WITHIN GROUP (ORDER BY agent_performance.response_time_hours) AS response_times,
			percentile_cont(0.5) WITHIN GROUP (ORDER BY agent_performance.conversation_count) AS conversation_count,
			agents.start_time,
			agents.end_time,
			array_agg(value) AS out_of_office
		FROM 
			agent_performance
		LEFT JOIN 
			agent_schedules ON agent_performance.agent_id = agent_schedules.agent_id
		JOIN 
			agents ON agent_performance.agent_id = agents.agent_id
		LEFT JOIN 
			unnest(agent_schedules.out_of_office) AS t(value) ON agent_schedules.agent_id = agent_performance.agent_id
		WHERE
			%s
		GROUP BY
			agents.start_time,
			agents.end_time;`
	getRunbooksQuery = `SELECT
			run_books.runbook_id,
			run_books.title,
			run_books.description,
			run_books.operating_system,
			run_books.language,
			run_books.script,
			run_books.status,
			run_books.exec_count,
			run_books.agent_id
		FROM
			run_books
		WHERE
			%s
		LIMIT
			%v
		OFFSET
			%v;`

	// backlog queries
	getBacklogIncidentsBaseQuery = `SELECT
			inc.incident_id,
			inc.case_number,
			priorities.priority_name AS priority,
			inc.description,
			inc.escalated,
			inc.escalation_date,
			inc.sla_missed,
			inc.last_outbound_date,
			inc.predicted,
			inc.attention_score,
			inc.prediction_score,
			inc.conversation_count,
			inc.sentiment_score,
			inc.start_date,
			inc.agent_id,
			inc.escalation_requested,
			inc.predicted_escalation_date,
			inc.escalation_request_date,
			inc.view_count,
			inc.client_name,
			agents.agent_name,
			inc.agent_count
		FROM
			(
				SELECT
					incidents.incident_id,
					incidents.case_number,
					incidents.priority_id,
					incidents.description,
					incidents.escalated,
					incidents.escalation_date,
					incidents.sla_missed,
					incidents.last_outbound_date,
					incidents.predicted,
					incidents.attention_score,
					incidents.prediction_score,
					incidents.conversation_count,
					incidents.sentiment_score,
					incidents.start_date,
					incidents.agent_id,
					incidents.escalation_requested,
					incidents.predicted_escalation_date,
					incidents.escalation_request_date,
					incidents.view_count,
					clients.client_name,
					COUNT (DISTINCT(agents.agent_id)) as agent_count
				FROM
					incidents
					JOIN clients ON clients.client_id = incidents.client_id
					JOIN service_company as sc ON clients.service_company_id = sc.service_company_id
					JOIN agents ON sc.service_company_id = agents.service_company_id
				WHERE
					%s
				GROUP BY
					incident_id,
					clients.client_name
			) as inc
		JOIN agents ON agents.agent_id = inc.agent_id
		JOIN priorities ON priorities.priority_id = inc.priority_id
		ORDER BY 
			%s
		LIMIT 
			%v 
		OFFSET 
			%v;`

	getIncidentSummaryBaseQuery = `SELECT
    	COUNT(DISTINCT incidents.incident_id) AS cases,
    	COUNT(DISTINCT incidents.agent_id) AS agents
	FROM
    	incidents
	WHERE
		%s;`

	// Sentiment queries

	// getSentimentMetricsPercentageQuery is used to get the sentiment metrics percentage
	getSentimentMetricsPercentageQuery = `SELECT
    (AVG(client_sentiments.sla_score) * 100.0 / 100) AS sla_percentage,
	(AVG(client_sentiments.efficiency_score) * 100.0 / 100) AS efficiency_percentage,
    (AVG(client_sentiments.customer_experience_Score) * 100.0 / 100) AS customer_experience_percentage,
    (AVG(client_sentiments.churn_risk) * 10.0 / 10) AS churn_risk_percentage
	FROM
   		client_sentiments;`

	// Sentiment score queries
	// getClientSentimentScoreBaseQuery is used to get the client sentiment scores
	getClientSentimentScoreBaseQuery = `SELECT
		client_id as id,
		DATE_TRUNC('month', date) as month,
		ROUND(AVG(%s), 2) as score
	FROM
		client_sentiments
	WHERE
		%s
	GROUP BY
		client_id,
		month
	ORDER BY
		month ASC;`

	// getUserSentimentScoreBaseQuery is used to get the user sentiment scores
	getUserSentimentScoreBaseQuery = `SELECT
		user_id as id,
		DATE_TRUNC('month', date) as month,
		ROUND(AVG(%s), 2) as score
	FROM
		user_sentiments
	WHERE
		%s
	GROUP BY
		user_id,
		month
	ORDER BY
		month ASC;`

	// Customer Insights queries
	// getTicketSummaryBaseQuery is a base query to get ticket summary
	getTicketSummaryBaseQuery = `SELECT
			SUM(
				CASE
					WHEN incidents.escalated = TRUE THEN 1
					ELSE 0
				END
			) AS escalations,
			clients.critical_cases_count AS critical_cases,
			clients.negative_sentiments_count AS negative_sentiments,
			clients.positive_sentiments_count AS positive_sentiments
	FROM
		incidents
	FULL OUTER JOIN 
		clients 
	ON 
		clients.client_id = incidents.client_id
	WHERE
		%s
	GROUP BY
		critical_cases,
		negative_sentiments,
		positive_sentiments;`

	// getCustomerScoreIndicatorBaseQuery is a base query to get customer score indicator
	getCustomerScoreIndicatorBaseQuery = `SELECT
			clients.client_name,
			clients.start_date AS active_since,
			ROUND(AVG(client_sentiments.customer_experience_score) * 100.0 / 100, 2) AS customer_experience
	FROM
    	clients
    FULL OUTER JOIN 
		client_sentiments 
	ON 
		clients.client_id = client_sentiments.client_id
	WHERE
    	%s
	GROUP BY
    	clients.client_name,
    	active_since;`

	// getResponseTimeMetricsBaseQuery is a base query to get response time metrics
	getResponseTimeMetricsBaseQuery = `SELECT
    -- Total unresponded tickets
    COUNT(incidents.incident_id) AS total_unresponded_tickets,

    -- Overdue tickets (Beyond SLA Deadline)
    COUNT(
        CASE
            WHEN (incidents.start_date + response_times.response_time) < CURRENT_TIMESTAMP
            AND (
                (incidents.priority_id = 3 AND EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - (incidents.start_date + response_times.response_time))) / 60 > 120) OR
                (incidents.priority_id = 2 AND EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - (incidents.start_date + response_times.response_time))) / 60 > 240) OR
                (incidents.priority_id = 1 AND EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - (incidents.start_date + response_times.response_time))) / 60 > 360)
            )
            THEN 1
        END
    ) AS overdue_tickets,

    -- Tickets approaching SLA deadline
    COUNT(
        CASE
            WHEN (incidents.start_date + response_times.response_time) >= CURRENT_TIMESTAMP
            AND (
                (incidents.priority_id = 3 AND EXTRACT(EPOCH FROM ((incidents.start_date + response_times.response_time) - CURRENT_TIMESTAMP)) / 60 BETWEEN 60 AND 120) OR
                (incidents.priority_id = 2 AND EXTRACT(EPOCH FROM ((incidents.start_date + response_times.response_time) - CURRENT_TIMESTAMP)) / 60 BETWEEN 180 AND 240) OR
                (incidents.priority_id = 1 AND EXTRACT(EPOCH FROM ((incidents.start_date + response_times.response_time) - CURRENT_TIMESTAMP)) / 60 BETWEEN 300 AND 360)
            )
            THEN 1
        END
    ) AS approaching_deadline_tickets,

    -- Tickets in early response window
    COUNT(
        CASE
            WHEN EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - incidents.start_date)) / 60 < 
                 CASE 
                     WHEN incidents.priority_id = 3 THEN 60 
                     WHEN incidents.priority_id = 2 THEN 180 
                     WHEN incidents.priority_id = 1 THEN 300 
                 END
            THEN 1
        END
    ) AS early_response_tickets,

    -- Uncategorized/custom SLA tickets
    COUNT(
        CASE
            WHEN incidents.priority_id NOT IN (1, 2, 3)
            THEN 1
        END
    ) AS uncategorized_tickets
FROM
    incidents
LEFT JOIN
    response_times ON response_times.incident_id = incidents.incident_id
WHERE %s 
	AND response_times.response_time IS NOT NULL;`

	// getCasesSummaryBaseQuery is a base query to get cases summary
	getCasesSummaryBaseQuery = `SELECT
		incidents.incident_id AS request_id,
    	incidents.case_number,
    	incidents.description,
    	incidents.start_date AS start_date,
    	priorities.priority_name AS priority,
    	incidents.sla_missed,
		incidents.churn_risk,
		clients.client_id,
    	clients.client_name,
		services.service_name
	FROM
    	incidents
	JOIN 
		clients 
	ON 
		incidents.client_id = clients.client_id
    JOIN 
		services 
	ON 
		services.service_id = incidents.service_id
	JOIN
	 	priorities
	ON
		priorities.priority_id = incidents.priority_id
	WHERE
    	%s
    ORDER BY 
		start_date DESC
	LIMIT
		%v
	OFFSET
		%v;`

	// getCasesCountBaseQuery is a base query to get cases count
	getCasesCountBaseQuery = `SELECT COUNT(*) FROM incidents WHERE %s`

	// getNumberOfUserAccountsForClientBaseQuery is a base query to get number of user accounts for client
	getNumberOfUserAccountsForClientBaseQuery = `SELECT
    	COUNT(users.user_id) AS number_of_user_accounts
	FROM
    	users
	WHERE
    	%s;`

	getNumberOfReportersForClientBaseQuery = `SELECT
    	COUNT(DISTINCT incidents.user_id) AS number_of_reporters
	FROM
    	incidents
	JOIN
		users
	ON
		incidents.user_id = users.user_id
	JOIN
		clients
	ON
		users.client_id = clients.client_id
	WHERE
    	%s`

	// getUserIncidentDetailsForClientBaseQuery is a base query to get account reporter insights
	getUserIncidentDetailsForClientBaseQuery = `SELECT
		incidents.incident_id,
		incidents.case_number,
		incidents.description,
		incidents.churn_risk,
		incidents.start_date,
		incidents.user_id,
		users.full_name,
		users.email,
		users.contact_number,
		clients.client_id,
		roles.role_name
	FROM
    	incidents
    JOIN 
		users 
	ON 
		users.user_id = incidents.user_id
    JOIN 
		clients 
	ON 
		users.client_id = clients.client_id
    JOIN 
		user_roles 
	ON 
		users.user_id = user_roles.users_user_id
    JOIN 
		roles 
	ON 
		user_roles.roles_role_id = roles.role_id
	WHERE
    	%s
	ORDER BY
    	incidents.incident_id DESC
	LIMIT
		%v
	OFFSET
		%v;`

	// getIncidentsCountBaseQuery is a base query to get incidents count
	getIncidentsCountBaseQuery = `SELECT
		COUNT(*)
	FROM
    	incidents
    JOIN 
		users 
	ON 
		users.user_id = incidents.user_id
    JOIN 
		clients 
	ON 
		users.client_id = clients.client_id
	WHERE
    	%s;`

	// getSentimentHistoryBaseQuery is a base query to get sentiment history
	getSentimentHistoryBaseQuery = `SELECT
	DATE_TRUNC('month', date) AS period,
	ROUND(AVG(overall_score),2) as positive_score_percentage
FROM
	client_sentiments
WHERE
	%s
GROUP BY
	month
ORDER BY
	month ASC;`

	// getNotesBaseQuery is a base query to get notes
	getNotesBaseQuery = `SELECT
	    notes.created_at,
    	notes.note_id,
    	notes.author_id,
		notes.client_id,
    	notes.content,
    	agents.agent_name
	FROM
    	notes
	JOIN
    	agents
	ON
    	notes.author_id = agents.agent_id
	WHERE
    	%s
	ORDER BY
		notes.created_at DESC
	LIMIT
		%v
	OFFSET
		%v;`

	getNotesCountBaseQuery = `SELECT COUNT(*) FROM notes WHERE %s`

	// getClientListQuery is a query to get the client list
	getClientListQuery = `SELECT
		client_id,
		client_name,
		email as client_email
	FROM
		clients;`

	// getClientsWithUsersQuery is a base query to get clients with users
	getClientsWithUsersQuery = `SELECT
		clients.client_id,
		clients.client_name,
		clients.email as client_email,
		users.user_id,
		users.full_name,
		users.email as user_email
	FROM
		clients
	JOIN
		users
	ON
		clients.client_id = users.client_id
	GROUP BY
		clients.client_id,
		users.user_id
	ORDER BY
		clients.client_id ASC;`

	getTotalBacklogIncidentsBaseQuery = `SELECT
		COUNT(*) AS total_incidents
	FROM
		incidents
	WHERE
		%s;`

	getClientQuery = "SELECT client_id FROM clients WHERE client_id = ?"

	getDailyIncidentsTrendBaseQuery = `SELECT
		start_date AS date,
		COUNT(*) AS total_incidents
	FROM
		incidents
	WHERE
		%s
	GROUP BY
		date
	ORDER BY
		date ASC;`

	getWeeklyIncidentsTrendBaseQuery = `SELECT
		DATE_TRUNC('week', start_date) AS date,
		COUNT(*) AS total_incidents
	FROM
		incidents
	WHERE
		%s
	GROUP BY
		date
	ORDER BY
		date ASC;`

	getMonthlyIncidentsTrendBaseQuery = `SELECT
		DATE_TRUNC('month', start_date) AS date,
		COUNT(*) AS total_incidents
	FROM
		incidents
	WHERE
		%s
	GROUP BY
		date
	ORDER BY
		date ASC;`
)
