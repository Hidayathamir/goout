package main

import (
	"path/filepath"

	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/repo/db"
	"github.com/Hidayathamir/goout/pkg/trace"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// gormPlayground is a function for initializing a GORM playground environment.
func gormPlayground(fn func(pg *gorm.DB)) {
	cfg, err := config.Load(filepath.Join("..", "..", "..", "..", "config.yml"))
	fatalIfErr(err)

	gormConfig := &gorm.Config{
		DryRun: true,
		Logger: logger.Default.LogMode(logger.Info),
	}

	pg, err := db.NewPostgres(cfg, db.WithGormConfig(gormConfig))
	fatalIfErr(err)

	fn(pg.DB)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
