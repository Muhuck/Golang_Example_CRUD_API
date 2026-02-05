package product

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(c *Product) error {
	return s.repo.Create(c)
}

func (s *Service) GetAll(filter ProductFilter) ([]*Product, error) {
	return s.repo.FindAll(filter)
}

func (s *Service) GetByID(id uint) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(c *Product) error {
	return s.repo.Update(c)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
