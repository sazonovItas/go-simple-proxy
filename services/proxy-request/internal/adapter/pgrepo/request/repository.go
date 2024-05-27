package pgrequest

import (
	"fmt"

	"github.com/sazonovItas/proxy-manager/pkg/postgresdb"
)

type requestRepository struct {
	db        *postgresdb.DB
	tableName string
}

func New(tableName string, db *postgresdb.DB) *requestRepository {
	return &requestRepository{
		tableName: tableName,
		db:        db,
	}
}

func (rr *requestRepository) table(query string) string {
	return fmt.Sprintf(query, rr.tableName)
}
