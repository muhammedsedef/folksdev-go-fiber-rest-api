package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang-gocb-couchbase/configuration"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/application/repository"
	"golang-gocb-couchbase/internal/folksdev-fiber-rest-api/domain"
)

type ICommandHandler interface {
	Save(ctx context.Context, command Command) error
}

type commandHandler struct {
	userRepository repository.IUserRepository
}

func NewCommandHandler(userRepository repository.IUserRepository) ICommandHandler {
	return &commandHandler{userRepository: userRepository}
}

func (h *commandHandler) Save(ctx context.Context, command Command) error {
	user, err := h.userRepository.GetByEmail(ctx, command.Email)

	if err != nil {
		return err
	}

	if user != nil {
		return errors.New(fmt.Sprintf("User Already Exist for given email: %s", command.Email))
	}

	if err := h.userRepository.Upsert(ctx, h.BuildEntity(command), configuration.PersistenceDurationInDaysForIntermediateStates); err != nil {
		return err
	}

	return nil
}

func (h *commandHandler) BuildEntity(command Command) *domain.User {
	return &domain.User{
		Id:        uuid.NewString(),
		FirstName: command.FirstName,
		LastName:  command.LastName,
		Email:     command.Email,
		Password:  command.Password,
		Age:       command.Age,
	}
}
