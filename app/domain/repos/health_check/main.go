package healthcheckrepo

import (
	"context"
	"gorm.io/gorm"
)

// Repo is an interface for HealthCheck repository
type Repo interface {
	CheckDatabaseReadiness(ctx context.Context) error
}

type repo struct {
	DB *gorm.DB
}

// New is a constructor method of Repo
func New(db *gorm.DB) Repo {
	return &repo{
		DB: db,
	}
}
