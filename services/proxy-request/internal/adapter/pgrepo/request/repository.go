package pgrequest

import (
	"fmt"

	"github.com/sazonovItas/proxy-manager/pkg/postgresdb"
)

type RequestRepository struct {
	db        *postgresdb.DB
	tableName string
}

func New(tableName string, db *postgresdb.DB) *RequestRepository {
	return &RequestRepository{
		tableName: tableName,
		db:        db,
	}
}

func (rr *RequestRepository) table(query string) string {
	return fmt.Sprintf(query, rr.tableName)
}
