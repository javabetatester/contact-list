package groups

type Service struct {
	Repository Repository
}

func (s *Service) Create(group Group) error {
	return s.Repository.Create(group)
}

func (s *Service) List() ([]Group, error) {
	return s.Repository.List()
}

func (s *Service) Update(group Group) error {
	return s.Repository.Update(group)
}

func (s *Service) Delete(id int) error {
	return s.Repository.Delete(id)
}

func (s *Service) GetById(id int) (Group, error) {
	return s.Repository.GetById(id)
}

func (s *Service) GetByName(name string) (Group, error) {
	return s.Repository.GetByName(name)
}

func (s *Service) AddContactToGroup(groupID int, contactID int) error {
	return s.Repository.AddContactToGroup(groupID, contactID)
}