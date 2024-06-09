package start

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/islombay/noutbuk_seller/config"
	"github.com/islombay/noutbuk_seller/pkg/logs"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func Init(cfg config.DBConfig, log logs.LoggerInterface) {
	migration(cfg, log)
}

func migration(cfg config.DBConfig, log logs.LoggerInterface) {
	var dbURL string

	dbURL = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PWD"),
		cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	fmt.Println(dbURL)

	migrationsPath := fmt.Sprintf("file://%s", cfg.MigrationsPath)

	log.Debug("initializing migrations")

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Panic("could not initialize migration", logs.Error(err))
		os.Exit(1)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Panic("could not migrate up", logs.Error(err))
		os.Exit(1)
	}

	log.Info("database migrated up")
}
