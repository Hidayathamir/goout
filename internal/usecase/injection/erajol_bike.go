package injection

import (
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/extapi"
	"github.com/Hidayathamir/goout/internal/repo"
	"github.com/Hidayathamir/goout/internal/repo/cache"
	"github.com/Hidayathamir/goout/internal/repo/db"
	"github.com/Hidayathamir/goout/internal/repomw/repologgermw"
	"github.com/Hidayathamir/goout/internal/usecase"
	"github.com/Hidayathamir/goout/internal/usecasemw/usecaseloggermw"
)

// InitUsecaseErajolBike initializes the ErajolBike usecase.
func InitUsecaseErajolBike(cfg *config.Config, pg *db.Postgres, redis *cache.Redis, extapiGocheck extapi.IGocheck) usecase.IErajolBike {
	cacheErajolBike := cache.NewErajolBike(cfg, redis)

	var repoErajolBike repo.IErajolBike
	repoErajolBike = repo.NewErajolBike(cfg, pg, cacheErajolBike)
	repoErajolBike = repologgermw.NewErajolBike(cfg, repoErajolBike)

	var usecaseErajolBike usecase.IErajolBike
	usecaseErajolBike = usecase.NewErajolBike(cfg, pg.TxManager, repoErajolBike, extapiGocheck)
	usecaseErajolBike = usecaseloggermw.NewErajolBike(cfg, usecaseErajolBike)

	return usecaseErajolBike
}
