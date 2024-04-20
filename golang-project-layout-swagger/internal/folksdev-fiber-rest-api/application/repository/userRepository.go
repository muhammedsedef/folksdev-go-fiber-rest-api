package repository

import (
	"context"
	"fmt"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/domain"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/pkg/utils"
)

type IUserRepository interface {
	Upsert(ctx context.Context, user *domain.User) error
	GetById(ctx context.Context, id string) (*domain.User, error)
	Get(ctx context.Context) ([]*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userRepository struct {
	userList []*domain.User
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		userList: utils.GetUserStub(),
	}
}

func (r *userRepository) Upsert(ctx context.Context, user *domain.User) error {
	r.userList = append(r.userList, user)
	return nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (*domain.User, error) {
	for _, user := range r.userList {
		if user.Id == id {
			return user, nil
		}
	}

	fmt.Printf("userRepository.GetById INFO Not found user by given id: %s\n", id)

	return nil, nil
}

func (r *userRepository) Get(ctx context.Context) ([]*domain.User, error) {
	users := r.userList

	if users == nil {
		fmt.Printf("userRepository.Get INFO not found users on datasource\n")
		return make([]*domain.User, 0), nil
	}

	return r.userList, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	for _, user := range r.userList {
		if user.Email == email {
			return user, nil
		}
	}

	fmt.Printf("userRepository.GetByEmail INFO Not found user by given email: %s\n", email)
	return nil, nil
}
