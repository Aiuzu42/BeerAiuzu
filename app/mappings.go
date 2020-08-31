package app

import "github.com/aiuzu42/BeerAiuzu/controllers"

func mapUrls() {
	router.GET("/beer", controllers.GetAllBeers)
	router.POST("/beer", controllers.PostBeer)
	router.GET("/beer/:ID", controllers.GetBeer)
	router.DELETE("/beer/:ID", controllers.DeleteBeer)
}
