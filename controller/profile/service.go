package profile

import profileentity "github.com/almanalfaruq/alfarpos-backend/model/profile"

type profileService interface {
	New(data profileentity.Profile) (profileentity.Profile, error)
	GetByID(id int64) (profileentity.Profile, error)
	Update(data profileentity.Profile) (profileentity.Profile, error)
	GetShopProfile() profileentity.Profile
}
