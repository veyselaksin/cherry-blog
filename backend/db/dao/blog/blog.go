
package blog

import (

	"gorm.io/gorm"
)

type DataAccess interface {
	// Define data access methods
}

type dataAccess struct {
	db *gorm.DB
}

func NewDataAccess(connection *gorm.DB) DataAccess {
	return dataAccess{db: connection}
}

	