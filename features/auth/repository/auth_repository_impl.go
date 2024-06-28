package repository

import (
	"github.com/4kpros/go-api/common/types"
	"github.com/4kpros/go-api/common/utils"
	"github.com/4kpros/go-api/features/auth/model"
	userModel "github.com/4kpros/go-api/features/user/model"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	Db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{Db: db}
}

func (repository *AuthRepositoryImpl) Create(user *model.NewUser) error {
	result := repository.Db.Create(user)
	return result.Error
}

func (repository *AuthRepositoryImpl) CreateUserAccountDetails(user *userModel.User) error {
	result := repository.Db.Create(user)
	return result.Error
}

func (repository *AuthRepositoryImpl) Update(user *model.NewUser) error {
	result := repository.Db.Model(user).Updates(user)
	return result.Error
}

func (repository *AuthRepositoryImpl) UpdatePasswordById(id string, password string) (*model.NewUser, error) {
	var user = &model.NewUser{
		Password: password,
	}
	result := repository.Db.Model(user).Where("id = ?", id).Update("password", user)
	return user, result.Error
}

func (repository *AuthRepositoryImpl) Delete(id string) (int64, error) {
	var user = &model.NewUser{}
	result := repository.Db.Where("id = ?", id).Delete(user)
	return result.RowsAffected, result.Error
}

func (repository *AuthRepositoryImpl) FindAll(filter *types.Filter, pagination *types.Pagination) ([]model.NewUser, error) {
	var users = []model.NewUser{}
	result := repository.Db.Scopes(utils.PaginationScope(users, pagination, filter, repository.Db)).Find(users)
	return users, result.Error
}

func (repository *AuthRepositoryImpl) FindById(id string) (*model.NewUser, error) {
	var user = &model.NewUser{}
	result := repository.Db.Where("id = ?", id).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepositoryImpl) FindByEmail(email string) (*model.NewUser, error) {
	var user = &model.NewUser{}
	result := repository.Db.Where("email = ? AND (provider is null OR provider = '')", email).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepositoryImpl) FindByPhoneNumber(phoneNumber int) (*model.NewUser, error) {
	var user = &model.NewUser{}
	result := repository.Db.Where("phoneNumber = ? AND (provider is null OR provider = '')", phoneNumber).Limit(1).Find(user)
	return user, result.Error
}

func (repository *AuthRepositoryImpl) FindByProvider(provider string, providerUserId string) (*model.NewUser, error) {
	var user = &model.NewUser{}
	result := repository.Db.Where("provider = ? AND providerUserId = ?", provider, providerUserId).Limit(1).Find(user)
	return user, result.Error
}
