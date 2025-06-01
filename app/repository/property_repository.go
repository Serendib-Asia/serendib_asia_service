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

// createPropertyAmenities creates amenities for a property
func (r *propertyRepository) createPropertyAmenities(tx *gorm.DB, propertyID uint, amenityIDs []int) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	if len(amenityIDs) == 0 {
		return nil
	}

	// Remove duplicates from amenityIDs
	uniqueAmenityIDs := make(map[int]bool)
	var uniqueIDs []int
	for _, id := range amenityIDs {
		if !uniqueAmenityIDs[id] {
			uniqueAmenityIDs[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	amenities := make([]dto.PropertyAmenity, len(uniqueIDs))
	for i, amenityID := range uniqueIDs {
		amenities[i] = dto.PropertyAmenity{
			PropertyID: propertyID,
			AmenityID:  uint(amenityID),
		}
	}
	if err := tx.Model(&dto.PropertyAmenity{}).Create(&amenities).Error; err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
		return err
	}
	return nil
}

// createPropertyUtilities creates utilities for a property
func (r *propertyRepository) createPropertyUtilities(tx *gorm.DB, propertyID uint, utilityIDs []int) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	if len(utilityIDs) == 0 {
		return nil
	}

	utilities := make([]dto.PropertyUtility, len(utilityIDs))
	for i, utilityID := range utilityIDs {
		utilities[i] = dto.PropertyUtility{
			PropertyID: propertyID,
			UtilityID:  uint(utilityID),
		}
	}
	if err := tx.Model(&dto.PropertyUtility{}).Create(&utilities).Error; err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyUtility"), log.TraceError(commonLogFields, err)...)
		return err
	}
	return nil
}

// createPropertyImages creates images for a property
func (r *propertyRepository) createPropertyImages(tx *gorm.DB, propertyID uint, imageURLs []string) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	if len(imageURLs) == 0 {
		return nil
	}

	propertyImages := make([]dto.PropertyImage, len(imageURLs))
	for i, url := range imageURLs {
		propertyImages[i] = dto.PropertyImage{
			PropertyID: propertyID,
			URL:        url,
			IsPrimary:  i == 0, // Set the first image as primary
		}
	}
	if err := tx.Model(&dto.PropertyImage{}).Create(&propertyImages).Error; err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyImage"), log.TraceError(commonLogFields, err)...)
		return err
	}
	return nil
}

