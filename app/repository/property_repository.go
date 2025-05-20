package repository

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/config/dbconfig"
	"github.com/chazool/serendib_asia_service/pkg/log"

	"gorm.io/gorm"
)

const (
	// Property repository methods
	PropertyRepositoryCreateMethod      = "PropertyRepositoryCreate"
	PropertyRepositoryGetByIDMethod     = "PropertyRepositoryGetByID"
	PropertyRepositoryUpdateMethod      = "PropertyRepositoryUpdate"
	PropertyRepositoryDeleteMethod      = "PropertyRepositoryDelete"
	PropertyRepositoryListMethod        = "PropertyRepositoryList"
	PropertyRepositoryCheckExistsMethod = "PropertyRepositoryCheckExists"
)

type PropertyRepository interface {
	Create(request dto.PropertyRequest) (uint, error)
	GetByID(id uint) (dto.PropertyResponse, error)
	Update(id uint, request dto.PropertyRequest) error
	Delete(id uint) error
	List(offset, limit int) ([]dto.PropertyResponse, error)
	CheckExists(id uint) (bool, error)
}

type propertyRepository struct {
	_                 struct{}
	repositoryContext Context
	db                *gorm.DB
}

// CreatePropertyRepository creates a new instance of PropertyRepository.
// It initializes the repository with the provided request ID and database connection.
// The function returns a PropertyRepository interface.
func CreatePropertyRepository(requestID string) PropertyRepository {
	return &propertyRepository{
		repositoryContext: CreateRepositoryContext(requestID),
		db:                dbconfig.GetDBConnection(),
	}
}

func (r *propertyRepository) Create(request dto.PropertyRequest) (uint, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyRepositoryCreateMethod), log.TraceMethodInputs(commonLogFields, request)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryCreateMethod), commonLogFields...)

	property := dto.Property{
		Title:           request.Title,
		Description:     request.Description,
		PurposeID:       uint(request.PurposeID),
		PropertyTypeID:  uint(request.PropertyTypeID),
		FurnitureTypeID: uint(request.FurnitureTypeID),
		ConditionID:     uint(request.ConditionID),
		Bedrooms:        request.Bedrooms,
		Bathrooms:       request.Bathrooms,
		Size:            request.Size,
		SizeUnit:        request.SizeUnit,
		City:            request.City,
		Address:         request.Address,
		PostalCode:      request.PostalCode,
		Latitude:        request.Latitude,
		Longitude:       request.Longitude,
		Price:           request.Price,
		PriceUnit:       request.PriceUnit,
		IsNegotiable:    request.IsNegotiable,
		RentalPeriod:    request.RentalPeriod,
		IsRefundable:    request.IsRefundable,
		PricingType:     request.PricingType,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&property).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("Property"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Insert amenities
		if len(request.AmenityIDs) > 0 {
			amenities := make([]dto.PropertyAmenity, len(request.AmenityIDs))
			for i, amenityID := range request.AmenityIDs {
				amenities[i] = dto.PropertyAmenity{
					PropertyID: property.ID,
					AmenityID:  uint(amenityID),
				}
			}
			if err := tx.Create(&amenities).Error; err != nil {
				log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
				return err
			}
		}

		// Insert utilities
		if len(request.UtilityIDs) > 0 {
			utilities := make([]dto.PropertyUtility, len(request.UtilityIDs))
			for i, utilityID := range request.UtilityIDs {
				utilities[i] = dto.PropertyUtility{
					PropertyID: property.ID,
					UtilityID:  uint(utilityID),
				}
			}
			if err := tx.Create(&utilities).Error; err != nil {
				log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyUtility"), log.TraceError(commonLogFields, err)...)
				return err
			}
		}

		return nil
	})

	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyRepositoryCreateMethod), logFields...)
		return 0, err
	}

	return property.ID, nil
}

