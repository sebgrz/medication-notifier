package db

import (
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigration(address string) {
	m, err := migrate.New("file://migrations", address)
	if err != nil {
		panic(fmt.Sprintf("migration failed: %s", err))
	}
	m.Up()

}
