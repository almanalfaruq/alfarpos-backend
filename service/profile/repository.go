package profile

import profileentity "github.com/almanalfaruq/alfarpos-backend/model/profile"

type profileRepo interface {
	New(data profileentity.Profile) (profileentity.Profile, error)
	FindByID(id int64) (profileentity.Profile, error)
	Update(data profileentity.Profile) (profileentity.Profile, error)
}
