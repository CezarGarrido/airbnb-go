package image

import (
	"context"

	entities "github.com/CezarGarrido/airbnb-go/entities"
)

// ApartmentImageRepo :
type ApartmentImageRepo interface {
	CreateWithTx(ctx context.Context, image *entities.ApartmentImage) (int64, error)
	FindByApartmenID(ctx context.Context, apartmentID int64) ([]*entities.ApartmentImage, error)
}
