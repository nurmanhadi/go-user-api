package usecase

import (
	"go-user-api/internal/entity"
	"go-user-api/internal/repository"
	"go-user-api/pkg/exception"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type IUserUsecase interface {
	FindById(id string) (*entity.User, error)
	Delete(id string) error
}
type userUsecase struct {
	userRepository repository.IUserRepository
	validation     *validator.Validate
	log            *logrus.Logger
}

func NewUserUsecase(userRepository repository.IUserRepository, validation *validator.Validate, log *logrus.Logger) IUserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		validation:     validation,
		log:            log,
	}
}
func (u *userUsecase) FindById(id string) (*entity.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		u.log.WithField("error", err).Warn("user not found")
		return nil, exception.UserNotFound
	}
	return user, nil
}
func (u *userUsecase) Delete(id string) error {
	_, err := u.userRepository.FindById(id)
	if err != nil {
		u.log.WithField("error", err).Warn("user not found")
		return exception.UserNotFound
	}
	err = u.userRepository.Delete(id)
	if err != nil {
		u.log.WithField("error", err).Warn("failed delete user from database")
		return err
	}
	return nil
}
