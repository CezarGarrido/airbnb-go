package apartment

import (
	"context"
	"database/sql"

	entities "github.com/CezarGarrido/airbnb-go/entities"
)

// NewPgApartmentRepo returns implement of Apartment repository interface
func NewPgApartmentRepo(db *sql.DB, tx *sql.Tx) ApartmentRepo {
	return &posgresApartmentRepo{db: db, tx: tx}
}

type posgresApartmentRepo struct {
	db *sql.DB
	tx *sql.Tx
}

func (pgApartmentRepo *posgresApartmentRepo) Create(ctx context.Context, apartment *entities.Apartment) (int64, error) {
	return -1, nil
}

func (pgApartmentRepo *posgresApartmentRepo) Update(ctx context.Context, apartment *entities.Apartment) (*entities.Apartment, error) {
	return nil, nil
}

func (pgApartmentRepo *posgresApartmentRepo) Delete(ctx context.Context, apartment *entities.Apartment) error {
	return nil
}

func (pgApartmentRepo *posgresApartmentRepo) FindAll(ctx context.Context) ([]*entities.Apartment, error) {
	query := "SELECT id, user_id, name, description, street, city, state, country, guests, bedrooms, beds, baths, likes, price, status, created_at, updated_at FROM apartments;"
	return pgApartmentRepo.fetch(ctx, query)
}

// fetch
func (pgApartmentRepo *posgresApartmentRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.Apartment, error) {
	rows, err := pgApartmentRepo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*entities.Apartment, 0)
	for rows.Next() {
		apartment := new(entities.Apartment)
		err := rows.Scan(
			&apartment.ID,
			&apartment.UserID,
			&apartment.Name,
			&apartment.Description,
			&apartment.Street,
			&apartment.City,
			&apartment.State,
			&apartment.Country,
			&apartment.Guests,
			&apartment.Bedrooms,
			&apartment.Beds,
			&apartment.Baths,
			&apartment.Likes,
			&apartment.Price,
			&apartment.Status,
			&apartment.CreateadAt,
			&apartment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, apartment)
	}
	return payload, nil
}
