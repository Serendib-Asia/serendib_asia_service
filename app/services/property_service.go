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
	// Property service methods
	PropertyServiceCreateMethod  = "PropertyServiceCreate"
	PropertyServiceGetByIDMethod = "PropertyServiceGetByID"
	PropertyServiceUpdateMethod  = "PropertyServiceUpdate"
	PropertyServiceDeleteMethod  = "PropertyServiceDelete"
	PropertyServiceListMethod    = "PropertyServiceList"
)

// PropertyService defines the interface for property service methods.
type PropertyService struct {
	_              struct{}
	serviceContext ServiceContext
	transaction    *gorm.DB
	propertyRepo   repository.PropertyRepository
	userRepo       repository.UserRepository
}

// CreatePropertyService creates a new instance of PropertyService.
// It initializes the service context with the provided request ID and transaction database.
// The function returns a PropertyService interface.
func CreatePropertyService(requestID string, transactionDB *gorm.DB) *PropertyService {
	return &PropertyService{
		serviceContext: CreateServiceContext(requestID),
		transaction:    transactionDB,
	}
}

// Create creates a new property
func (service *PropertyService) Create(request dto.PropertyRequest) (response *dto.Property, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyServiceCreateMethod), log.TraceMethodInputs(commonLogFields, request)...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(PropertyServiceCreateMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(PropertyServiceCreateMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.propertyRepo = repository.CreatePropertyRepository(service.serviceContext.RequestID)

	// Validate images count
	if len(request.Images) == 0 {
		errRes := custom.BuildBadReqErrResult(constant.ErrCodeInvalidInput, "at least one property image is required", constant.Empty)
		return nil, &errRes
	}
	if len(request.Images) > 6 {
		errRes := custom.BuildBadReqErrResult(constant.ErrCodeInvalidInput, "maximum of 6 property images allowed", constant.Empty)
		return nil, &errRes
	}

	// Create property
	propertyID, err := service.propertyRepo.Create(request)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.PropertyRepositoryCreateMethod), logFields...)
		return nil, buildSelectErrFromRepo("property", err)
	}

	property, err := service.propertyRepo.GetByID(propertyID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.PropertyRepositoryGetByIDMethod), logFields...)
		return nil, buildSelectErrFromRepo("property", err)
	}

	return &property, nil
}

// GetByID retrieves a property by ID
func (service *PropertyService) GetByID(propertyID uint) (response dto.Property, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyServiceGetByIDMethod), log.TraceMethodInputs(commonLogFields, propertyID)...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(PropertyServiceGetByIDMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(PropertyServiceGetByIDMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.propertyRepo = repository.CreatePropertyRepository(service.serviceContext.RequestID)
	property, err := service.propertyRepo.GetByID(propertyID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.PropertyRepositoryGetByIDMethod), logFields...)
		return response, buildSelectErrFromRepo("property", err)
	}

	return property, nil
}

// Update updates a property
func (service *PropertyService) Update(propertyID uint, request dto.PropertyRequest) (response dto.Property, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyServiceUpdateMethod), log.TraceMethodInputs(commonLogFields, propertyID, request)...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(PropertyServiceUpdateMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(PropertyServiceUpdateMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.propertyRepo = repository.CreatePropertyRepository(service.serviceContext.RequestID)
	err := service.propertyRepo.Update(propertyID, request)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.PropertyRepositoryUpdateMethod), logFields...)
		return response, buildSelectErrFromRepo("property", err)
	}

	property, err := service.propertyRepo.GetByID(propertyID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.PropertyRepositoryGetByIDMethod), logFields...)
		return response, buildSelectErrFromRepo("property", err)
	}

	return property, nil
}

// Delete deletes a property
func (service *PropertyService) Delete(propertyID uint) (response dto.Property, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyServiceDeleteMethod), log.TraceMethodInputs(commonLogFields, propertyID)...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(PropertyServiceDeleteMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(PropertyServiceDeleteMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.propertyRepo = repository.CreatePropertyRepository(service.serviceContext.RequestID)
	property, err := service.propertyRepo.GetByID(propertyID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.PropertyRepositoryGetByIDMethod), logFields...)
		return response, buildSelectErrFromRepo("property", err)
	}

	err = service.propertyRepo.Delete(propertyID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.PropertyRepositoryDeleteMethod), logFields...)
		return response, buildSelectErrFromRepo("property", err)
	}

	return property, nil
}

// List lists properties with pagination
func (service *PropertyService) List(offset, limit int) (response []dto.Property, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyServiceListMethod), log.TraceMethodInputs(commonLogFields, offset, limit)...)

	defer func() {
		// Panic handling
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(PropertyServiceListMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(PropertyServiceListMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.propertyRepo = repository.CreatePropertyRepository(service.serviceContext.RequestID)
	properties, err := service.propertyRepo.List(offset, limit)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.PropertyRepositoryListMethod), logFields...)
		return nil, buildSelectErrFromRepo("properties", err)
	}

	return properties, nil
}