func (r *propertyRepository) GetByID(id uint) (dto.PropertyResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyRepositoryGetByIDMethod), log.TraceMethodInputs(commonLogFields, id)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryGetByIDMethod), commonLogFields...)

	var property dto.Property
	err := r.db.Preload("PropertyAmenities").Preload("PropertyUtilities").Preload("PropertyImages").
		Where("id = ?", id).
		First(&property).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("Property"), log.TraceError(commonLogFields, err)...)
		return dto.PropertyResponse{}, err
	}

	amenities := make([]int, len(property.PropertyAmenities))
	for i, amenity := range property.PropertyAmenities {
		amenities[i] = int(amenity.AmenityID)
	}

	utilities := make([]int, len(property.PropertyUtilities))
	for i, utility := range property.PropertyUtilities {
		utilities[i] = int(utility.UtilityID)
	}

	images := make([]string, len(property.PropertyImages))
	for i, image := range property.PropertyImages {
		images[i] = image.URL
	}

	response := dto.PropertyResponse{
		ID:              property.ID,
		UserID:          property.UserID,
		Title:           property.Title,
		Description:     property.Description,
		PurposeID:       int(property.PurposeID),
		PropertyTypeID:  int(property.PropertyTypeID),
		FurnitureTypeID: int(property.FurnitureTypeID),
		ConditionID:     int(property.ConditionID),
		Bedrooms:        property.Bedrooms,
		Bathrooms:       property.Bathrooms,
		Size:            property.Size,
		SizeUnit:        property.SizeUnit,
		City:            property.City,
		Address:         property.Address,
		PostalCode:      property.PostalCode,
		Latitude:        property.Latitude,
		Longitude:       property.Longitude,
		Price:           property.Price,
		PriceUnit:       property.PriceUnit,
		IsNegotiable:    property.IsNegotiable,
		RentalPeriod:    property.RentalPeriod,
		IsRefundable:    property.IsRefundable,
		PricingType:     property.PricingType,
		CreatedAt:       property.CreatedAt,
		Amenities:       amenities,
		Utilities:       utilities,
		Images:          images,
	}

	log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryGetByIDMethod), log.TraceMethodOutputWithErr(commonLogFields, response, err)...)
	return response, nil
}

func (r *propertyRepository) Update(id uint, request dto.PropertyRequest) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyRepositoryUpdateMethod), log.TraceMethodInputs(commonLogFields, id, request)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryUpdateMethod), commonLogFields...)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Update property
		if err := tx.Model(&dto.Property{}).Where("id = ?", id).Updates(map[string]interface{}{
			"title":             request.Title,
			"description":       request.Description,
			"purpose_id":        uint(request.PurposeID),
			"property_type_id":  uint(request.PropertyTypeID),
			"furniture_type_id": uint(request.FurnitureTypeID),
			"condition_id":      uint(request.ConditionID),
			"bedrooms":          request.Bedrooms,
			"bathrooms":         request.Bathrooms,
			"size":              request.Size,
			"size_unit":         request.SizeUnit,
			"city":              request.City,
			"address":           request.Address,
			"postal_code":       request.PostalCode,
			"latitude":          request.Latitude,
			"longitude":         request.Longitude,
			"price":             request.Price,
			"price_unit":        request.PriceUnit,
			"is_negotiable":     request.IsNegotiable,
			"rental_period":     request.RentalPeriod,
			"is_refundable":     request.IsRefundable,
			"pricing_type":      request.PricingType,
		}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("Property"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Update amenities
		if err := tx.Where("property_id = ?", id).Delete(&dto.PropertyAmenity{}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
			return err
		}

		if len(request.AmenityIDs) > 0 {
			amenities := make([]dto.PropertyAmenity, len(request.AmenityIDs))
			for i, amenityID := range request.AmenityIDs {
				amenities[i] = dto.PropertyAmenity{
					PropertyID: id,
					AmenityID:  uint(amenityID),
				}
			}
			if err := tx.Create(&amenities).Error; err != nil {
				log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
				return err
			}
		}

		// Update utilities
		if err := tx.Where("property_id = ?", id).Delete(&dto.PropertyUtility{}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyUtility"), log.TraceError(commonLogFields, err)...)
			return err
		}

		if len(request.UtilityIDs) > 0 {
			utilities := make([]dto.PropertyUtility, len(request.UtilityIDs))
			for i, utilityID := range request.UtilityIDs {
				utilities[i] = dto.PropertyUtility{
					PropertyID: id,
					UtilityID:  uint(utilityID),
				}
			}
			if err := tx.Create(&utilities).Error; err != nil {
				log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyUtility"), log.TraceError(commonLogFields, err)...)
				return err
			}
		}

		return nil
	})

	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyRepositoryUpdateMethod), logFields...)
		return err
	}

	return nil
}

