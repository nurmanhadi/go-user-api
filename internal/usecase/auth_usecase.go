package usecase

import (
	"go-user-api/internal/entity"
	"go-user-api/internal/model"
	"go-user-api/internal/repository"
	"go-user-api/pkg/exception"
	"go-user-api/pkg/hash"
	"go-user-api/pkg/token"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type IAuthUsecase interface {
	Register(request *model.AuthRegisterRequest) error
	Login(request *model.AuthLoginRequest) (*model.AuthResponseAccessToken, error)
}
type authUsecase struct {
	userRepository repository.IUserRepository
	validation     *validator.Validate
	log            *logrus.Logger
	viper          *viper.Viper
}

func NewAuthUsecase(userRepository repository.IUserRepository, validation *validator.Validate, log *logrus.Logger, viper *viper.Viper) IAuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
		validation:     validation,
		log:            log,
		viper:          viper,
	}
}
func (u *authUsecase) Register(request *model.AuthRegisterRequest) error {
	err := u.validation.Struct(request)
	if err != nil {
		u.log.WithField("error", err).Warn("failed validation register request")
		return err
	}
	email := strings.ToLower(request.Email)
	total, err := u.userRepository.Count(email)
	if err != nil {
		u.log.WithField("error", err).Warn("failed count user from database")
		return err
	}
	if total > 0 {
		u.log.WithField("error", total).Warn("email already exists")
		return exception.EmailAlreadyExists
	}
	hashPassword, err := hash.Password(request.Password)
	if err != nil {
		u.log.WithField("error", err).Warn("failed generate hash password")
		return err
	}
	id := uuid.NewString()
	user := &entity.User{
		Id:       id,
		Name:     request.Name,
		Email:    email,
		Password: hashPassword,
	}
	err = u.userRepository.Add(user)
	if err != nil {
		u.log.WithField("error", err).Warn("failed add user to database")
		return err
	}
	return nil
}
func (u *authUsecase) Login(request *model.AuthLoginRequest) (*model.AuthResponseAccessToken, error) {
	err := u.validation.Struct(request)
	if err != nil {
		u.log.WithField("error", err).Warn("failed validation login request")
		return nil, err
	}
	email := strings.ToLower(request.Email)
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		u.log.WithField("error", err).Warn("email or password wrong")
		return nil, exception.EmailOrPasswordWrong
	}
	err = hash.ComparePassword(user.Password, request.Password)
	if err != nil {
		u.log.WithField("error", err).Warn("email or password wrong")
		return nil, exception.EmailOrPasswordWrong
	}
	token, err := token.GenerateAccessToken(user.Id, u.viper)
	if err != nil {
		u.log.WithField("error", err).Warn("failed generate access token")
		return nil, err
	}
	result := &model.AuthResponseAccessToken{
		AccessToken: token,
	}
	return result, nil
}
