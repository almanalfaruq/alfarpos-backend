package profile

import (
	profileentity "github.com/almanalfaruq/alfarpos-backend/model/profile"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

type ProfileRepo struct {
	db util.DBIface
}

func NewProfile(db util.DBIface) *ProfileRepo {
	return &ProfileRepo{
		db: db,
	}
}

func (r *ProfileRepo) New(data profileentity.Profile) (profileentity.Profile, error) {
	db := r.db.GetDb()
	return data, db.Create(&data).Error
}

func (r *ProfileRepo) FindByID(id int64) (profileentity.Profile, error) {
	var resp profileentity.Profile
	db := r.db.GetDb()
	return resp, db.Where("profiles.id = ?", id).Error
}

func (r *ProfileRepo) Update(data profileentity.Profile) (profileentity.Profile, error) {
	var oldProfile profileentity.Profile
	db := r.db.GetDb()
	err := db.Where("profiles.id = ?", data.ID).First(&oldProfile).Error
	if err != nil {
		return profileentity.Profile{}, err
	}
	return data, db.Save(&data).Error
}
