package repologgermw

import (
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/repo"
)

// ErajolBike represents the implementation of the ErajolBike logger middleware.
type ErajolBike struct {
	cfg  *config.Config
	next repo.IErajolBike
}

var _ repo.IErajolBike = &ErajolBike{}

// NewErajolBike creates a new instane of ErajolBike logger middleware.
func NewErajolBike(cfg *config.Config, next repo.IErajolBike) *ErajolBike {
	return &ErajolBike{
		cfg:  cfg,
		next: next,
	}
}
