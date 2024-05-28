package pguser

import (
	"fmt"

	"github.com/sazonovItas/proxy-manager/pkg/postgresdb"
)

type userRepository struct {
	db        *postgresdb.DB
	tableName string
}

func New(db *postgresdb.DB, tableName string) *userRepository {
	return &userRepository{
		db:        db,
		tableName: tableName,
	}
}

func (ur *userRepository) table(query string) string {
	return fmt.Sprintf(query, ur.tableName)
}
