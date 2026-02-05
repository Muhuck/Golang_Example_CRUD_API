package product

type Repository interface {
	FindByID(id uint) (*Product, error)
	FindAll(filter ProductFilter) ([]*Product, error)
	Create(product *Product) error
	Update(product *Product) error
	Delete(id uint) error
}
