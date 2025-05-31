package dbconfig

import (
	"fmt"

	"github.com/chazool/serendib_asia_service/pkg/config"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbCon *gorm.DB

// GetDBConnection returns the current database connection
func GetDBConnection() *gorm.DB {
	return dbCon
}

// SetDBConnection sets the current database connection
func SetDBConnection(db *gorm.DB) {
	dbCon = db
}

// InitDBConnection initializes the database connection
// This function initializes the database connection based on the configuration
// It supports both PostgreSQL and MySQL databases
func InitDBConnection() (err error) {
	log.Logger.Debug(log.TraceMsgFuncStart(InitDBConnectionMethod))
	defer log.Logger.Debug(log.TraceMsgFuncEnd(InitDBConnectionMethod))

	// Initialize Default
	err = initDefaultDB()
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(initDefaultDBMethod), zap.Error(err))
		return fmt.Errorf("failed to connect to default database: %w", err)
	}

	return nil
}

func initDB(dbConfig config.DBConfig) (*gorm.DB, error) {
	log.Logger.Debug(log.TraceMsgFuncStart(initDBMethod))
	defer log.Logger.Debug(log.TraceMsgFuncEnd(initDBMethod))

	var (
		dialector gorm.Dialector
		dsn       string
	)

	switch dbConfig.DBType {
	case config.PG:
		if dbConfig.ISCloudSQL {
			log.Logger.Debug(fmt.Sprintf(InitializingDBConn, dbConfig.ISCloudSQL, dbConfig.DBType))

			dsn = fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=%s",
				dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBSSLMode)
		} else {
			log.Logger.Debug(fmt.Sprintf(InitializingDBConn, dbConfig.ISCloudSQL, dbConfig.DBType))

			dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
				dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort, dbConfig.DBSSLMode)
		}

		dialector = postgres.Open(dsn)

	case config.MYSQL:
		if dbConfig.ISCloudSQL {
			log.Logger.Debug(fmt.Sprintf(InitializingDBConn, dbConfig.ISCloudSQL, dbConfig.DBType))

			dsn = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=%s",
				dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBName, dbConfig.DBSSLMode)
		} else {
			log.Logger.Debug(fmt.Sprintf(InitializingDBConn, dbConfig.ISCloudSQL, dbConfig.DBType))

			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s",
				dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName, dbConfig.DBSSLMode)
		}
		dialector = mysql.Open(dsn)
	}
	db, err := gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhen(OpenGormDBConnection), zap.Error(err))
		return nil, err
	}

	return db, nil
}

func initDefaultDB() error {
	log.Logger.Debug(log.TraceMsgFuncStart(initDefaultDBMethod))
	defer log.Logger.Debug(log.TraceMsgFuncEnd(initDefaultDBMethod))

	dbConfig := config.GetConfig().DBConfig

	db, err := initDB(dbConfig)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(initDBMethod), zap.Error(err))
		return err
	}

	SetDBConnection(db)
	return nil
}

// InitDBConWithAutoMigrate initializes the database connection and auto migrates the provided models
func InitDBConWithAutoMigrate(dst ...any) error {
	log.Logger.Debug(log.TraceMsgFuncStart(InitDBConWithAutoMigrateMethod))
	defer log.Logger.Debug(log.TraceMsgFuncEnd(InitDBConWithAutoMigrateMethod))

	err := InitDBConnection()
	if err != nil {
		log.Logger.Error(constant.DBConnectionOpenError, zap.Error(err))
		return err
	}

	err = dbCon.AutoMigrate(dst...)
	if err != nil {
		log.Logger.Error(constant.DBErrorOccurredWhenAutoMigrate, zap.Error(err))
		return err
	}

	return nil
}
