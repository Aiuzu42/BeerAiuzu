package business

import (
	"net/http"
	"strconv"

	"github.com/aiuzu42/BeerAiuzu/models"
	"github.com/aiuzu42/BeerAiuzu/storage"
)

var (
	impl storage.Storage = &storage.Memory{}
)

func SelectImplementation(code int) {
	if code == 1 {
		impl = &storage.Memory{}
	}
}

func ListBeers() ([]models.Beer, *models.ApiError) {
	return impl.ListAllBeers()
}

func AddBeer(beer models.Beer) *models.ApiError {
	if err := validateBeer(beer); err != nil {
		return err
	}
	return impl.AddBeer(beer)
}

func BeerDetails(id string) (models.Beer, *models.ApiError) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Beer{}, &models.ApiError{Message: "Wrong parameters", Status: http.StatusBadRequest}
	}
	b, errImpl := impl.BeerDetails(idInt)
	if errImpl != nil {
		return models.Beer{}, errImpl
	}
	if b.ID == 0 {
		return models.Beer{}, &models.ApiError{Message: "Not found", Status: http.StatusNotFound}
	} else {
		return b, nil
	}
}

func DeleteBeer(id string) *models.ApiError {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return &models.ApiError{Message: "Wrong parameters", Status: http.StatusBadRequest}
	}
	b, errImpl := impl.DeleteBeer(idInt)
	if err != nil {
		return errImpl
	}
	if !b {
		return &models.ApiError{Message: "Not found", Status: http.StatusNotFound}
	} else {
		return nil
	}
}

func validateBeer(beer models.Beer) *models.ApiError {
	if beer.ID != 0 {
		return &models.ApiError{Message: "Field ID not allowed", Status: http.StatusBadRequest}
	}
	if beer.Name == "" {
		return &models.ApiError{Message: "Field Name is mandatory", Status: http.StatusBadRequest}
	}
	return nil
}
