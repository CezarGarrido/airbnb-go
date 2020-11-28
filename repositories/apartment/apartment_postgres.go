package apartment

import (
	"context"
	"database/sql"
	"errors"

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

func (pgApartmentRepo *posgresApartmentRepo) CreateWithTx(ctx context.Context, apartment *entities.Apartment) (int64, error) {
	query := `INSERT INTO apartments
	(uuid, user_id, "name", description, room_type, property_type, street, city, state, country, guests, bedrooms, beds, baths, price, status, created_at)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15, $16, $17)
	RETURNING id;`

	stmt, err := pgApartmentRepo.tx.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()
	var ID int64
	err = stmt.QueryRowContext(ctx,
		apartment.UUID,
		apartment.UserID,
		apartment.Name,
		apartment.Description,
		apartment.RoomType,
		apartment.PropertyType,
		apartment.Street,
		apartment.City,
		apartment.State,
		apartment.Country,
		apartment.Guests,
		apartment.Bedrooms,
		apartment.Beds,
		apartment.Baths,
		apartment.Price,
		apartment.Status,
		apartment.CreatedAt,
	).Scan(&ID)

	if err != nil {
		return -1, err
	}

	return ID, nil
}

func (pgApartmentRepo *posgresApartmentRepo) Update(ctx context.Context, apartment *entities.Apartment) (*entities.Apartment, error) {
	return nil, nil
}

func (pgApartmentRepo *posgresApartmentRepo) Delete(ctx context.Context, apartment *entities.Apartment) error {
	return nil
}

func (pgApartmentRepo *posgresApartmentRepo) FindAll(ctx context.Context) ([]*entities.Apartment, error) {
	query := "SELECT id, uuid, user_id, name, description, street, city, state, country, guests, bedrooms, beds, baths, likes, price, status, created_at, updated_at FROM apartments;"
	return pgApartmentRepo.fetch(ctx, query)
}

func (pgApartmentRepo *posgresApartmentRepo) FindByUUID(ctx context.Context, UUID string) (*entities.Apartment, error) {
	query := "SELECT id, uuid, user_id, name, description, street, city, state, country, guests, bedrooms, beds, baths, likes, price, status, created_at, updated_at FROM apartments WHERE uuid=$1;"
	rows, err := pgApartmentRepo.fetch(ctx, query, UUID)
	if err != nil {
		return nil, err
	}
	if len(rows) <= 0 {
		return nil, errors.New("Apartamento não encontrado")
	}
	return rows[0], nil
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
			&apartment.UUID,
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
			&apartment.CreatedAt,
			&apartment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, apartment)
	}
	return payload, nil
}