func (r *propertyRepository) Delete(id uint) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyRepositoryDeleteMethod), log.TraceMethodInputs(commonLogFields, id)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryDeleteMethod), commonLogFields...)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Delete amenities
		if err := tx.Where("property_id = ?", id).Delete(&dto.PropertyAmenity{}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Delete utilities
		if err := tx.Where("property_id = ?", id).Delete(&dto.PropertyUtility{}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyUtility"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Delete images
		if err := tx.Where("property_id = ?", id).Delete(&dto.PropertyImage{}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyImage"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Delete property
		if err := tx.Delete(&dto.Property{}, "id = ?", id).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("Property"), log.TraceError(commonLogFields, err)...)
			return err
		}

		return nil
	})

	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(PropertyRepositoryDeleteMethod), logFields...)
		return err
	}

	return nil
}

func (r *propertyRepository) List(offset, limit int) ([]dto.PropertyResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyRepositoryListMethod), log.TraceMethodInputs(commonLogFields, offset, limit)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryListMethod), commonLogFields...)

	var properties []dto.Property
	err := r.db.Preload("PropertyAmenities").Preload("PropertyUtilities").Preload("PropertyImages").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&properties).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("Property"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	responses := make([]dto.PropertyResponse, len(properties))
	for i, property := range properties {
		amenities := make([]int, len(property.PropertyAmenities))
		for j, amenity := range property.PropertyAmenities {
			amenities[j] = int(amenity.AmenityID)
		}

		utilities := make([]int, len(property.PropertyUtilities))
		for j, utility := range property.PropertyUtilities {
			utilities[j] = int(utility.UtilityID)
		}

		images := make([]string, len(property.PropertyImages))
		for j, image := range property.PropertyImages {
			images[j] = image.URL
		}

		responses[i] = dto.PropertyResponse{
			ID:              property.ID,
			UserID:          property.UserID,
			Title:           property.Title,
			Description:     property.Description,
			PurposeID:       int(property.PurposeID),
			PropertyTypeID:  int(property.PropertyTypeID),
			FurnitureTypeID: int(property.FurnitureTypeID),
			ConditionID:     int(property.ConditionID),
			Bedrooms:        property.Bedrooms,
			Bathrooms:       property.Bathrooms,
			Size:            property.Size,
			SizeUnit:        property.SizeUnit,
			City:            property.City,
			Address:         property.Address,
			PostalCode:      property.PostalCode,
			Latitude:        property.Latitude,
			Longitude:       property.Longitude,
			Price:           property.Price,
			PriceUnit:       property.PriceUnit,
			IsNegotiable:    property.IsNegotiable,
			RentalPeriod:    property.RentalPeriod,
			IsRefundable:    property.IsRefundable,
			PricingType:     property.PricingType,
			CreatedAt:       property.CreatedAt,
			Amenities:       amenities,
			Utilities:       utilities,
			Images:          images,
		}
	}

	log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryListMethod), log.TraceMethodOutputWithErr(commonLogFields, responses, err)...)
	return responses, nil
}

func (r *propertyRepository) CheckExists(id uint) (bool, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyRepositoryCheckExistsMethod), log.TraceMethodInputs(commonLogFields, id)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryCheckExistsMethod), commonLogFields...)

	var count int64
	err := r.db.Model(&dto.Property{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenCounting("Property"), log.TraceError(commonLogFields, err)...)
		return false, err
	}

	return count > 0, nil
}
