package repository

// Context use to create repository context
type Context struct {
	_         struct{}
	RequestID string
}

// CreateRepositoryContext create repository context
func CreateRepositoryContext(requestID string) Context {
	return Context{
		RequestID: requestID,
	}
}
