package migrations

import (
	"embed"

	"github.com/RecoLabs/go-clickhouse/chmigrate"
)

var Migrations = chmigrate.NewMigrations()

//go:embed *.sql
var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
