package usecase

import (
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/extapi"
	"github.com/Hidayathamir/goout/internal/repo"
	"github.com/Hidayathamir/goout/internal/repo/cache"
	"github.com/Hidayathamir/goout/internal/repo/db"
)

// InitUsecaseErajolBike initializes the ErajolBike usecase.
func InitUsecaseErajolBike(cfg *config.Config, pg *db.Postgres, redis *cache.Redis, extapiGocheck extapi.IGocheck) *ErajolBike {
	cacheErajolBike := cache.NewErajolBike(cfg, redis)
	repoErajolBike := repo.NewErajolBike(cfg, pg, cacheErajolBike)
	usecaseErajolBike := NewErajolBike(cfg, pg.TxManager, repoErajolBike, extapiGocheck)
	return usecaseErajolBike
}
