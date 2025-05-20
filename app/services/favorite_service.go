package services

import (
	"github.com/chazool/serendib_asia_service/app/repository"
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/custom"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"github.com/chazool/serendib_asia_service/pkg/utils/constant"
)

const (
	// Favorite service methods
	FavoriteServiceAddMethod    = "FavoriteServiceAdd"
	FavoriteServiceRemoveMethod = "FavoriteServiceRemove"
	FavoriteServiceListMethod   = "FavoriteServiceList"
)

type FavoriteService interface {
	AddFavorite(userID, propertyID uint) *custom.ErrorResult
	RemoveFavorite(userID, propertyID uint) *custom.ErrorResult
	ListFavorites(userID uint) ([]dto.FavoriteResponse, *custom.ErrorResult)
}

type favoriteService struct {
	_              struct{}
	serviceContext ServiceContext
	favoriteRepo   repository.FavoriteRepository
	propertyRepo   repository.PropertyRepository
}

// CreateFavoriteService creates a new instance of FavoriteService
func CreateFavoriteService(requestID string) FavoriteService {
	return &favoriteService{
		serviceContext: CreateServiceContext(requestID),
		favoriteRepo:   repository.CreateFavoriteRepository(requestID),
		propertyRepo:   repository.CreatePropertyRepository(requestID),
	}
}

func (s *favoriteService) AddFavorite(userID, propertyID uint) *custom.ErrorResult {
	commonLogFields := log.CommonLogField(s.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteServiceAddMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteServiceAddMethod), commonLogFields...)

	// Check if property exists
	exists, err := s.propertyRepo.CheckExists(propertyID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteServiceAddMethod), log.TraceError(commonLogFields, err)...)
		return buildSelectErrFromRepo("Property", err)
	}
	if !exists {
		errRes := custom.BuildBadReqErrResult(constant.ErrRecordNotFoundCode, constant.ErrRecordNotFoundMsg, "Property")
		return &errRes
	}

	// Add to favorites
	err = s.favoriteRepo.Add(userID, propertyID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteServiceAddMethod), log.TraceError(commonLogFields, err)...)
		return buildInsertErrFromRepo("Favorite", err)
	}

	return nil
}

func (s *favoriteService) RemoveFavorite(userID, propertyID uint) *custom.ErrorResult {
	commonLogFields := log.CommonLogField(s.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteServiceRemoveMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteServiceRemoveMethod), commonLogFields...)

	// Remove from favorites
	err := s.favoriteRepo.Remove(userID, propertyID)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteServiceRemoveMethod), log.TraceError(commonLogFields, err)...)
		return buildDeleteErrFromRepo("Favorite", err)
	}

	return nil
}

func (s *favoriteService) ListFavorites(userID uint) ([]dto.FavoriteResponse, *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(s.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteServiceListMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteServiceListMethod), commonLogFields...)

	// Get favorites
	favorites, _, err := s.favoriteRepo.List(userID, 1, 100) // TODO: Add pagination
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(FavoriteServiceListMethod), log.TraceError(commonLogFields, err)...)
		return nil, buildSelectErrFromRepo("Favorites", err)
	}

	return favorites, nil
}
