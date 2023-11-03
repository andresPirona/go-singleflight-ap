package main

import (
	"fmt"
	"github.com/andresPirona/go-singleflight-ap/domain/entity"
	"github.com/andresPirona/go-singleflight-ap/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
	"net/http"
	"time"
)

var group singleflight.Group

func GetCountryList(c *gin.Context) {

	fmt.Println(" ========= Simulación de una consulta costosa ========== ")
	// Simulación de una consulta costosa
	time.Sleep(5 * time.Second)

	countryRepo := services.NewCountryImplementation()
	countryList := countryRepo.GetAll()

	c.JSON(http.StatusOK, countryList)

}

func GetCountrySingleflightList(c *gin.Context) {
	result, err, _ := group.Do("countryList", func() (interface{}, error) {
		fmt.Println(" ========= Simulación de una consulta costosa ========== ")
		// Simulación de una consulta costosa
		time.Sleep(5 * time.Second)

		countryRepo := services.NewCountryImplementation()
		return countryRepo.GetAll(), nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos del país"})
		return
	}

	countryData, ok := result.([]entity.Country)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en el formato de los datos del país"})
		return
	}

	c.JSON(http.StatusOK, countryData)
}
