package repository

import (
	"github.com/chazool/serendib_asia_service/app/routes/dto"
	"github.com/chazool/serendib_asia_service/pkg/config/dbconfig"
	"github.com/chazool/serendib_asia_service/pkg/log"

	"gorm.io/gorm"
)

const (
	// User repository methods
	UserRepositoryRegisterMethod         = "UserRepositoryRegister"
	UserRepositoryLoginMethod            = "UserRepositoryLogin"
	UserRepositoryGetProfileMethod       = "UserRepositoryGetProfile"
	UserRepositoryUpdateProfileMethod    = "UserRepositoryUpdateProfile"
	UserRepositoryUpdatePasswordMethod   = "UserRepositoryUpdatePassword"
	UserRepositoryCheckEmailExistsMethod = "UserRepositoryCheckEmailExists"
)

type UserRepository interface {
	Register(request *dto.UserRegisterRequest) (*dto.UserProfileResponse, error)
	Login(email, password string) (*dto.UserProfileResponse, error)
	GetProfile(userID uint) (*dto.UserProfileResponse, error)
	UpdateProfile(userID uint, request *dto.UserUpdateProfileRequest) (*dto.UserProfileResponse, error)
	UpdatePassword(userID uint, currentPassword, newPassword string) error
	CheckEmailExists(email string) (bool, error)
}

type userRepository struct {
	_                 struct{}
	repositoryContext Context
	db                *gorm.DB
}

// CreateUserRepository creates a new instance of UserRepository
func CreateUserRepository(requestID string) UserRepository {
	return &userRepository{
		repositoryContext: CreateRepositoryContext(requestID),
		db:                dbconfig.GetDBConnection(),
	}
}

func (r *userRepository) Register(request *dto.UserRegisterRequest) (*dto.UserProfileResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryRegisterMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryRegisterMethod), commonLogFields...)

	// TODO: Hash password before storing
	user := &dto.UserProfileResponse{
		FullName:    request.FullName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}

	err := r.db.Table("users").Create(user).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("User"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Login(email, password string) (*dto.UserProfileResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryLoginMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryLoginMethod), commonLogFields...)

	var user dto.UserProfileResponse
	err := r.db.Table("users").
		Where("email = ? AND password_hash = ?", email, password). // TODO: Use proper password hashing
		First(&user).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("User"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetProfile(userID uint) (*dto.UserProfileResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryGetProfileMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryGetProfileMethod), commonLogFields...)

	var user dto.UserProfileResponse
	err := r.db.Table("users").
		Where("id = ?", userID).
		First(&user).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("UserProfile"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateProfile(userID uint, request *dto.UserUpdateProfileRequest) (*dto.UserProfileResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryUpdateProfileMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryUpdateProfileMethod), commonLogFields...)

	updates := map[string]interface{}{
		"full_name":     request.FullName,
		"phone_number":  request.PhoneNumber,
		"profile_image": request.ProfileImage,
	}

	err := r.db.Table("users").
		Where("id = ?", userID).
		Updates(updates).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("UserProfile"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	return r.GetProfile(userID)
}

func (r *userRepository) UpdatePassword(userID uint, currentPassword, newPassword string) error {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryUpdatePasswordMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryUpdatePasswordMethod), commonLogFields...)

	// TODO: Hash passwords before comparing and storing
	err := r.db.Table("users").
		Where("id = ? AND password_hash = ?", userID, currentPassword).
		Update("password_hash", newPassword).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenUpdating("UserPassword"), log.TraceError(commonLogFields, err)...)
		return err
	}

	return nil
}

func (r *userRepository) CheckEmailExists(email string) (bool, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryCheckEmailExistsMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryCheckEmailExistsMethod), commonLogFields...)

	var count int64
	err := r.db.Table("users").
		Where("email = ?", email).
		Count(&count).Error

	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenCounting("UserEmail"), log.TraceError(commonLogFields, err)...)
		return false, err
	}

	return count > 0, nil
}
