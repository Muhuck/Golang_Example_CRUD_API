package transaction

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(c *Transaction) error {
	return s.repo.Create(c)
}

func (s *Service) GetAll() ([]*Transaction, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uint) (*Transaction, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(c *Transaction) error {
	return s.repo.Update(c)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
