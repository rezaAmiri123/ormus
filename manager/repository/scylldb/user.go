package scylldb

import (
	"fmt"

	"github.com/rezaAmiri123/ormus/manager/entity"
)

func (a StorageAdapter) DoesUserExistsByEmail(email string) (bool, error) {
	fmt.Println(email)

	panic("implement me")
}

func (a StorageAdapter) Register(u entity.User) (*entity.User, error) {
	fmt.Println(u)

	panic("implement me")
}

func (a StorageAdapter) GetUserByEmail(email string) (*entity.User, error) {
	fmt.Println(email)

	panic("implement me")
}
