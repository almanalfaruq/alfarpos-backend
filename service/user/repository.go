package user

import (
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
)

type userRepositoryIface interface {
	FindAll() ([]userentity.User, error)
	FindById(id int64) (userentity.User, error)
	FindByUsername(username string) (userentity.User, error)
	FindByUsernameForLogin(username string) (userentity.User, error)
	New(userData userentity.User) (userentity.User, error)
	Update(userData userentity.User) (userentity.User, error)
	Delete(id int64) (userentity.User, error)
}
