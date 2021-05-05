package profile

import profileentity "github.com/almanalfaruq/alfarpos-backend/model/profile"

type ProfileService struct {
	repo profileRepo
}

func NewProfile(repo profileRepo) *ProfileService {
	return &ProfileService{
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
