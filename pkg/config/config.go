package config

import (
	"encoding/json"
	"log"
	"slices"
	"sort"
	"time"

	lg "github.com/chazool/serendib_asia_service/pkg/log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	appConfig *CommonConfig
)

// configuration constance
const (
	SrvListenPort                     = "SRV_LISTEN_PORT"
	ChildFiberProcessIdleTimeout      = "CHILD_FIBER_PROCESS_IDLE_TIMEOUT"
	LogDestination                    = "LOG_DESTINATION"
	AllowedPriorities                 = "ALLOWED_PRIORITIES"
	AuthServerBaseURL                 = "AUTH_SERVER_BASE_URL"
	AiMlServiceBaseURL                = "AI_ML_SERVICE_BASE_URL"
	AuthServerAllUsersDetailsEndpoint = "SEC_SRV_GET_ALL_USER_DETAILS_ENDPOINT"
	AuthServerTokenValidateEndpoint   = "TOKEN_VALIDATE_ENDPOINT"
	AuthServerUserDetailsEndpoint     = "USER_DETAILS_ENDPOINT"
	AiMlServiceCodeGenEndpoint        = "CODE_GENERATION_ENDPOINT"
	ICX_APP                           = "TENANT_NAME"
	SDeskTicketBaseURL                = "SDESK_BASE_URL"
	DefaultTicketBaseURL              = "DEFAULT_BASE_URL"
	DefaultClient                     = "DEFAULT_CLIENT"
	SDeskClient                       = "SDESK"
	// log constance
	LogFileName                 = "LOG_FILE_NAME"
	LogMaxSizeMb                = "LOG_MAX_SIZE_MB"
	LogMaxBackupDays            = "LOG_MAX_BACKUP_DAYS"
	LogMaxAgeDaysBeforeRollover = "LOG_MAX_AGE_DAYS_BEFORE_ROLLOVER"
	LogCompressionEnabled       = "LOG_COMPOSITION_ENABLED"
	LogLevel                    = "LOG_LEVEL"
	LogFormat                   = "LOG_FORMAT"
	Pprofenabled                = "PPROF_ENABLED"
	// db constance
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBName     = "DB_NAME"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBSSLMode  = "DB_SSLMODE"
	ISCloudSQL = "IS_CLOUD_SQL"

	// client db
	ClientDBHost     = "CLIENT_DB_HOST"
	ClientDBPort     = "CLIENT_DB_PORT"
	ClientDBName     = "CLIENT_DB_NAME"
	ClientDBUser     = "CLIENT_DB_USER"
	ClientDBPassword = "CLIENT_DB_PASSWORD"
	ClientDBSSLMode  = "CLIENT_DB_SSLMODE"

	// log constance values
	Console = "console"
	File    = "file"
	Debug   = "DEBUG"
	JSON    = "json"
)

// configuration values
const (
	// log values
	LogFileNameValue                 = "app.log"
	LogMaxSizeMbValue                = 100
	LogMaxBackupDaysValue            = 30
	LogMaxAgeDaysBeforeRolloverValue = 28
	LogCompressionEnabledValue       = true
	// priorities
	DefaultAllowedPriorities = "1 2 3"
)

// CommonConfig is a struct that holds the common configuration for the application
type CommonConfig struct {
	_ struct{}
	LogConfig
	DBConfig
	ClientDBConfig                    DBConfig
	SrvListenPort                     string
	ChildFiberProcessIdleTimeout      time.Duration
	Pprofenabled                      bool
	AllowedPriorities                 []string
	AuthServerBaseURL                 string
	AiMlServiceBaseURL                string
	AuthServerAllUsersDetailsEndpoint string
	AuthServerTokenValidateEndpoint   string
	AuthServerUserDetailsEndpoint     string
	AiMlServiceCodeGenEndpoint        string
	ICX_APP                           string
	SDeskTicketBaseURL                string
	DefaultTicketBaseURL              string
	DefaultClient                     string
	SDeskClient                       string
}

// LogConfig is a struct that holds the log configuration for the application
type LogConfig struct {
	_                           struct{}
	LogDestination              string
	LogFileName                 string
	LogMaxSizeMb                int
	LogMaxBackupDays            int
	LogMaxAgeDaysBeforeRollover int
	LogCompression              bool
	LogLevel                    string
	LogFormat                   string
	AppEnvironment              string
}

type DBType string

const (
	PG    DBType = "pg"
	MYSQL DBType = "mysql"
)

// DBConfig is a struct that holds the database configuration for the application
type DBConfig struct {
	_          struct{}
	DBType     DBType
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
	ISCloudSQL bool
}

