package http

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CezarGarrido/airbnb-go/driver"
	"github.com/CezarGarrido/airbnb-go/entities"
	"github.com/CezarGarrido/airbnb-go/lib/render"
	apartmentRepo "github.com/CezarGarrido/airbnb-go/repositories/apartment"
	apartmentImageRepo "github.com/CezarGarrido/airbnb-go/repositories/apartment/image"
	"github.com/CezarGarrido/multi"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// NewApartment :
func NewApartment(db *driver.DB) *Apartment {
	return &Apartment{
		db:                 db.SQL,
		ApartmentRepo:      apartmentRepo.NewPgApartmentRepo(db.SQL, nil),
		ApartmentImageRepo: apartmentImageRepo.NewPgApartmentImageRepo(db.SQL, nil),
	}
}

// Apartment :
type Apartment struct {
	db                 *sql.DB
	ApartmentRepo      apartmentRepo.ApartmentRepo
	ApartmentImageRepo apartmentImageRepo.ApartmentImageRepo
}

// FindAll :
func (ap *Apartment) FindAll(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	apartments, err := ap.ApartmentRepo.FindAll(ctx)
	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	for _, apartment := range apartments {
		apartment.Images, _ = ap.ApartmentImageRepo.FindByApartmenID(ctx, apartment.ID)
	}

	render.HTTPWriteJSON(w, http.StatusOK, apartments)
}

// FindByUUID :
func (ap *Apartment) FindByUUID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	apartment, err := ap.ApartmentRepo.FindByUUID(ctx, mux.Vars(r)["uuid"])
	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	apartment.Images, _ = ap.ApartmentImageRepo.FindByApartmenID(ctx, apartment.ID)

	render.HTTPWriteJSON(w, http.StatusOK, apartment)
}

// CreateFormData :
func (ap *Apartment) CreateFormData(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	err := r.ParseMultipartForm(200000)
	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	formdata := r.MultipartForm
	files := formdata.File["files"]

	var Apartment entities.Apartment

	Apartment.Name = r.PostFormValue("name")

	Apartment.Description = r.PostFormValue("description")
	Apartment.RoomType = r.PostFormValue("room_type")
	Apartment.PropertyType = r.PostFormValue("property_type")
	Apartment.Street = r.PostFormValue("street")
	Apartment.City = r.PostFormValue("city")
	Apartment.State = r.PostFormValue("state")
	Apartment.Country = r.PostFormValue("country")
	Apartment.Guests, _ = strconv.Atoi(r.PostFormValue("guests"))
	Apartment.Bedrooms, _ = strconv.Atoi(r.PostFormValue("bedrooms"))
	Apartment.Beds, _ = strconv.Atoi(r.PostFormValue("beds"))
	Apartment.Baths, _ = strconv.Atoi(r.PostFormValue("baths"))
	Apartment.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 64)
	Apartment.UUID = uuid.New().String()
	Apartment.UserID = 1
	Apartment.Status = "available"

	values, err := multi.New().Run("Inserir apartamento e imagens", ap.db, func(tx *sql.Tx) (multi.Result, error) {

		repoApartmentTx := apartmentRepo.NewPgApartmentRepo(nil, tx)
		repoApartmentImageTx := apartmentImageRepo.NewPgApartmentImageRepo(nil, tx)

		ApartmentImages := make([]*entities.ApartmentImage, 0)

		newID, err := repoApartmentTx.CreateWithTx(ctx, &Apartment)
		if err != nil {
			return nil, err
		}
		Apartment.ID = newID

		for i := range files {

			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				return nil, err
			}
			UUID := uuid.New().String()

			fileName := UUID + "-" + files[i].Filename

			ApartmentImage := entities.ApartmentImage{
				ApartmentID: Apartment.ID,
				FileName:    fileName,
				FileSize:    files[i].Size,
				URL:         "http://localhost:3000/storage/locations/images/" + fileName,
			}

			ApartmentImage.UUID = UUID
			ApartmentImage.CreatedAt = time.Now()

			out, err := os.Create("./storage/locations/images/" + fileName)
			defer out.Close()
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(out, file)
			if err != nil {
				return nil, err
			}
			imageID, err := repoApartmentImageTx.CreateWithTx(ctx, &ApartmentImage)
			if err != nil {
				return nil, err
			}
			ApartmentImage.ID = imageID
			ApartmentImages = append(ApartmentImages, &ApartmentImage)
		}

		Apartment.Images = ApartmentImages

		result := multi.Result{
			"apartment": Apartment,
		}

		return result, nil
	})
	if err != nil {
		log.Println(err.Error())
		render.HTTPWriteJSON(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	m, _ := values.(multi.Result)

	render.HTTPWriteJSON(w, http.StatusOK, m["apartment"])

}
