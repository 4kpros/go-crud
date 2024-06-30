package repository

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/features/user/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

// Create implements UserRepository.
func (repository *UserRepositoryImpl) Create(user *model.User) error {
	return repository.Db.Create(user).Error
}

// Update implements UserRepository.
func (repository *UserRepositoryImpl) Update(user *model.User) error {
	return repository.Db.Model(user).Updates(user).Error
}

// Delete implements UserRepository.
func (repository *UserRepositoryImpl) Delete(id string) (int64, error) {
	var user = &model.User{}
	result := repository.Db.Where("id = ?", id).Delete(user)
	return result.RowsAffected, result.Error
}

// FindAll implements UserRepository.
func (repository *UserRepositoryImpl) FindAll(filter *types.Filter, pagination *types.Pagination) ([]model.User, error) {
	var users = []model.User{}
	result := repository.Db.Scopes(utils.PaginationScope(users, pagination, filter, repository.Db)).Find(users)
	return users, result.Error
}

// FindById implements UserRepository.
func (repository *UserRepositoryImpl) FindById(id string) (*model.User, error) {
	var user = &model.User{}
	result := repository.Db.Where("id = ?", id).Limit(1).Find(user)
	return user, result.Error
}
