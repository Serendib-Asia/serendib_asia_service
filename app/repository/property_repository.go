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

// updateRelatedEntities is a generic function to update related entities with soft delete support
func updateRelatedEntities[T any, K comparable](
	tx *gorm.DB,
	propertyID uint,
	newIDs []K,
	getExisting func(tx *gorm.DB, propertyID uint) ([]T, error),
	compareKey func(T) K,
	isDeleted func(T) bool,
	createNew func(K) T,
	deleteWhere func(*gorm.DB, []K) *gorm.DB,
	entityName string,
	requestID string,
	extraOps ...func(tx *gorm.DB, newIDs []K) error,
) error {
	commonLogFields := log.CommonLogField(requestID)

	// Remove duplicates from newIDs
	uniqueIDsMap := make(map[K]bool)
	var uniqueIDs []K
	for _, id := range newIDs {
		if !uniqueIDsMap[id] {
			uniqueIDsMap[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	// Get existing entities including soft-deleted
	existing, err := getExisting(tx, propertyID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting(entityName), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Create maps for existing and soft-deleted entities
	existingMap := make(map[K]bool)
	softDeletedMap := make(map[K]bool)
	for _, entity := range existing {
		key := compareKey(entity)
		if isDeleted(entity) {
			softDeletedMap[key] = true
		} else {
			existingMap[key] = true
		}
	}

	// Determine entities to add, reactivate, or deactivate
	var toAdd []T
	var toReactivate []K
	for _, id := range uniqueIDs {
		if !existingMap[id] {
			if softDeletedMap[id] {
				// Mark for reactivation
				toReactivate = append(toReactivate, id)
			} else {
				// Add new entity
				toAdd = append(toAdd, createNew(id))
			}
		}
	}

	// Determine entities to deactivate
	var toDeactivate []K
	for _, entity := range existing {
		key := compareKey(entity)
		if !isDeleted(entity) { // Only check non-deleted entities
			found := false
			for _, id := range uniqueIDs {
				if id == key {
					found = true
					break
				}
			}
			if !found {
				toDeactivate = append(toDeactivate, key)
			}
		}
	}

	// Deactivate entities (soft delete)
	if len(toDeactivate) > 0 {
		if err := deleteWhere(tx.Model(new(T)), toDeactivate).Delete(new(T)).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting(entityName), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	// Reactivate soft-deleted entities
	if len(toReactivate) > 0 {
		if err := deleteWhere(tx.Unscoped().Model(new(T)), toReactivate).
			Update("deleted_at", nil).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating(entityName), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	// Add new entities
	if len(toAdd) > 0 {
		if err := tx.Model(new(T)).Create(&toAdd).Error; err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting(entityName), log.TraceError(commonLogFields, err)...)
			return err
		}
	}

	// Execute extra operations (e.g., for images)
	for _, op := range extraOps {
		if err := op(tx, newIDs); err != nil {
			return err
		}
	}

	return nil
}

// updatePropertyAmenities updates amenities for a property
func (r *propertyRepository) updatePropertyAmenities(tx *gorm.DB, propertyID uint, amenityIDs []int) error {
	getExisting := func(tx *gorm.DB, propertyID uint) ([]dto.PropertyAmenity, error) {
		var existing []dto.PropertyAmenity
		err := tx.Unscoped().Where(&dto.PropertyAmenity{PropertyID: propertyID}).Find(&existing).Error
		return existing, err
	}

	compareKey := func(a dto.PropertyAmenity) int {
		return int(a.AmenityID)
	}

	isDeleted := func(a dto.PropertyAmenity) bool {
		return a.DeletedAt.Valid
	}

	createNew := func(id int) dto.PropertyAmenity {
		return dto.PropertyAmenity{
			PropertyID: propertyID,
			AmenityID:  uint(id),
		}
	}

	deleteWhere := func(tx *gorm.DB, ids []int) *gorm.DB {
		uintIDs := make([]uint, len(ids))
		for i, id := range ids {
			uintIDs[i] = uint(id)
		}
		return tx.Where(&dto.PropertyAmenity{PropertyID: propertyID}).Where("amenity_id IN ?", uintIDs)
	}

	return updateRelatedEntities(
		tx,
		propertyID,
		amenityIDs,
		getExisting,
		compareKey,
		isDeleted,
		createNew,
		deleteWhere,
		"PropertyAmenity",
		r.repositoryContext.RequestID,
	)
}

// updatePropertyUtilities updates utilities for a property
func (r *propertyRepository) updatePropertyUtilities(tx *gorm.DB, propertyID uint, utilityIDs []int) error {
	getExisting := func(tx *gorm.DB, propertyID uint) ([]dto.PropertyUtility, error) {
		var existing []dto.PropertyUtility
		err := tx.Unscoped().Where(&dto.PropertyUtility{PropertyID: propertyID}).Find(&existing).Error
		return existing, err
	}

	compareKey := func(u dto.PropertyUtility) int {
		return int(u.UtilityID)
	}

	isDeleted := func(u dto.PropertyUtility) bool {
		return u.DeletedAt.Valid
	}

	createNew := func(id int) dto.PropertyUtility {
		return dto.PropertyUtility{
			PropertyID: propertyID,
			UtilityID:  uint(id),
		}
	}

	deleteWhere := func(tx *gorm.DB, ids []int) *gorm.DB {
		uintIDs := make([]uint, len(ids))
		for i, id := range ids {
			uintIDs[i] = uint(id)
		}
		return tx.Where(&dto.PropertyUtility{PropertyID: propertyID}).Where("utility_id IN ?", uintIDs)
	}

	return updateRelatedEntities(
		tx,
		propertyID,
		utilityIDs,
		getExisting,
		compareKey,
		isDeleted,
		createNew,
		deleteWhere,
		"PropertyUtility",
		r.repositoryContext.RequestID,
	)
}

// updatePropertyImages updates images for a property
func (r *propertyRepository) updatePropertyImages(tx *gorm.DB, propertyID uint, imageURLs []string) error {
	getExisting := func(tx *gorm.DB, propertyID uint) ([]dto.PropertyImage, error) {
		var existing []dto.PropertyImage
		err := tx.Unscoped().Where(&dto.PropertyImage{PropertyID: propertyID}).Find(&existing).Error
		return existing, err
	}

	compareKey := func(i dto.PropertyImage) string {
		return i.URL
	}

	isDeleted := func(i dto.PropertyImage) bool {
		return i.DeletedAt.Valid
	}

	createNew := func(url string) dto.PropertyImage {
		return dto.PropertyImage{
			PropertyID: propertyID,
			URL:        url,
			IsPrimary:  false, // Will be updated in extraOps
		}
	}

	deleteWhere := func(tx *gorm.DB, urls []string) *gorm.DB {
		return tx.Where(&dto.PropertyImage{PropertyID: propertyID}).Where("url IN ?", urls)
	}

	extraOps := []func(tx *gorm.DB, newURLs []string) error{
		func(tx *gorm.DB, newURLs []string) error {
			if len(newURLs) == 0 {
				return nil
			}
			// Reset all images to non-primary
			if err := tx.Model(&dto.PropertyImage{}).Where(&dto.PropertyImage{PropertyID: propertyID}).Update("is_primary", false).Error; err != nil {
				log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("PropertyImage"), log.TraceError(log.CommonLogField(r.repositoryContext.RequestID), err)...)
				return err
			}
			// Set the first image as primary
			if err := tx.Model(&dto.PropertyImage{}).Where(&dto.PropertyImage{PropertyID: propertyID, URL: newURLs[0]}).Update("is_primary", true).Error; err != nil {
				log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("PropertyImage"), log.TraceError(log.CommonLogField(r.repositoryContext.RequestID), err)...)
				return err
			}
			return nil
		},
	}

	return updateRelatedEntities(
		tx,
		propertyID,
		imageURLs,
		getExisting,
		compareKey,
		isDeleted,
		createNew,
		deleteWhere,
		"PropertyImage",
		r.repositoryContext.RequestID,
		extraOps...,
	)
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
