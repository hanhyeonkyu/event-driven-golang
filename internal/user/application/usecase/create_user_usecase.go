package usecase

import (
	"context"
	"fmt"

	"github.com/hanhyeonkyu/event-driven-golang/internal/user/domain/entity"
	"github.com/hanhyeonkyu/event-driven-golang/internal/user/domain/event"
	"github.com/hanhyeonkyu/event-driven-golang/internal/user/domain/queue"
)

type CreateUserUseCase struct {
	publisher queue.Publisher
}

func NewCreateUserUseCase(publisher queue.Publisher) *CreateUserUseCase {
	return &CreateUserUseCase{
		publisher: publisher,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, name, email string) error {
	fmt.Println("--- CreateUserUseCase ---")
	user, err := entity.NewUserEntity(name, email)
	if err != nil {
		return err
	}
	event := event.UserRegisteredEvent{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}
	err = u.publisher.Publish(ctx, event)
	if err != nil {
		return err
	}
	return nil
}
