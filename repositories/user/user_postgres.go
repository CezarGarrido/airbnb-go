package user

import (
	"context"
	"database/sql"

	entities "github.com/CezarGarrido/airbnb-go/entities"
)

// NewPgUserRepo returns implement of user repository interface
func NewPgUserRepo(db *sql.DB, tx *sql.Tx) UserRepo {
	return &posgresUserRepo{db: db, tx: tx}
}

type posgresUserRepo struct {
	db *sql.DB
	tx *sql.Tx
}

func (pgUserRepo *posgresUserRepo) Create(ctx context.Context, user *entities.User) (int64, error) {

	query := `INSERT INTO users (first_name, last_name, email, password, created_at) VALUES($1,$2,$3,$4,$5) RETURNING id`

	stmt, err := pgUserRepo.db.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	var ID int64

	err = stmt.QueryRowContext(ctx,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
	).Scan(&ID)

	if err != nil {
		return -1, err
	}

	return ID, nil
}
