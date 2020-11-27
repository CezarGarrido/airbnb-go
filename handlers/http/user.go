package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/CezarGarrido/airbnb-go/driver"
	"github.com/CezarGarrido/airbnb-go/entities"
	"github.com/CezarGarrido/airbnb-go/lib/render"
	userRepo "github.com/CezarGarrido/airbnb-go/repositories/user"
	"github.com/google/uuid"
)

// NewUser :
func NewUser(db *driver.DB) *User {
	return &User{
		UserRepo: userRepo.NewPgUserRepo(db.SQL, nil),
	}
}

type User struct {
	UserRepo userRepo.UserRepo
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var user entities.User

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = user.Validate()
	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	user.CreatedAt = time.Now()
	user.UUID = uuid.New().String()
	newID, err := u.UserRepo.Create(ctx, &user)

	if err != nil && err == userRepo.ErrorDuplicateEmail {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Já existe uma conta para esse email!")
		return
	}

	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	user.ID = newID

	render.HTTPWriteJSON(w, http.StatusOK, user)

}
