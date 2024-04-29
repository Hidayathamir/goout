package repo

import (
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/repo/cache"
	"github.com/Hidayathamir/goout/internal/repo/db"
)

// IErajolBike defines the interface for the ErajolBike repository.
type IErajolBike interface{}

// ErajolBike represents the implementation of the ErajolBike repository.
type ErajolBike struct {
	cfg             *config.Config
	pg              *db.Postgres
	cacheErajolBike cache.IErajolBike
}

var _ IErajolBike = &ErajolBike{}

// NewErajolBike creates a new instance of the ErajolBike repository.
func NewErajolBike(cfg *config.Config, pg *db.Postgres, cacheErajolBike cache.IErajolBike) *ErajolBike {
	return &ErajolBike{
		cfg:             cfg,
		pg:              pg,
		cacheErajolBike: cacheErajolBike,
	}
}
