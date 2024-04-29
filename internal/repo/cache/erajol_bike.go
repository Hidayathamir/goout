package cache

import "github.com/Hidayathamir/goout/internal/config"

// IErajolBike defines the interface for the ErajolBike cache.
type IErajolBike interface{}

// ErajolBike represents the implementation of the ErajolBike cache.
type ErajolBike struct {
	cfg   *config.Config
	redis *Redis
}

var _ IErajolBike = &ErajolBike{}

// NewErajolBike creates a new instance of the ErajolBike cache.
func NewErajolBike(cfg *config.Config, redis *Redis) *ErajolBike {
	return &ErajolBike{
		cfg:   cfg,
		redis: redis,
	}
}
