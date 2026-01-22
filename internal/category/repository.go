package category

type Repository interface {
	FindByID(id uint) (*Category, error)
	FindAll() ([]*Category, error)
	Create(category *Category) error
	Update(category *Category) error
	Delete(id uint) error
}
