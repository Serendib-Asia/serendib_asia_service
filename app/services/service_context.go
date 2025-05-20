package services

// ServiceContext used to define service interface
type ServiceContext struct {
	_         struct{}
	RequestID string
}

// CreateServiceContext use to create service context
func CreateServiceContext(requestID string) ServiceContext {
	return ServiceContext{RequestID: requestID}
}
