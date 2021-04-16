package healthcheckrepo

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

func (repo *repo) CheckDatabaseReadiness(ctx context.Context) error {
	errLocationMsg := "[healthcheckrepo/check database readiness] %s"
	result := repo.DB.Exec("SELECT 1=1")
	if result.Error != nil {
		return errors.Wrap(result.Error, fmt.Sprintf(errLocationMsg, "database is not ready to serve the request"))
	}
	return nil
}
