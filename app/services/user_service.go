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
	// User service methods
	UserServiceRegisterMethod       = "UserServiceRegister"
	UserServiceLoginMethod          = "UserServiceLogin"
	UserServiceGetProfileMethod     = "UserServiceGetProfile"
	UserServiceUpdateProfileMethod  = "UserServiceUpdateProfile"
	UserServiceUpdatePasswordMethod = "UserServiceUpdatePassword"
)

// UserService defines the interface for user service methods
type UserService struct {
	_              struct{}
	serviceContext ServiceContext
	transaction    *gorm.DB
	userRepo       repository.UserRepository
}

// CreateUserService creates a new instance of UserService
func CreateUserService(requestID string, transactionDB *gorm.DB) *UserService {
	return &UserService{
		serviceContext: CreateServiceContext(requestID),
		transaction:    transactionDB,
	}
}

// Register handles user registration
func (service *UserService) Register(request *dto.UserRegisterRequest) (response *dto.UserLoginResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserServiceRegisterMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(UserServiceRegisterMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(UserServiceRegisterMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.userRepo = repository.CreateUserRepository(service.serviceContext.RequestID)

	// Check if email already exists
	exists, err := service.userRepo.CheckEmailExists(request.Email)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.UserRepositoryCheckEmailExistsMethod), logFields...)
		return nil, buildSelectErrFromRepo("user email", err)
	}
	if exists {
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(UserServiceRegisterMethod), commonLogFields...)
		errRes := custom.BuildBadReqErrResult(constant.DuplicateEmailErrorCode, constant.DuplicateEmailErrorMessage, "Email")
		return nil, &errRes
	}

	// Register user
	user, err := service.userRepo.Register(request)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.UserRepositoryRegisterMethod), logFields...)
		return nil, buildInsertErrFromRepo("user", err)
	}

	// TODO: Generate JWT token
	token := "dummy-token" // Replace with actual JWT token generation

	response = &dto.UserLoginResponse{
		User:  *user,
		Token: token,
	}

	return response, nil
}

// Login handles user login
func (service *UserService) Login(request *dto.UserLoginRequest) (response *dto.UserLoginResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserServiceLoginMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(UserServiceLoginMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(UserServiceLoginMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.userRepo = repository.CreateUserRepository(service.serviceContext.RequestID)

	// Login user
	user, err := service.userRepo.Login(request.Email, request.Password)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.UserRepositoryLoginMethod), logFields...)
		errRes := custom.BuildBadReqErrResult(constant.InvalidCredentialsCode, constant.InvalidCredentialsMessage, "Credentials")
		return nil, &errRes
	}

	// TODO: Generate JWT token
	token := "dummy-token" // Replace with actual JWT token generation

	response = &dto.UserLoginResponse{
		User:  *user,
		Token: token,
	}

	return response, nil
}

// GetProfile retrieves user profile
func (service *UserService) GetProfile(userID uint) (response *dto.UserProfileResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserServiceGetProfileMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(UserServiceGetProfileMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(UserServiceGetProfileMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.userRepo = repository.CreateUserRepository(service.serviceContext.RequestID)

	response, err := service.userRepo.GetProfile(userID)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.UserRepositoryGetProfileMethod), logFields...)
		errRes := custom.BuildBadReqErrResult(constant.UserNotFoundCode, constant.UserNotFoundMessage, "User")
		return nil, &errRes
	}

	return response, nil
}

// UpdateProfile updates user profile
func (service *UserService) UpdateProfile(userID uint, request *dto.UserUpdateProfileRequest) (response *dto.UserProfileResponse, errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserServiceUpdateProfileMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(UserServiceUpdateProfileMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(UserServiceUpdateProfileMethod), log.TraceMethodOutputs(commonLogFields, response, errResult)...)
	}()

	service.userRepo = repository.CreateUserRepository(service.serviceContext.RequestID)

	response, err := service.userRepo.UpdateProfile(userID, request)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.UserRepositoryUpdateProfileMethod), logFields...)
		return nil, buildUpdateErrFromRepo("user profile", err)
	}

	return response, nil
}

// UpdatePassword updates user password
func (service *UserService) UpdatePassword(userID uint, request *dto.UserUpdatePasswordRequest) (errResult *custom.ErrorResult) {
	commonLogFields := log.CommonLogField(service.serviceContext.RequestID)
	log.Logger.Debug(log.TraceMsgFuncStart(UserServiceUpdatePasswordMethod), commonLogFields...)

	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error(constant.PanicOccurred, log.TraceStack(commonLogFields, debug.Stack())...)
			errResult = buildPanicErr(UserServiceUpdatePasswordMethod)
		}
		log.Logger.Debug(log.TraceMsgFuncEnd(UserServiceUpdatePasswordMethod), log.TraceMethodOutputs(commonLogFields, nil, errResult)...)
	}()

	service.userRepo = repository.CreateUserRepository(service.serviceContext.RequestID)

	err := service.userRepo.UpdatePassword(userID, request.CurrentPassword, request.NewPassword)
	if err != nil {
		logFields := log.TraceError(commonLogFields, err)
		log.Logger.Error(log.TraceMsgErrorOccurredFrom(repository.UserRepositoryUpdatePasswordMethod), logFields...)
		return buildUpdateErrFromRepo("user password", err)
	}

	return nil
}
