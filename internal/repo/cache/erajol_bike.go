package cache

import "github.com/Hidayathamir/goout/internal/config"

// IErajolBike -.
type IErajolBike interface{}

// ErajolBike -.
type ErajolBike struct {
	cfg   config.Config
	redis *Redis
}

var _ IErajolBike = &ErajolBike{}

// NewErajolBike -.
func NewErajolBike(cfg config.Config, redis *Redis) *ErajolBike {
	return &ErajolBike{
		cfg:   cfg,
		redis: redis,
	}
}
