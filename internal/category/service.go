package category

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(c *Category) error {
	return s.repo.Create(c)
}

func (s *Service) GetAll() ([]*Category, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uint) (*Category, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(c *Category) error {
	return s.repo.Update(c)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
