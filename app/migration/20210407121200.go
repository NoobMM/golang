package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrate20210407121200Up = []string{
	`CREATE TABLE "wallets" (
		id		uuid NOT NULL DEFAULT uuid_generate_v4(),
		name	varchar	NOT NULL,
		balance	bigint	NOT NULL,
		PRIMARY KEY (id)
	)`,
}

var migrate20210407121200Down = []string{
	`DROP TABLE "wallets"`,
}

func migrate20210407121200() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20210107121200",
		Migrate: func(db *gorm.DB) error {
			for _, sql := range migrate20210407121200Up {
				err := db.Exec(sql).Error
				if err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(db *gorm.DB) error {
			for _, sql := range migrate20210407121200Down {
				err := db.Exec(sql).Error
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}
