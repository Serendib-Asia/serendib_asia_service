package main

import (
	"github.com/chazool/serendib_asia_service/app/routes"
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/app/routes/handler/validator"

	"github.com/chazool/serendib_asia_service/pkg/config"
	"github.com/chazool/serendib_asia_service/pkg/config/appconfig"
	"github.com/chazool/serendib_asia_service/pkg/config/dbconfig"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"go.uber.org/zap"
)

func init() {
	config.InitConfig()

	err := dbconfig.InitDBConWithAutoMigrate(&dto.Property{})
	if err != nil {
		log.Logger.Error(constant.DBInitFailError, zap.Error(err))
	}

	utils.HTTPClientImplInstance = utils.NewHTTPClientUtil()
	validator.InitValidator()
}

// @SecurityDefinitions.api apiKey
// @Scheme bearer
// @In header
// @Name Authorization
// @title Serendib Asia Service
// @version V1.0
// @description This is Serendib Asia Service API's
// @contact.name Serendib Asia Support
// @contact.email support@serendib.asia
// @BasePath /api/v1
// @schemes http https
// Serendib Asia Service
// This is the main entry point for the Serendib Asia Service
// It initializes the configuration, database connection, and starts the API routes
func main() {
	appconfig.Start(routes.APIRoutes)
}
