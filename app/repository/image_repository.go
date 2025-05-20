package repository

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/config/dbconfig"
	"github.com/chazool/serendib_asia_service/pkg/log"

	"gorm.io/gorm"
)

const (
	// Image repository methods
	ImageRepositoryUploadMethod     = "ImageRepositoryUpload"
	ImageRepositoryDeleteMethod     = "ImageRepositoryDelete"
	ImageRepositorySetPrimaryMethod = "ImageRepositorySetPrimary"
	ImageRepositoryListMethod       = "ImageRepositoryList"
)

type ImageRepository interface {
	Upload(propertyID uint, url string, isPrimary bool) (*dto.ImageResponse, error)
	Delete(imageID uint) error
	SetPrimary(imageID uint) error
	List(propertyID uint) ([]dto.ImageResponse, error)
}

type imageRepository struct {
	_                 struct{}
	repositoryContext Context
	db                *gorm.DB
}

// CreateImageRepository creates a new instance of ImageRepository
func CreateImageRepository(requestID string) ImageRepository {
	return &imageRepository{
		repositoryContext: CreateRepositoryContext(requestID),
		db:                dbconfig.GetDBConnection(),
	}
}

func (r *imageRepository) Upload(propertyID uint, url string, isPrimary bool) (*dto.ImageResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ImageRepositoryUploadMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ImageRepositoryUploadMethod), commonLogFields...)

	// If this is the primary image, unset any existing primary images
	if isPrimary {
		err := r.db.Table("property_images").
			Where("property_id = ?", propertyID).
			Update("is_primary", false).Error
		if err != nil {
			log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("PropertyImages"), log.TraceError(commonLogFields, err)...)
			return nil, err
		}
	}

	var response dto.ImageResponse
	err := r.db.Table("property_images").Create(map[string]interface{}{
		"property_id": propertyID,
		"url":         url,
		"is_primary":  isPrimary,
	}).Scan(&response).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("PropertyImage"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return &response, nil
}

func (r *imageRepository) Delete(imageID uint) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ImageRepositoryDeleteMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ImageRepositoryDeleteMethod), commonLogFields...)

	err := r.db.Table("property_images").
		Where("id = ?", imageID).
		Delete(&struct{}{}).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("PropertyImage"), log.TraceError(commonLogFields, err)...)
		return err
	}
	return nil
}

func (r *imageRepository) SetPrimary(imageID uint) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ImageRepositorySetPrimaryMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ImageRepositorySetPrimaryMethod), commonLogFields...)

	// Get the property ID for this image
	var propertyID uint
	err := r.db.Table("property_images").
		Select("property_id").
		Where("id = ?", imageID).
		Scan(&propertyID).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PropertyImage"), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Unset any existing primary images
	err = r.db.Table("property_images").
		Where("property_id = ?", propertyID).
		Update("is_primary", false).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("PropertyImages"), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Set the new primary image
	err = r.db.Table("property_images").
		Where("id = ?", imageID).
		Update("is_primary", true).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("PropertyImage"), log.TraceError(commonLogFields, err)...)
		return err
	}

	return nil
}

func (r *imageRepository) List(propertyID uint) ([]dto.ImageResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ImageRepositoryListMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(ImageRepositoryListMethod), commonLogFields...)

	var images []dto.ImageResponse
	err := r.db.Table("property_images").
		Where("property_id = ?", propertyID).
		Find(&images).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("PropertyImages"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}
	return images, nil
}
