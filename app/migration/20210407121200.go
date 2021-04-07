package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrate20210407121200Up = []string{
	`CRAETE EXTENSION "uuid_ossp"`,
	`CRAETE TABLE "wallets" (
		id		uuid PRIMARY KEY NOT NULL DEFAULT generate_uuid_v4(),
		name	varchar	NOT NULL,
		balance	bigint	NOT NULL
	)`,
}

var migrate20210407121200Down = []string{
	`DROP TABLE "wallets"`,
	`DROP EXTENSION "uuid_ossp"`,
}

func migrate20210407121200() *gormigrate.Migration {
	return &gormigrate.Migration{
	ID: "20210107121200",
		Migrate: func(db *gorm.DB) error {
			for _, sql := range migrate20210407121200Up {
				err := db.Exec(sql).Error
				if err != nil{
					return err
				}
			}
			return nil
			},
			Rollback: func(db *gorm.DB) error {
				for _, sql := range migrate20210407121200Down {
					err := db.Exec(sql).Error
					if err != nil{
						return err
					}
				}
				return nil
			},
		}
}