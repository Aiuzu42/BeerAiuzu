package app

import (
	"github.com/aiuzu42/BeerAiuzu/business"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func Start() {
	business.SelectImplementation(1)
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
