package repository

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/config/dbconfig"
	"github.com/chazool/serendib_asia_service/pkg/log"

	"gorm.io/gorm"
)

const (
	// Favorite repository methods
	FavoriteRepositoryAddMethod    = "FavoriteRepositoryAdd"
	FavoriteRepositoryRemoveMethod = "FavoriteRepositoryRemove"
	FavoriteRepositoryListMethod   = "FavoriteRepositoryList"
)

type FavoriteRepository interface {
	Add(userID, propertyID uint) error
	Remove(userID, propertyID uint) error
	List(userID uint, page, pageSize int) ([]dto.FavoriteResponse, int64, error)
}

type favoriteRepository struct {
	_                 struct{}
	repositoryContext Context
	db                *gorm.DB
}

// CreateFavoriteRepository creates a new instance of FavoriteRepository
func CreateFavoriteRepository(requestID string) FavoriteRepository {
	return &favoriteRepository{
		repositoryContext: CreateRepositoryContext(requestID),
		db:                dbconfig.GetDBConnection(),
	}
}

func (r *favoriteRepository) Add(userID, propertyID uint) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteRepositoryAddMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteRepositoryAddMethod), commonLogFields...)

	err := r.db.Table("favourites").Create(map[string]interface{}{
		"user_id":     userID,
		"property_id": propertyID,
	}).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("Favorite"), log.TraceError(commonLogFields, err)...)
		return err
	}
	return nil
}

func (r *favoriteRepository) Remove(userID, propertyID uint) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteRepositoryRemoveMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteRepositoryRemoveMethod), commonLogFields...)

	err := r.db.Table("favourites").
		Where("user_id = ? AND property_id = ?", userID, propertyID).
		Delete(&struct{}{}).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenDeleting("Favorite"), log.TraceError(commonLogFields, err)...)
		return err
	}
	return nil
}

func (r *favoriteRepository) List(userID uint, page, pageSize int) ([]dto.FavoriteResponse, int64, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(FavoriteRepositoryListMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(FavoriteRepositoryListMethod), commonLogFields...)

	var favorites []dto.FavoriteResponse
	var total int64

	query := r.db.Table("favourites").
		Select("favourites.*, properties.*").
		Joins("JOIN properties ON properties.id = favourites.property_id").
		Where("favourites.user_id = ?", userID)

	err := query.Count(&total).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenCounting("Favorites"), log.TraceError(commonLogFields, err)...)
		return nil, 0, err
	}

	err = query.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Scan(&favorites).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("Favorites"), log.TraceError(commonLogFields, err)...)
		return nil, 0, err
	}

	return favorites, total, nil
}
