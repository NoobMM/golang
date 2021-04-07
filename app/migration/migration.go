package migration

import (
	"github.com/NoobMM/golang/app/domain/interfaces/connectors"
	"github.com/go-gormigrate/gormigrate/v2"
)


func migrateEngine() *gormigrate.Gormigrate {
	db := connectors.ConnectPostgresDB(true)

	migrateOpt := gormigrate.DefaultOptions
	return gormigrate.New(db, migrateOpt, []*gormigrate.Migration{
		migrate20210407121200(),
	})

}

func Migrate() error {
	err := migrateEngine().Migrate()
	if err != nil {
		return err
	}
	return nil
}
