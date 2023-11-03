package services

import (
	"encoding/json"
	"github.com/andresPirona/go-singleflight-ap/domain/entity"
	"github.com/andresPirona/go-singleflight-ap/domain/repository"
	"log"
	"os"
)

type implementationCountry struct{}

func (i implementationCountry) GetAll() []entity.Country {

	var staticCountryData []entity.Country

	// Abre el archivo JSON
	file, err := os.Open("resources/countries.json")
	if err != nil {
		log.Fatalf("Error al abrir el archivo JSON: %v", err)
	}
	defer file.Close()

	// Decodifica el archivo JSON en la estructura de datos
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&staticCountryData); err != nil {
		log.Fatalf("Error al decodificar el archivo JSON: %v", err)
	}

	return staticCountryData

}

func NewCountryImplementation() repository.CountryRepository {
	return &implementationCountry{}
}
