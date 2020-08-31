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

func (m *Memory) BeerDetails(id int) (models.Beer, *models.ApiError) {
	for _, b := range m.beers {
		if b.ID == id {
			return b, nil
		}
	}
	return models.Beer{}, nil
}

func (m *Memory) DeleteBeer(id int) (bool, *models.ApiError) {
	for i, _ := range m.beers {
		if m.beers[i].ID == id {
			copy(m.beers[i:], m.beers[i+1:])
			m.beers = m.beers[:len(m.beers)-1]
			return true, nil
		}
	}
	return false, nil
}
