package storage

import "github.com/aiuzu42/BeerAiuzu/models"

type Memory struct {
	beers   []models.Beer
	counter int
}

func (m *Memory) ListAllBeers() ([]models.Beer, *models.ApiError) {
	return m.beers, nil
}

func (m *Memory) AddBeer(beer models.Beer) *models.ApiError {
	m.counter = m.counter + 1
	beer.ID = m.counter
	m.beers = append(m.beers, beer)
	return nil
}
