// Package main -.
package main

import (
	"github.com/Hidayathamir/goout/internal/config"
	"github.com/Hidayathamir/goout/internal/repo/db"
	"github.com/Hidayathamir/goout/internal/repo/db/migration/migrate"
	"github.com/Hidayathamir/goout/pkg/trace"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.Load("./config.yml")
	fatalIfErr(err)

	pg, err := db.NewPostgres(*cfg)
	fatalIfErr(err)

	err = migrate.Up(pg.DB)
	fatalIfErr(err)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
