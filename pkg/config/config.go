package config

import (
	"encoding/json"
	"log"
	"slices"
	"sort"
	"time"

	"github.com/chazool/serendib_asia_service/pkg/config/firebase"
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
	SrvListenPort                = "SRV_LISTEN_PORT"
	ChildFiberProcessIdleTimeout = "CHILD_FIBER_PROCESS_IDLE_TIMEOUT"
	LogDestination               = "LOG_DESTINATION"
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
	FirebaseConfig               firebase.Config
	ChildFiberProcessIdleTimeout time.Duration
	SrvListenPort                string
	Pprofenabled                 bool
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

	// Set Firebase default config
	firebase.SetDefaultConfig()

	// you can supply "console" or "File". if json, logging formant is in json
	viper.SetDefault(LogFileName, JSON)
	viper.SetDefault(Pprofenabled, "true")
}

func (config *CommonConfig) setDefaultDBConfig() {
	viper.SetDefault(DBHost, "localhost")
	viper.SetDefault(DBPort, "5432")
	viper.SetDefault(DBName, "serendib")
	viper.SetDefault(DBUser, "postgres")
	viper.SetDefault(DBPassword, "postgres")
	viper.SetDefault(DBSSLMode, "disable")
	viper.SetDefault(ISCloudSQL, false)
}

// BuildConfig is using to build the application configuration
func (config *CommonConfig) BuildConfig() *CommonConfig {
	config.setDefaultConfig()

	// Call the SetDBDefaultConfig function
	config.setDefaultDBConfig()

	viper.AutomaticEnv()
	logConfig, logger := config.getLogConfig()

	config = &CommonConfig{
		LogConfig:                    logConfig,
		DBConfig:                     config.getDBConfig(),
		FirebaseConfig:               firebase.GetConfig(),
		ChildFiberProcessIdleTimeout: viper.GetDuration(ChildFiberProcessIdleTimeout),
		SrvListenPort:                viper.GetString(SrvListenPort),
		Pprofenabled:                 viper.GetBool(Pprofenabled),
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
