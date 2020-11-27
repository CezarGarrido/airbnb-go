package user

import (
	"context"

	entities "github.com/CezarGarrido/airbnb-go/entities"
)

// UserRepo :
type UserRepo interface {
	Create(ctx context.Context, user *entities.User) (int64, error)
}
