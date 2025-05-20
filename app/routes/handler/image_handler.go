package handler

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/app/services"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"
	"github.com/chazool/serendib_asia_service/pkg/web"
	"github.com/chazool/serendib_asia_service/pkg/web/responsebuilder"

	"github.com/gofiber/fiber/v2"
)

const (
	// Image handler methods
	HandleUploadImageMethod     = "HandleUploadImage"
	HandleDeleteImageMethod     = "HandleDeleteImage"
	HandleSetPrimaryImageMethod = "HandleSetPrimaryImage"
	HandleListImagesMethod      = "HandleListImages"
)

// HandleUploadImage handles uploading a property image
// @Summary Upload a property image
// @Description Uploads an image for a property with optional primary flag
// @Tags properties
// @Accept multipart/form-data
// @Produce json
// @Param propertyId path int true "Property ID"
// @Param image formData file true "Image file"
// @Param isPrimary formData bool false "Set as primary image"
// @Success 200 {object} dto.ImageResponse
// @Failure 400 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties/{propertyId}/images [post]
func HandleUploadImage(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleUploadImageMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleUploadImageMethod), commonLogFields...)

	var (
		statusCode   int
		errorResult  *custom.ErrorResult
		errRes       custom.ErrorResult
		response     *dto.ImageResponse
		imageService = services.CreateImageService(requestID, nil)
	)

	propertyID, err := ctx.ParamsInt("propertyId")
	if err != nil || propertyID <= 0 {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleUploadImageMethod), commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, "Invalid property ID")
		errorResult = &errRes
		statusCode, errRes = HandleError(errorResult)
	} else {
		file, err := ctx.FormFile("image")
		if err != nil {
			logFields := log.TraceError(commonLogFields, err)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleUploadImageMethod), logFields...)
			errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, "Invalid image file")
			errorResult = &errRes
			statusCode, errRes = HandleError(errorResult)
		} else {
			// TODO: Implement file upload to storage and get URL
			imageURL := "https://example.com/images/" + file.Filename // Placeholder URL

			request := &dto.UploadImageRequest{
				URL:       imageURL,
				IsPrimary: ctx.FormValue("isPrimary") == "true",
			}

			response, errorResult = imageService.Upload(uint(propertyID), request)
			if errorResult != nil {
				logFields := log.TraceCustomError(commonLogFields, *errorResult)
				log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.ImageServiceUploadMethod), logFields...)
				statusCode, errRes = HandleError(errorResult)
			}
		}
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		Response:      response,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

// HandleDeleteImage handles deleting a property image
// @Summary Delete a property image
// @Description Deletes an image from a property
// @Tags properties
// @Accept json
// @Produce json
// @Param imageId path int true "Image ID"
// @Success 200 {object} custom.ErrorResult
// @Failure 400 {object} custom.ErrorResult
// @Failure 404 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties/images/{imageId} [delete]
func HandleDeleteImage(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleDeleteImageMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleDeleteImageMethod), commonLogFields...)

	var (
		statusCode   int
		errorResult  *custom.ErrorResult
		errRes       custom.ErrorResult
		imageService = services.CreateImageService(requestID, nil)
	)

	imageID, err := ctx.ParamsInt("imageId")
	if err != nil || imageID <= 0 {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleDeleteImageMethod), commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, "Invalid image ID")
		errorResult = &errRes
		statusCode, errRes = HandleError(errorResult)
	} else {
		errorResult = imageService.Delete(uint(imageID))
		if errorResult != nil {
			logFields := log.TraceCustomError(commonLogFields, *errorResult)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.ImageServiceDeleteMethod), logFields...)
			statusCode, errRes = HandleError(errorResult)
		}
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

// HandleSetPrimaryImage handles setting an image as primary
// @Summary Set primary image
// @Description Sets an image as the primary image for a property
// @Tags properties
// @Accept json
// @Produce json
// @Param imageId path int true "Image ID"
// @Success 200 {object} custom.ErrorResult
// @Failure 400 {object} custom.ErrorResult
// @Failure 404 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties/images/{imageId}/primary [put]
func HandleSetPrimaryImage(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleSetPrimaryImageMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleSetPrimaryImageMethod), commonLogFields...)

	var (
		statusCode   int
		errorResult  *custom.ErrorResult
		errRes       custom.ErrorResult
		imageService = services.CreateImageService(requestID, nil)
	)

	imageID, err := ctx.ParamsInt("imageId")
	if err != nil || imageID <= 0 {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleSetPrimaryImageMethod), commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, "Invalid image ID")
		errorResult = &errRes
		statusCode, errRes = HandleError(errorResult)
	} else {
		errorResult = imageService.SetPrimary(uint(imageID))
		if errorResult != nil {
			logFields := log.TraceCustomError(commonLogFields, *errorResult)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.ImageServiceSetPrimaryMethod), logFields...)
			statusCode, errRes = HandleError(errorResult)
		}
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}

// HandleListImages handles listing all images for a property
// @Summary List property images
// @Description Lists all images associated with a property
// @Tags properties
// @Accept json
// @Produce json
// @Param propertyId path int true "Property ID"
// @Success 200 {object} dto.ImageListResponse
// @Failure 400 {object} custom.ErrorResult
// @Failure 500 {object} custom.ErrorResult
// @Router /api/properties/{propertyId}/images [get]
func HandleListImages(ctx *fiber.Ctx) error {
	requestID := web.GetRequestID(ctx)
	commonLogFields := log.CommonLogField(requestID)
	log.Logger.Info(log.TraceMsgFuncStart(HandleListImagesMethod), commonLogFields...)
	defer log.Logger.Info(log.TraceMsgFuncEnd(HandleListImagesMethod), commonLogFields...)

	var (
		statusCode   int
		errorResult  *custom.ErrorResult
		errRes       custom.ErrorResult
		response     *dto.ImageListResponse
		imageService = services.CreateImageService(requestID, nil)
	)

	propertyID, err := ctx.ParamsInt("propertyId")
	if err != nil || propertyID <= 0 {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(HandleListImagesMethod), commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.BindingErrorCode, constant.InvalidRequestErrorMessage, "Invalid property ID")
		errorResult = &errRes
		statusCode, errRes = HandleError(errorResult)
	} else {
		response, errorResult = imageService.List(uint(propertyID))
		if errorResult != nil {
			logFields := log.TraceCustomError(commonLogFields, *errorResult)
			log.Logger.Error(log.TraceMsgErrorOccurredFrom(services.ImageServiceListMethod), logFields...)
			statusCode, errRes = HandleError(errorResult)
		}
	}

	responseBuilder := responsebuilder.APIResponse{
		Ctx:           ctx,
		HTTPStatus:    statusCode,
		ErrorResponse: errRes,
		Response:      response,
		RequestID:     requestID,
	}
	responseBuilder.BuildAPIResponse()

	return nil
}
