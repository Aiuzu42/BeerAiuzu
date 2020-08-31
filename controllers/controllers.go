package controllers

import (
	"net/http"

	"github.com/aiuzu42/BeerAiuzu/business"
	"github.com/aiuzu42/BeerAiuzu/models"
	"github.com/gin-gonic/gin"
)

func GetAllBeers(c *gin.Context) {
	result, err := business.ListBeers()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if len(result) < 1 {
		c.Writer.WriteHeader(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func PostBeer(c *gin.Context) {
	var beerRequest models.Beer
	if err := c.ShouldBindJSON(&beerRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.ApiError{Message: "Bad Request", Status: http.StatusBadRequest})
		return
	}
	apiErr := business.AddBeer(beerRequest)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)
}

func GetBeer(c *gin.Context) {
	id := c.Param("ID")
	b, apiErr := business.BeerDetails(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, b)
}

func DeleteBeer(c *gin.Context) {
	id := c.Param("ID")
	apiErr := business.DeleteBeer(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}
