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
// @Name Dashboard
// @title ICX Dashboard Service
// @version V1.0
// @description This is ICX Dashboard Service API's
// @contact.name ICX Support
// @contact.email support@ivedha.com
// @BasePath /api/v1
// @schemes http https
func main() {
	appconfig.Start(routes.APIRoutes)
}
