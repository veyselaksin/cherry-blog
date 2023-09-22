
package service

import (
	Dao "cherry-blog/db/dao/blog"

)

type Service interface {
	// Define service methods
}

type service struct {
	data Dao.DataAccess
}

func NewService(data Dao.DataAccess) Service {
	return service{data: data}
}

