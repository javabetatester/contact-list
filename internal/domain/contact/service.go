package contact

type Service struct {
	Repository Repository
}

func (s *Service) Create(contact Contact) error {
	return s.Repository.Create(contact)
}

func (s *Service) List() ([]Contact, error) {
	return s.Repository.List()
}

func (s *Service) Update(contact Contact) error {
	return s.Repository.Update(contact)
}

func (s *Service) Delete(contact Contact) error {
	return s.Repository.Delete(contact)
}

func (s *Service) GetById(id int) (Contact, error) {
	return s.Repository.GetById(id)
}

func (s *Service) GetByName(name string) (Contact, error) {
	return s.Repository.GetByName(name)
}

func (s *Service) GetByEmail(email string) (Contact, error) {
	return s.Repository.GetByEmail(email)
}

func (s *Service) GetByPhone(phone string) (Contact, error) {
	return s.Repository.GetByPhone(phone)
}
