
package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s application_name='%s' sslmode=disable timezone=Europe/Istanbul",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		"5432",
		"cherry-blog",
	)

	connection, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil
	}
	return connection
}
