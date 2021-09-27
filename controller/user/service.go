package user

import (
	userentity "github.com/almanalfaruq/alfarpos-backend/model/user"
)

type userServiceIface interface {
	LoginUser(userData string) (userentity.UserResponse, error)
	NewUser(userData string) (userentity.User, error)
	UpdateUser(userData string) (userentity.User, error)
}
