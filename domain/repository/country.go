package repository

import "github.com/andresPirona/go-singleflight-ap/domain/entity"

type CountryRepository interface {
	GetAll() []entity.Country
}
