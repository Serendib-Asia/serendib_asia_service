package dbconfig

// methods
const (
	initDBMethod                   = "initDB"
	initDefaultDBMethod            = "initDefaultDB"
	InitDBConnectionMethod         = "InitDBConnection"
	InitDBConWithAutoMigrateMethod = "InitDBConWithAutoMigrate"
)

// log constants
const (
	OpenGormDBConnection = "Open GORM DB connection"
	InitializingDBConn   = "Initializing CloudSql: %v connection for %s"
)
