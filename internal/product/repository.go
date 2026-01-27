package product

type Repository interface {
	FindByID(id uint) (*Product, error)
	FindAll() ([]*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id uint) error
}
