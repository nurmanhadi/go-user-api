package repository

import (
	"context"
	"database/sql"
	"go-user-api/internal/entity"
	"go-user-api/internal/repository/source"

	"github.com/sirupsen/logrus"
)

type IUserRepository interface {
	Add(user *entity.User) error
	Count(email string) (int, error)
	FindById(id string) (*entity.User, error)
	Delete(id string) error
	FindByEmail(email string) (*entity.User, error)
	UpdatePassword(id string, password string) error
}
type userRepository struct {
	db  *sql.DB
	ctx context.Context
	log *logrus.Logger
}

func NewUserRepository(db *sql.DB, ctx context.Context, log *logrus.Logger) IUserRepository {
	return &userRepository{
		db:  db,
		ctx: ctx,
		log: log,
	}
}
func (r *userRepository) Add(user *entity.User) error {
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_ADD)
	if err != nil {
		r.log.WithError(err).Error("failed prepare context add user")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		r.log.WithError(err).Error("failed exec context add user")
		return err
	}
	return nil
}
func (r *userRepository) Count(email string) (int, error) {
	var count int
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_COUNT_BY_EMAIL)
	if err != nil {
		r.log.WithError(err).Error("failed prepare context count user")
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(r.ctx, &email).Scan(&count)
	if err != nil {
		r.log.WithError(err).Error("failed query row context count user")
		return 0, err
	}
	return count, nil
}
func (r *userRepository) FindById(id string) (*entity.User, error) {
	user := new(entity.User)
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_FIND_BY_ID)
	if err != nil {
		r.log.WithError(err).Error("failed prepare context user find by id")
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(r.ctx, &id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		r.log.WithError(err).Error("failed query row context user find by id")
		return nil, err
	}
	return user, nil
}
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_FIND_BY_EMAIL)
	if err != nil {
		r.log.WithError(err).Error("failed prepare context user find by email")
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(r.ctx, &email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		r.log.WithError(err).Error("failed query row context user find by email")
		return nil, err
	}
	return user, nil
}
func (r *userRepository) Delete(id string) error {
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_DELETE_BY_ID)
	if err != nil {
		r.log.WithError(err).Error("failed prepare context delete user")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &id)
	if err != nil {
		r.log.WithError(err).Error("failed exec context delete user")
		return err
	}
	return nil
}
func (r *userRepository) UpdatePassword(id string, password string) error {
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_UPDATE_PASSWORD)
	if err != nil {
		r.log.WithError(err).Error("failed prepare context update password user")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &password, &id)
	if err != nil {
		r.log.WithError(err).Error("failed exec context update password user")
		return err
	}
	return nil
}
