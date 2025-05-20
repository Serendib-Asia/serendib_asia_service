package repository

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/config/dbconfig"
	"github.com/chazool/serendib_asia_service/pkg/log"

	"gorm.io/gorm"
)

const (
	// Lookup repository methods
	LookupRepositoryGetPurposeTypesMethod   = "LookupRepositoryGetPurposeTypes"
	LookupRepositoryGetPropertyTypesMethod  = "LookupRepositoryGetPropertyTypes"
	LookupRepositoryGetFurnitureTypesMethod = "LookupRepositoryGetFurnitureTypes"
	LookupRepositoryGetConditionsMethod     = "LookupRepositoryGetConditions"
	LookupRepositoryGetUtilitiesMethod      = "LookupRepositoryGetUtilities"
	LookupRepositoryGetAmenitiesMethod      = "LookupRepositoryGetAmenities"
)

type LookupRepository interface {
	GetPurposeTypes() ([]dto.LookupResponse, error)
	GetPropertyTypes() ([]dto.LookupResponse, error)
	GetFurnitureTypes() ([]dto.LookupResponse, error)
	GetConditions() ([]dto.LookupResponse, error)
	GetUtilities() ([]dto.LookupResponse, error)
	GetAmenities() ([]dto.LookupResponse, error)
}

type lookupRepository struct {
	_                 struct{}
	repositoryContext Context
	db                *gorm.DB
}

// CreateLookupRepository creates a new instance of LookupRepository
func CreateLookupRepository(requestID string) LookupRepository {
	return &lookupRepository{
		repositoryContext: CreateRepositoryContext(requestID),
		db:                dbconfig.GetDBConnection(),
	}
}

func (r *lookupRepository) GetPurposeTypes() ([]dto.LookupResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupRepositoryGetPurposeTypesMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(LookupRepositoryGetPurposeTypesMethod), commonLogFields...)

	var purposes []dto.LookupResponse
	err := r.db.Table("purpose_types").Find(&purposes).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PurposeTypes"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return purposes, nil
}

func (r *lookupRepository) GetPropertyTypes() ([]dto.LookupResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupRepositoryGetPropertyTypesMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(LookupRepositoryGetPropertyTypesMethod), commonLogFields...)

	var types []dto.LookupResponse
	err := r.db.Table("property_types").Find(&types).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PropertyTypes"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return types, nil
}

func (r *lookupRepository) GetFurnitureTypes() ([]dto.LookupResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupRepositoryGetFurnitureTypesMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(LookupRepositoryGetFurnitureTypesMethod), commonLogFields...)

	var types []dto.LookupResponse
	err := r.db.Table("furniture_types").Find(&types).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("FurnitureTypes"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return types, nil
}

func (r *lookupRepository) GetConditions() ([]dto.LookupResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupRepositoryGetConditionsMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(LookupRepositoryGetConditionsMethod), commonLogFields...)

	var conditions []dto.LookupResponse
	err := r.db.Table("property_conditions").Find(&conditions).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PropertyConditions"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return conditions, nil
}

func (r *lookupRepository) GetUtilities() ([]dto.LookupResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupRepositoryGetUtilitiesMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(LookupRepositoryGetUtilitiesMethod), commonLogFields...)

	var utilities []dto.LookupResponse
	err := r.db.Table("utilities").Find(&utilities).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("Utilities"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return utilities, nil
}

func (r *lookupRepository) GetAmenities() ([]dto.LookupResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(LookupRepositoryGetAmenitiesMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(LookupRepositoryGetAmenitiesMethod), commonLogFields...)

	var amenities []dto.LookupResponse
	err := r.db.Table("amenities").Find(&amenities).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("Amenities"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return amenities, nil
}
