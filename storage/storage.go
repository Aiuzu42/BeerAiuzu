package storage

import "github.com/aiuzu42/BeerAiuzu/models"

type Storage interface {
	ListAllBeers() ([]models.Beer, *models.ApiError)
	AddBeer(beer models.Beer) *models.ApiError
	BeerDetails(id int) (models.Beer, *models.ApiError)
	DeleteBeer(id int) (bool, *models.ApiError)
}
