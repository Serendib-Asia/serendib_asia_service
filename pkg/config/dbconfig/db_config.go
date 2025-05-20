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
var sdeskDBCon *gorm.DB

// GetDBConnection returns the current database connection
func GetDBConnection() *gorm.DB {
	return dbCon
}

// GetSDeskDBConnection returns the current sdesk database connection
func GetSDeskDBConnection() *gorm.DB {
	return sdeskDBCon
}

// SetDBConnection sets the current database connection
func SetDBConnection(db *gorm.DB) {
	dbCon = db
}

// SetSDeskDBConnection sets the current sdesk database connection
func SetSDeskDBConnection(sdeskDB *gorm.DB) {
	sdeskDBCon = sdeskDB
}

// InitDBConnection initializes the database connection
func InitDBConnection() (err error) {
	log.Logger.Debug(log.TraceMsgFuncStart(InitDBConnectionMethod))
	defer log.Logger.Debug(log.TraceMsgFuncEnd(InitDBConnectionMethod))

	// Initialize Default
	err = initDefaultDB()
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(initDefaultDBMethod), zap.Error(err))
		return fmt.Errorf("failed to connect to default database: %w", err)
	}

	switch config.GetConfig().DefaultClient {
	case constant.SDesk:
		// Initialize sDesk Databse
		err = initSDeskDB()
		if err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(initSDeskDBMethod), zap.Error(err))
			return fmt.Errorf("failed to connect to sDesk database: %w", err)
		}
	default:
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(InitDBConnectionMethod), zap.Any("unable to find default client", config.GetConfig().DefaultClient))
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

func initSDeskDB() error {
	log.Logger.Debug(log.TraceMsgFuncStart(initSDeskDBMethod))
	defer log.Logger.Debug(log.TraceMsgFuncEnd(initSDeskDBMethod))

	sdeskConfig := config.GetConfig().ClientDBConfig

	sdeskDB, err := initDB(sdeskConfig)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(initDBMethod), zap.Error(err))
		return err
	}

	SetSDeskDBConnection(sdeskDB)
	return nil
}

// InitDBConWithAutoMigrate initializes the database connection and auto migrates the provided models in ICX Database
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