// setDefaultConfig is using added application default configurations
func (config *CommonConfig) setDefaultConfig() {
	viper.SetDefault(SrvListenPort, "8080")

	dur, _ := time.ParseDuration("10s")
	viper.SetDefault(ChildFiberProcessIdleTimeout, dur)

	// log default config
	// you can supply "console" or "File". if console, logging goes to stdout, if tile, goes to LOG_FILE_NAME
	viper.SetDefault(LogDestination, Console)
	viper.SetDefault(LogFileName, LogFileNameValue)
	viper.SetDefault(LogMaxSizeMb, LogMaxSizeMbValue)
	viper.SetDefault(LogMaxBackupDays, LogMaxBackupDaysValue)
	viper.SetDefault(LogMaxAgeDaysBeforeRollover, LogMaxAgeDaysBeforeRolloverValue)
	viper.SetDefault(LogCompressionEnabled, LogCompressionEnabledValue)
	viper.SetDefault(LogLevel, Debug)
	viper.SetDefault(LogFormat, Console)

	// you can supply "console" or "File". if json, logging formant is in json
	viper.SetDefault(LogFileName, JSON)
	viper.SetDefault(Pprofenabled, "true")
	viper.SetDefault(AllowedPriorities, DefaultAllowedPriorities)
	viper.SetDefault(AuthServerBaseURL, "http://localhost:8081")
	viper.SetDefault(AiMlServiceBaseURL, "https://icx-ai-ml-non-prod-fhhapl7sza-uc.a.run.app")
	viper.SetDefault(AuthServerAllUsersDetailsEndpoint, "/icx/auth/v1/users")
	viper.SetDefault(AuthServerTokenValidateEndpoint, "/icx/auth/v1/validate/token/icx-dashboard")
	viper.SetDefault(AuthServerUserDetailsEndpoint, "/icx/auth/v1/user")
	viper.SetDefault(AiMlServiceCodeGenEndpoint, "/code-generate")
	viper.SetDefault(ICX_APP, "icx-dashboard")

	viper.SetDefault(SDeskTicketBaseURL, "https://testinstance.uat.sdesk.co.uk/")
	viper.SetDefault(DefaultTicketBaseURL, "https://testinstance.uat.sdesk.co.uk/")
	viper.SetDefault(DefaultClient, "")
	viper.SetDefault(SDeskClient, "sdesk")
}

func (config *CommonConfig) setDefaultDBConfig() {
	viper.SetDefault(DBHost, "localhost")
	viper.SetDefault(DBPort, "5432")
	viper.SetDefault(DBName, "icx_dashboard_db")
	viper.SetDefault(DBUser, "postgres")
	viper.SetDefault(DBPassword, "postgres")
	viper.SetDefault(DBSSLMode, "disable")
	viper.SetDefault(ISCloudSQL, false)
}

func (config *CommonConfig) setClientDBConfig() {
	viper.SetDefault(ClientDBHost, "34.141.37.153")
	viper.SetDefault(ClientDBPort, "3306")
	viper.SetDefault(ClientDBName, "sdesk")
	viper.SetDefault(ClientDBUser, "sdesk_db")
	viper.SetDefault(ClientDBPassword, "P00074950d")
	viper.SetDefault(ClientDBSSLMode, true)
}

// BuildConfig is using to build the application configuration
func (config *CommonConfig) BuildConfig() *CommonConfig {
	config.setDefaultConfig()

	// Call the SetDBDefaultConfig function
	config.setDefaultDBConfig()
	config.setClientDBConfig()

	viper.AutomaticEnv()
	logConfig, logger := config.getLogConfig()

	config = &CommonConfig{
		LogConfig:                         logConfig,
		DBConfig:                          config.getDBConfig(),
		ClientDBConfig:                    config.getClientDBConfig(),
		ChildFiberProcessIdleTimeout:      viper.GetDuration(ChildFiberProcessIdleTimeout),
		SrvListenPort:                     viper.GetString(SrvListenPort),
		Pprofenabled:                      viper.GetBool(Pprofenabled),
		AllowedPriorities:                 viper.GetStringSlice(AllowedPriorities),
		AuthServerBaseURL:                 viper.GetString(AuthServerBaseURL),
		AiMlServiceBaseURL:                viper.GetString(AiMlServiceBaseURL),
		AuthServerAllUsersDetailsEndpoint: viper.GetString(AuthServerAllUsersDetailsEndpoint),
		AuthServerTokenValidateEndpoint:   viper.GetString(AuthServerTokenValidateEndpoint),
		AuthServerUserDetailsEndpoint:     viper.GetString(AuthServerUserDetailsEndpoint),
		AiMlServiceCodeGenEndpoint:        viper.GetString(AiMlServiceCodeGenEndpoint),
		ICX_APP:                           viper.GetString(ICX_APP),
		SDeskTicketBaseURL:                viper.GetString(SDeskTicketBaseURL),
		DefaultTicketBaseURL:              viper.GetString(DefaultTicketBaseURL),
		DefaultClient:                     viper.GetString(DefaultClient),
		SDeskClient:                       viper.GetString(SDeskClient),
	}

	configJSONPresntation, _ := json.Marshal(config)
	logger.Info("Settup Config", zap.String("AppConfig", string(configJSONPresntation)))

	appConfig = config
	config.buildLogger()

	return config
}

