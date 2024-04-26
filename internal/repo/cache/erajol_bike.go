package cache

import "github.com/Hidayathamir/goout/internal/config"

type IErajolBike interface{}

type ErajolBike struct {
	cfg   config.Config
	redis *Redis
}

var _ IErajolBike = &ErajolBike{}

func NewErajolBike(cfg config.Config, redis *Redis) *ErajolBike {
	return &ErajolBike{
		cfg:   cfg,
		redis: redis,
	}
}
