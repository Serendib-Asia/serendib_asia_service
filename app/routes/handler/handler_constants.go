package handler

// Constants for handler methods.
const (
	// Common handlers
	HandleStaticDocMethod      = "HandleStaticDoc"
	GetMasterDataHandlerMethod = "GetMasterDataHandler"

	// agent handlers
	GetAgentListHandlerMethod               = "GetAgentListHandler"
	GetAgentCompletedCaseCountHandlerMethod = "GetAgentCompletedCaseCountHandler"
	GetAgentSkillScoresHandlerMethod        = "GetAgentSkillScoresHandler"
	AgentStatusHandlerMethod                = "AgentStatusHandler"
	ValidateAgentHandlerMethod              = "ValidateAgentHandler"

	// property handlers
	CreatePropertyHandlerMethod = "CreatePropertyHandler"
	GetPropertyHandlerMethod    = "GetPropertyHandler"
	UpdatePropertyHandlerMethod = "UpdatePropertyHandler"
	DeletePropertyHandlerMethod = "DeletePropertyHandler"
	ListPropertiesHandlerMethod = "ListPropertiesHandler"

	// Lookup handler method constants
	GetPurposeTypesHandlerMethod       = "GetPurposeTypesHandler"
	GetPropertyTypesHandlerMethod      = "GetPropertyTypesHandler"
	GetFurnitureTypesHandlerMethod     = "GetFurnitureTypesHandler"
	GetPropertyConditionsHandlerMethod = "GetPropertyConditionsHandler"
	GetUtilitiesHandlerMethod          = "GetUtilitiesHandler"
	GetAmenitiesHandlerMethod          = "GetAmenitiesHandler"

	// User handler method constants
	RegisterHandlerMethod       = "RegisterHandler"
	LoginHandlerMethod          = "LoginHandler"
	GetProfileHandlerMethod     = "GetProfileHandler"
	UpdateProfileHandlerMethod  = "UpdateProfileHandler"
	UpdatePasswordHandlerMethod = "UpdatePasswordHandler"

	// Favorite handler methods
	AddToFavoritesHandlerMethod      = "AddToFavoritesHandler"
	RemoveFromFavoritesHandlerMethod = "RemoveFromFavoritesHandler"
	ListFavoritesHandlerMethod       = "ListFavoritesHandler"

	// Image handler methods
	UploadImageHandlerMethod     = "UploadImageHandler"
	DeleteImageHandlerMethod     = "DeleteImageHandler"
	SetPrimaryImageHandlerMethod = "SetPrimaryImageHandler"
	ListImagesHandlerMethod      = "ListImagesHandler"
)

const (
	dotFwdSlash   = "./"
	htmlString    = "html"
	staticDocName = "static.html"
	docFileName   = "document.html"
)

// Constants for query parameters.
const (
	QueryAllKey    = "queryAll"
	DateRequestKey = "dateRequest"
)
