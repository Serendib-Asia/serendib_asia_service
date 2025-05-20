package services

import (
	"runtime/debug"

	"github.com/chazool/serendib_asia_service/app/repository"
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"

	"gorm.io/gorm"
)

const (
	// Lookup service methods
	LookupServiceGetPurposeTypesMethod   = "LookupServiceGetPurposeTypes"
	LookupServiceGetPropertyTypesMethod  = "LookupServiceGetPropertyTypes"
	LookupServiceGetFurnitureTypesMethod = "LookupServiceGetFurnitureTypes"
	LookupServiceGetConditionsMethod     = "LookupServiceGetConditions"
	LookupServiceGetUtilitiesMethod      = "LookupServiceGetUtilities"
	LookupServiceGetAmenitiesMethod      = "LookupServiceGetAmenities"
)

// LookupService defines the interface for lookup service methods
type LookupService struct {
	_              struct{}
	serviceContext ServiceContext
	transaction    *gorm.DB
	lookupRepo     repository.LookupRepository
}

// CreateLookupService creates a new instance of LookupService
func CreateLookupService(requestID string, transactionDB *gorm.DB) *LookupService {
	return &LookupService{
		serviceContext: CreateServiceContext(requestID),
		transaction:    transactionDB,
	}
}

// GetPurposeTypes retrieves all purpose types
func (service *LookupService) GetPurposeTypes() (response []dto.LookupResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupServiceGetPurposeTypesMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(LookupServiceGetPurposeTypesMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(LookupServiceGetPurposeTypesMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.lookupRepo = repository.CreateLookupRepository(service.serviceContext.RequestID)
	response, err := service.lookupRepo.GetPurposeTypes()
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.LookupRepositoryGetPurposeTypesMethod), logFields...)
		return nil, buildSelectErrFromRepo("purpose types", err)
	}

	return response, nil
}

// GetPropertyTypes retrieves all property types
func (service *LookupService) GetPropertyTypes() (response []dto.LookupResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupServiceGetPropertyTypesMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(LookupServiceGetPropertyTypesMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(LookupServiceGetPropertyTypesMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.lookupRepo = repository.CreateLookupRepository(service.serviceContext.RequestID)
	response, err := service.lookupRepo.GetPropertyTypes()
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.LookupRepositoryGetPropertyTypesMethod), logFields...)
		return nil, buildSelectErrFromRepo("property types", err)
	}

	return response, nil
}

// GetFurnitureTypes retrieves all furniture types
func (service *LookupService) GetFurnitureTypes() (response []dto.LookupResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupServiceGetFurnitureTypesMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(LookupServiceGetFurnitureTypesMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(LookupServiceGetFurnitureTypesMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.lookupRepo = repository.CreateLookupRepository(service.serviceContext.RequestID)
	response, err := service.lookupRepo.GetFurnitureTypes()
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.LookupRepositoryGetFurnitureTypesMethod), logFields...)
		return nil, buildSelectErrFromRepo("furniture types", err)
	}

	return response, nil
}

// GetConditions retrieves all property conditions
func (service *LookupService) GetConditions() (response []dto.LookupResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupServiceGetConditionsMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(LookupServiceGetConditionsMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(LookupServiceGetConditionsMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.lookupRepo = repository.CreateLookupRepository(service.serviceContext.RequestID)
	response, err := service.lookupRepo.GetConditions()
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.LookupRepositoryGetConditionsMethod), logFields...)
		return nil, buildSelectErrFromRepo("property conditions", err)
	}

	return response, nil
}

// GetUtilities retrieves all utilities
func (service *LookupService) GetUtilities() (response []dto.LookupResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupServiceGetUtilitiesMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(LookupServiceGetUtilitiesMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(LookupServiceGetUtilitiesMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.lookupRepo = repository.CreateLookupRepository(service.serviceContext.RequestID)
	response, err := service.lookupRepo.GetUtilities()
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.LookupRepositoryGetUtilitiesMethod), logFields...)
		return nil, buildSelectErrFromRepo("utilities", err)
	}

	return response, nil
}

// GetAmenities retrieves all amenities
func (service *LookupService) GetAmenities() (response []dto.LookupResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupServiceGetAmenitiesMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(LookupServiceGetAmenitiesMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(LookupServiceGetAmenitiesMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.lookupRepo = repository.CreateLookupRepository(service.serviceContext.RequestID)
	response, err := service.lookupRepo.GetAmenities()
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.LookupRepositoryGetAmenitiesMethod), logFields...)
		return nil, buildSelectErrFromRepo("amenities", err)
	}

	return response, nil
}
