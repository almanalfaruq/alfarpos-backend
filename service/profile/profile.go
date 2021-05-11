package profile

import (
	profileentity "github.com/almanalfaruq/alfarpos-backend/model/profile"
	"github.com/almanalfaruq/alfarpos-backend/util"
)

type ProfileService struct {
	cfg  *util.Config
	repo profileRepo
}

func NewProfile(cfg *util.Config, repo profileRepo) *ProfileService {
	return &ProfileService{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *ProfileService) New(data profileentity.Profile) (profileentity.Profile, error) {
	return s.repo.New(data)
}

func (s *ProfileService) GetByID(id int64) (profileentity.Profile, error) {
	return s.repo.FindByID(id)
}

func (s *ProfileService) Update(data profileentity.Profile) (profileentity.Profile, error) {
	return s.repo.Update(data)
}

func (s *ProfileService) GetShopProfile() profileentity.Profile {
	return profileentity.Profile{
		Name:            s.cfg.ShopProfile.Name,
		Address:         s.cfg.ShopProfile.Address,
		Phone:           s.cfg.ShopProfile.Phone,
		ThankyouMessage: s.cfg.ShopProfile.ThankyouMessage,
		FootNote:        s.cfg.ShopProfile.FootNote,
	}
}
