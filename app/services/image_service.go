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
	// Image service methods
	ImageServiceUploadMethod     = "ImageServiceUpload"
	ImageServiceDeleteMethod     = "ImageServiceDelete"
	ImageServiceSetPrimaryMethod = "ImageServiceSetPrimary"
	ImageServiceListMethod       = "ImageServiceList"
)

// ImageService defines the interface for image service methods
type ImageService struct {
	_              struct{}
	serviceContext ServiceContext
	transaction    *gorm.DB
	imageRepo      repository.ImageRepository
}

// CreateImageService creates a new instance of ImageService
func CreateImageService(requestID string, transactionDB *gorm.DB) *ImageService {
	return &ImageService{
		serviceContext: CreateServiceContext(requestID),
		transaction:    transactionDB,
	}
}

// Upload uploads a property image
func (service *ImageService) Upload(propertyID uint, request *dto.UploadImageRequest) (response *dto.ImageResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ImageServiceUploadMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(ImageServiceUploadMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(ImageServiceUploadMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.imageRepo = repository.CreateImageRepository(service.serviceContext.RequestID)

	image, err := service.imageRepo.Upload(propertyID, request.URL, request.IsPrimary)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.ImageRepositoryUploadMethod), logFields...)
		return nil, buildInsertErrFromRepo("image", err)
	}

	response = &dto.ImageResponse{
		ID:         image.ID,
		PropertyID: image.PropertyID,
		URL:        image.URL,
		IsPrimary:  image.IsPrimary,
	}

	return response, nil
}

// Delete deletes a property image
func (service *ImageService) Delete(imageID uint) (errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ImageServiceDeleteMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(ImageServiceDeleteMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(ImageServiceDeleteMethod), log.TraceMethodOutputs(commonLogFields, nil, errResult)...)
	}()

	service.imageRepo = repository.CreateImageRepository(service.serviceContext.RequestID)

	err := service.imageRepo.Delete(imageID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.ImageRepositoryDeleteMethod), logFields...)
		return buildDeleteErrFromRepo("image", err)
	}

	return nil
}

// SetPrimary sets an image as primary
func (service *ImageService) SetPrimary(imageID uint) (errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ImageServiceSetPrimaryMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(ImageServiceSetPrimaryMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(ImageServiceSetPrimaryMethod), log.TraceMethodOutputs(commonLogFields, nil, errResult)...)
	}()

	service.imageRepo = repository.CreateImageRepository(service.serviceContext.RequestID)

	err := service.imageRepo.SetPrimary(imageID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.ImageRepositorySetPrimaryMethod), logFields...)
		return buildUpdateErrFromRepo("image", err)
	}

	return nil
}

// List lists all images for a property
func (service *ImageService) List(propertyID uint) (response *dto.ImageListResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(ImageServiceListMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(ImageServiceListMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(ImageServiceListMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.imageRepo = repository.CreateImageRepository(service.serviceContext.RequestID)

	images, err := service.imageRepo.List(propertyID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.ImageRepositoryListMethod), logFields...)
		return nil, buildSelectErrFromRepo("images", err)
	}

	response = &dto.ImageListResponse{
		Items: images,
	}

	return response, nil
}
