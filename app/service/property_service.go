package service

import (
	"context"

	"github.com/chazool/serendib_asia_service/app/repository"
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"
)

type propertyService struct {
	userRepo repository.UserRepository
}

func (s *propertyService) Create(ctx context.Context, request dto.PropertyRequest) (*dto.Property, *custom.ErrorResult) {
	// Validate user exists
	user, err := s.userRepo.GetProfile(request.UserID)
	if err != nil {
		if errResult, ok := err.(*custom.ErrorResult); ok {
			return nil, errResult
		}
		errRes := custom.BuildInternalServerErrResult(constant.ErrCodeInternalServerError, err.Error(), constant.Empty)
		return nil, &errRes
	}
	if user == nil {
		errRes := custom.BuildNotFoundErrResult(constant.UserNotFoundCode, constant.UserNotFoundMessage, constant.Empty)
		return nil, &errRes
	}

	// Validate images count
	if len(request.Images) == 0 {
		errRes := custom.BuildBadReqErrResult(constant.ErrCodeInvalidInput, "at least one property image is required", constant.Empty)
		return nil, &errRes
	}
	if len(request.Images) > 6 {
		errRes := custom.BuildBadReqErrResult(constant.ErrCodeInvalidInput, "maximum of 6 property images allowed", constant.Empty)
		return nil, &errRes
	}

	// Validate purpose type exists
	// ... existing code ...

	return nil, nil
}
