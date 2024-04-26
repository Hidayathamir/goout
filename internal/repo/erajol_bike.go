package repo

import (
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/repo/cache"
	"github.com/Hidayathamir/goout/internal/repo/db"
)

type IErajolBike interface{}

type ErajolBike struct {
	cfg             config.Config
	pg              *db.Postgres
	cacheErajolBike cache.IErajolBike
}

var _ IErajolBike = &ErajolBike{}

func NewErajolBike(cfg config.Config, pg *db.Postgres, cacheErajolBike cache.IErajolBike) *ErajolBike {
	return &ErajolBike{
		cfg:             cfg,
		pg:              pg,
		cacheErajolBike: cacheErajolBike,
	}
}
