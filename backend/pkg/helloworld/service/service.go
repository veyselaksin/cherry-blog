
package service

import (
	Dao "cherry-blog/db/dao/helloworld"

)

type Service interface {
	// Define service methods
	Helloworld() (string, int, error)
}

type service struct {
	data Dao.DataAccess
}

func NewService(data Dao.DataAccess) Service {
	return service{data: data}
}

func (s service) Helloworld() (string, int, error) {
	helloWorld := "Hello World!"
	return helloWorld, 200, nil
}

