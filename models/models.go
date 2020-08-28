package models

type Beer struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Size         int      `json:"size"`
	Presentation string   `json:"presentation"`
	Alcohol      float32  `json:"alcohol"`
	Country      string   `json:"country"`
	Brewery      string   `json:"brewery"`
	Brand        string   `json:"brand"`
	Ingredients  []string `json:"ingredients"`
	Edition      string   `json:"edition"`
	Year         int      `json:"year"`
}

type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
