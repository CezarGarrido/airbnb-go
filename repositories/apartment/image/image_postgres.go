package image

import (
	"context"
	"database/sql"

	entities "github.com/CezarGarrido/airbnb-go/entities"
)

// NewPgApartmentImageRepo returns implement of user repository interface
func NewPgApartmentImageRepo(db *sql.DB, tx *sql.Tx) ApartmentImageRepo {
	return &posgresApartmentImageRepo{db: db, tx: tx}
}

type posgresApartmentImageRepo struct {
	db *sql.DB
	tx *sql.Tx
}

func (pgUserRepo *posgresApartmentImageRepo) CreateWithTx(ctx context.Context, image *entities.ApartmentImage) (int64, error) {

	query := `INSERT INTO apartment_images (uuid, apartment_id, file_name, file_size, url, created_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id`

	stmt, err := pgUserRepo.tx.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	var ID int64

	err = stmt.QueryRowContext(ctx,
		image.UUID,
		image.ApartmentID,
		image.FileName,
		image.FileSize,
		image.URL,
		image.CreatedAt,
	).Scan(&ID)

	if err != nil {
		return -1, err
	}

	return ID, nil
}

func (pgUserRepo *posgresApartmentImageRepo) FindByApartmenID(ctx context.Context, apartmentID int64) ([]*entities.ApartmentImage, error) {
	return pgUserRepo.fetch(ctx, "SELECT id, uuid, apartment_id, file_name, file_size, url, created_at, updated_at FROM apartment_images WHERE apartment_id=$1", apartmentID)
}

func (pgUserRepo *posgresApartmentImageRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.ApartmentImage, error) {
	rows, err := pgUserRepo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*entities.ApartmentImage, 0)
	for rows.Next() {
		apartment := new(entities.ApartmentImage)
		err := rows.Scan(
			&apartment.ID,
			&apartment.UUID,
			&apartment.ApartmentID,
			&apartment.FileName,
			&apartment.FileSize,
			&apartment.URL,
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
