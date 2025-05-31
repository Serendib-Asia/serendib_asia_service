package repository

import (
	appdto "github.com/chazool/serendib_asia_service/app/routes/dto"
	internaldto "github.com/chazool/serendib_asia_service/internal/dto"
	"github.com/chazool/serendib_asia_service/pkg/config/dbconfig"
	"github.com/chazool/serendib_asia_service/pkg/log"
	"golang.org/x/crypto/bcrypt"

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
	Register(request *appdto.UserRegisterRequest) (*appdto.UserProfileResponse, error)
	Login(email, password string) (*appdto.UserProfileResponse, error)
	GetProfile(userID uint) (*appdto.UserProfileResponse, error)
	UpdateProfile(userID uint, request *appdto.UserUpdateProfileRequest) (*appdto.UserProfileResponse, error)
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

func (r *userRepository) Register(request *appdto.UserRegisterRequest) (*appdto.UserProfileResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryRegisterMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryRegisterMethod), commonLogFields...)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhen("hashing password"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	// Create user
	newUser := &internaldto.User{
		FullName:     request.Name,
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
	}

	err = r.db.Create(newUser).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenInserting("User"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	// Convert to response DTO
	response := &appdto.UserProfileResponse{
		ID:           newUser.ID,
		FullName:     newUser.FullName,
		Email:        newUser.Email,
		PhoneNumber:  newUser.PhoneNumber,
		ProfileImage: newUser.ProfileImage,
		CreatedAt:    newUser.CreatedAt,
	}

	return response, nil
}

func (r *userRepository) Login(email, password string) (*appdto.UserProfileResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryLoginMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryLoginMethod), commonLogFields...)

	var user internaldto.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("User"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhen("verifying password"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	// Convert to response DTO
	response := &appdto.UserProfileResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		ProfileImage: user.ProfileImage,
		CreatedAt:    user.CreatedAt,
	}

	return response, nil
}

func (r *userRepository) GetProfile(userID uint) (*appdto.UserProfileResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryGetProfileMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryGetProfileMethod), commonLogFields...)

	var user internaldto.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("UserProfile"), log.TraceError(commonLogFields, err)...)
		return nil, err
	}

	// Convert to response DTO
	response := &appdto.UserProfileResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		ProfileImage: user.ProfileImage,
		CreatedAt:    user.CreatedAt,
	}

	return response, nil
}

func (r *userRepository) UpdateProfile(userID uint, request *appdto.UserUpdateProfileRequest) (*appdto.UserProfileResponse, error) {
	commonLogFields := log.CommonLogField(r.repositoryContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserRepositoryUpdateProfileMethod), commonLogFields...)
	defer log.Logger.Debug(log.TraceMsgFuncEnd(UserRepositoryUpdateProfileMethod), commonLogFields...)

	updates := map[string]interface{}{
		"full_name":     request.FullName,
		"phone_number":  request.PhoneNumber,
		"profile_image": request.ProfileImage,
	}

	err := r.db.Model(&internaldto.User{}).
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

	// Get current user
	var user internaldto.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenSelecting("UserPassword"), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Verify current password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(currentPassword))
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhen("verifying current password"), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhen("hashing new password"), log.TraceError(commonLogFields, err)...)
		return err
	}

	// Update password
	err = r.db.Model(&user).Update("password_hash", string(hashedPassword)).Error
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
	err := r.db.Model(&internaldto.User{}).
		Where("email = ?", email).
		Count(&count).Error
	if err != nil {
		log.Logger.Error(log.TraceMsgErrorOccurredWhenCounting("UserEmail"), log.TraceError(commonLogFields, err)...)
		return false, err
	}

	return count > 0, nil
}
