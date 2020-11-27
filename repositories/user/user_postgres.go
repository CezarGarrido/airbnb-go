package user

import (
	"context"
	"database/sql"
	"errors"

	entities "github.com/CezarGarrido/airbnb-go/entities"
	"github.com/lib/pq"
)

// ErrorDuplicateEmail :
var ErrorDuplicateEmail = errors.New("Email já cadastrado.")

// NewPgUserRepo returns implement of user repository interface
func NewPgUserRepo(db *sql.DB, tx *sql.Tx) UserRepo {
	return &posgresUserRepo{db: db, tx: tx}
}

type posgresUserRepo struct {
	db *sql.DB
	tx *sql.Tx
}

func (pgUserRepo *posgresUserRepo) Create(ctx context.Context, user *entities.User) (int64, error) {

	query := `INSERT INTO users (uuid, first_name, last_name, email, password, created_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id`

	stmt, err := pgUserRepo.db.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	var ID int64

	err = stmt.QueryRowContext(ctx,
		user.UUID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
	).Scan(&ID)

	pqErr := err.(*pq.Error)

	if pqErr.Code == pq.ErrorCode("23505") {
		return -1, ErrorDuplicateEmail
	}

	if err != nil {
		return -1, err
	}

	return ID, nil
}