// mapRequestToProperty maps PropertyRequest to Property entity
func (r *propertyRepository) mapRequestToProperty(request dto.PropertyRequest) dto.Property {
	return dto.Property{
		UserID:          request.UserID,
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
}

// updatePropertyAmenities updates amenities for a property
func (r *propertyRepository) updatePropertyAmenities(tx *gorm.DB, propertyID uint, amenityIDs []int) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)

	// Remove duplicates from amenityIDs
	uniqueAmenityIDs := make(map[int]bool)
	var uniqueIDs []int
	for _, id := range amenityIDs {
		if !uniqueAmenityIDs[id] {
			uniqueAmenityIDs[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	// Get existing amenities including soft deleted ones
	var existingAmenities []dto.PropertyAmenity
	if err := tx.Unscoped().Where(&dto.PropertyAmenity{PropertyID: propertyID}).Find(&existingAmenities).Error; err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Create maps for easier lookup
	existingMap := make(map[uint]bool)
	softDeletedMap := make(map[uint]bool)
	for _, amenity := range existingAmenities {
		if amenity.DeletedAt.Valid {
			softDeletedMap[amenity.AmenityID] = true
		} else {
			existingMap[amenity.AmenityID] = true
		}
	}

	// Find amenities to add, reactivate, or deactivate
	var toAdd []dto.PropertyAmenity
	toReactivate := make(map[uint]bool)
	toDeactivate := make(map[uint]bool)

	// Check which amenities need to be added or reactivated
	for _, id := range uniqueIDs {
		if !existingMap[uint(id)] {
			if softDeletedMap[uint(id)] {
				// Reactivate soft deleted amenity
				toReactivate[uint(id)] = true
			} else {
				// Add new amenity
				toAdd = append(toAdd, dto.PropertyAmenity{
					PropertyID: propertyID,
					AmenityID:  uint(id),
				})
			}
		}
	}

	// Check which amenities need to be deactivated
	for _, existing := range existingAmenities {
		if !existing.DeletedAt.Valid { // Only check non-deleted amenities
			found := false
			for _, id := range uniqueIDs {
				if uint(id) == existing.AmenityID {
					found = true
					break
				}
			}
			if !found {
				toDeactivate[existing.AmenityID] = true
			}
		}
	}

	// First, deactivate amenities that are no longer needed
	if len(toDeactivate) > 0 {
		amenityIDsToDeactivate := make([]uint, 0, len(toDeactivate))
		for id := range toDeactivate {
			amenityIDsToDeactivate = append(amenityIDsToDeactivate, id)
		}
		if err := tx.Where(&dto.PropertyAmenity{PropertyID: propertyID}).Where("amenity_id IN ?", amenityIDsToDeactivate).Delete(&dto.PropertyAmenity{}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	// Then, reactivate soft deleted amenities
	if len(toReactivate) > 0 {
		amenityIDsToReactivate := make([]uint, 0, len(toReactivate))
		for id := range toReactivate {
			amenityIDsToReactivate = append(amenityIDsToReactivate, id)
		}
		if err := tx.Unscoped().Model(&dto.PropertyAmenity{}).
			Where(&dto.PropertyAmenity{PropertyID: propertyID}).
			Where("amenity_id IN ?", amenityIDsToReactivate).
			Update("deleted_at", nil).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	// Finally, add new amenities one by one to handle potential duplicates
	for _, amenity := range toAdd {
		// Check if the record already exists
		var count int64
		if err := tx.Model(&dto.PropertyAmenity{}).
			Where(&dto.PropertyAmenity{
				PropertyID: propertyID,
				AmenityID:  amenity.AmenityID,
			}).Count(&count).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Only insert if it doesn't exist
		if count == 0 {
			if err := tx.Create(&amenity).Error; err != nil {
				log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyAmenity"), log.TraceError(commonLogFields, err)...)
				return err
			}
		}
	}

	return nil
}

// updatePropertyUtilities updates utilities for a property
func (r *propertyRepository) updatePropertyUtilities(tx *gorm.DB, propertyID uint, utilityIDs []int) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)

	// Remove duplicates from utilityIDs
	uniqueUtilityIDs := make(map[int]bool)
	var uniqueIDs []int
	for _, id := range utilityIDs {
		if !uniqueUtilityIDs[id] {
			uniqueUtilityIDs[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	// Get existing utilities
	var existingUtilities []dto.PropertyUtility
	if err := tx.Where(&dto.PropertyUtility{PropertyID: propertyID}).Find(&existingUtilities).Error; err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PropertyUtility"), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Create maps for easier lookup
	existingMap := make(map[uint]bool)
	for _, utility := range existingUtilities {
		existingMap[utility.UtilityID] = true
	}

	// Find utilities to add and remove
	var toAdd []dto.PropertyUtility
	toRemove := make(map[uint]bool)

	// Check which utilities need to be added
	for _, id := range uniqueIDs {
		if !existingMap[uint(id)] {
			toAdd = append(toAdd, dto.PropertyUtility{
				PropertyID: propertyID,
				UtilityID:  uint(id),
			})
		}
	}

	// Check which utilities need to be removed
	for _, existing := range existingUtilities {
		found := false
		for _, id := range uniqueIDs {
			if uint(id) == existing.UtilityID {
				found = true
				break
			}
		}
		if !found {
			toRemove[existing.UtilityID] = true
		}
	}

	// Remove utilities that are no longer needed
	if len(toRemove) > 0 {
		utilityIDsToRemove := make([]uint, 0, len(toRemove))
		for id := range toRemove {
			utilityIDsToRemove = append(utilityIDsToRemove, id)
		}
		if err := tx.Where(&dto.PropertyUtility{PropertyID: propertyID}).Where("utility_id IN ?", utilityIDsToRemove).Delete(&dto.PropertyUtility{}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyUtility"), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	// Add new utilities
	if len(toAdd) > 0 {
		if err := tx.Create(&toAdd).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyUtility"), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	return nil
}

// updatePropertyImages updates images for a property
func (r *propertyRepository) updatePropertyImages(tx *gorm.DB, propertyID uint, imageURLs []string) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)

	// Get existing images
	var existingImages []dto.PropertyImage
	if err := tx.Where(&dto.PropertyImage{PropertyID: propertyID}).Find(&existingImages).Error; err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PropertyImage"), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Create maps for easier lookup
	existingMap := make(map[string]bool)
	for _, image := range existingImages {
		existingMap[image.URL] = true
	}

	// Find images to add and remove
	var toAdd []dto.PropertyImage
	toRemove := make(map[string]bool)

	// Check which images need to be added
	for i, url := range imageURLs {
		if !existingMap[url] {
			toAdd = append(toAdd, dto.PropertyImage{
				PropertyID: propertyID,
				URL:        url,
				IsPrimary:  i == 0, // Set the first image as primary
			})
		}
	}

	// Check which images need to be removed
	for _, existing := range existingImages {
		found := false
		for _, url := range imageURLs {
			if url == existing.URL {
				found = true
				break
			}
		}
		if !found {
			toRemove[existing.URL] = true
		}
	}

	// Remove images that are no longer needed
	if len(toRemove) > 0 {
		urlsToRemove := make([]string, 0, len(toRemove))
		for url := range toRemove {
			urlsToRemove = append(urlsToRemove, url)
		}
		if err := tx.Where(&dto.PropertyImage{PropertyID: propertyID}).Where("url IN ?", urlsToRemove).Delete(&dto.PropertyImage{}).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyImage"), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	// Add new images
	if len(toAdd) > 0 {
		if err := tx.Create(&toAdd).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyImage"), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	// Update primary image if needed
	if len(imageURLs) > 0 {
		// Reset all images to non-primary
		if err := tx.Model(&dto.PropertyImage{}).Where(&dto.PropertyImage{PropertyID: propertyID}).Update("is_primary", false).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("PropertyImage"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Set the first image as primary
		if err := tx.Model(&dto.PropertyImage{}).Where(&dto.PropertyImage{PropertyID: propertyID, URL: imageURLs[0]}).Update("is_primary", true).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("PropertyImage"), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	return nil
}

func (r *propertyRepository) Create(request dto.PropertyRequest) (uint, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(PropertyRepositoryCreateMethod), log.TraceMethodInputs(commonLogFields, request)...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(PropertyRepositoryCreateMethod), commonLogFields...)

	// Map request to property entity
	property := r.mapRequestToProperty(request)

	// Create property and its associations in a transaction
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Create the property
		if err := tx.Create(&property).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("Property"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Create property images
		if err := r.createPropertyImages(tx, property.ID, request.Images); err != nil {
			return err
		}

		// Create amenities
		if err := r.createPropertyAmenities(tx, property.ID, request.AmenityIDs); err != nil {
			return err
		}

		// Create utilities
		if err := r.createPropertyUtilities(tx, property.ID, request.UtilityIDs); err != nil {
			return err
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
		property := r.mapRequestToProperty(request)
		property.ID = id

		if err := tx.Model(&dto.Property{}).Where(dto.Property{ID: id}).Updates(property).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("Property"), log.TraceError(commonLogFields, err)...)
			return err
		}

		// Update amenities
		if err := r.updatePropertyAmenities(tx, id, request.AmenityIDs); err != nil {
			return err
		}

		// Update utilities
		if err := r.updatePropertyUtilities(tx, id, request.UtilityIDs); err != nil {
			return err
		}

		// Update images
		if err := r.updatePropertyImages(tx, id, request.Images); err != nil {
			return err
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