func (config *CommonConfig) getDBConfig() DBConfig {
	return DBConfig{
		DBType:     PG,
		DBHost:     viper.GetString(DBHost),
		DBPort:     viper.GetString(DBPort),
		DBName:     viper.GetString(DBName),
		DBUser:     viper.GetString(DBUser),
		DBPassword: viper.GetString(DBPassword),
		DBSSLMode:  viper.GetString(DBSSLMode),
		ISCloudSQL: viper.GetBool(ISCloudSQL),
	}
}

func (config *CommonConfig) getClientDBConfig() DBConfig {
	return DBConfig{
		DBType:     MYSQL,
		DBHost:     viper.GetString(ClientDBHost),
		DBPort:     viper.GetString(ClientDBPort),
		DBName:     viper.GetString(ClientDBName),
		DBUser:     viper.GetString(ClientDBUser),
		DBPassword: viper.GetString(ClientDBPassword),
		DBSSLMode:  viper.GetString(ClientDBSSLMode),
		ISCloudSQL: viper.GetBool(ISCloudSQL),
	}
}

// getLogConfig is using set up the zap logger configuration
func (config *CommonConfig) getLogConfig() (LogConfig, *zap.Logger) {
	configLogger, err := zap.NewDevelopmentConfig().Build()
	if err != nil {
		log.Fatal("Failed to create logger", err)
	}

	defer configLogger.Sync()

	logDestination := viper.GetString(LogDestination)
	supportedLogDests := []string{Console, File} // supported log destinations
	sort.Strings(supportedLogDests)

	if !slices.Contains(supportedLogDests, logDestination) {
		configLogger.Fatal("Invalid log destination specified", zap.String(LogDestination, logDestination), zap.Strings("supportedLogDestinations", supportedLogDests))
	} else {
		configLogger.Info("log destination is set to ", zap.String(LogDestination, logDestination))
	}

	supportedLogFormats := []string{Console, JSON}
	logFormat := viper.GetString(LogFormat)

	if !slices.Contains(supportedLogFormats, logFormat) {
		configLogger.Fatal("Invalid Log Format specified", zap.String(LogFormat, logFormat), zap.Strings("supportedLogFormats", supportedLogFormats))
	} else {
		configLogger.Info("log format is se to", zap.String(LogFormat, logFormat))
	}

	logConfig := LogConfig{
		LogDestination:              logDestination,
		LogFormat:                   logFormat,
		LogFileName:                 viper.GetString(LogFileName),
		LogLevel:                    viper.GetString(LogLevel),
		LogMaxSizeMb:                viper.GetInt(LogMaxSizeMb),
		LogMaxBackupDays:            viper.GetInt(LogMaxBackupDays),
		LogMaxAgeDaysBeforeRollover: viper.GetInt(LogMaxAgeDaysBeforeRollover),
		LogCompression:              viper.GetBool(LogCompressionEnabled),
		AppEnvironment:              viper.GetString("APP_ENVIRONMENT"),
	}

	return logConfig, configLogger
}

func (config CommonConfig) buildLogger() {
	ZapLogLevel := map[string]zapcore.Level{
		"DEBUG":  zapcore.DebugLevel,
		"INFO":   zapcore.InfoLevel,
		"WARN":   zapcore.WarnLevel,
		"ERROR":  zapcore.ErrorLevel,
		"FATAL":  zapcore.FatalLevel,
		"PANIC":  zapcore.PanicLevel,
		"DPANIC": zapcore.DPanicLevel,
	}

	var (
		logLevel        = ZapLogLevel[config.LogLevel]
		err       error = nil
		core      zapcore.Core
		zapLogger *zap.Logger
	)

	if logLevel == 0 {
		log.Fatalf("can't initialize zap logger - unsupported log level %v", logLevel)
	}

	if config.LogDestination == File {
		LogConfig := zap.NewDevelopmentEncoderConfig()
		LogConfig.FunctionKey = "F"

		var enc zapcore.Encoder

		if config.LogFormat == JSON {
			enc = zapcore.NewJSONEncoder(LogConfig)
		} else {
			enc = zapcore.NewConsoleEncoder(LogConfig)
		}

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.LogFileName,
			MaxSize:    config.LogMaxSizeMb,
			MaxBackups: config.LogMaxBackupDays,
			MaxAge:     config.LogMaxAgeDaysBeforeRollover,
			Compress:   config.LogCompression,
		})

		core = zapcore.NewCore(enc, w, logLevel)
		zapLogger = zap.New(core)
	} else {
		logConfig := zap.NewDevelopmentConfig()
		logConfig.Level = zap.NewAtomicLevelAt(logLevel)
		logConfig.Encoding = config.LogFormat
		zapLogger, err = logConfig.Build()
	}

	if err != nil {
		log.Fatalf("can't initialize zap logger %v", err)
	}

	lg.Logger = zapLogger
}

// GetConfig returns the application configuration
func GetConfig() *CommonConfig {
	return appConfig
}

// InitConfig initializes the application configuration
func InitConfig() {
	config := &CommonConfig{}
	appConfig = config.BuildConfig()
}
