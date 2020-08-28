package business

import (
	"net/http"

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

func validateBeer(beer models.Beer) *models.ApiError {
	if beer.ID != 0 {
		return &models.ApiError{Message: "Field ID not allowed", Status: http.StatusBadRequest}
	}
	if beer.Name == "" {
		return &models.ApiError{Message: "Field Name is mandatory", Status: http.StatusBadRequest}
	}
	return nil
}
