package query

import (
	"context"
	"errors"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/application/repository"
	"golang-project-layout-swagger/internal/folksdev-fiber-rest-api/domain"
)

type IUserQueryService interface {
	GetById(ctx context.Context, id string) (*domain.User, error)
	Get(ctx context.Context) ([]*domain.User, error)
}

type userQueryService struct {
	userRepository repository.IUserRepository
}

func NewUserQueryService(userRepository repository.IUserRepository) IUserQueryService {
	return &userQueryService{
		userRepository: userRepository,
	}
}

func (q *userQueryService) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := q.userRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("not found error")
	}

	return user, nil
}

func (q *userQueryService) Get(ctx context.Context) ([]*domain.User, error) {
	users, err := q.userRepository.Get(ctx)

	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, errors.New("not found users")
	}

	return users, nil
}
