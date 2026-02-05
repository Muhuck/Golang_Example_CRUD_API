package product

import "gorm.io/gorm"

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db}
}

func (r *gormRepository) Create(c *Product) error {
	return r.db.Create(c).Error
}

func (r *gormRepository) FindByID(id uint) (*Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	return &product, err

}

func (r *gormRepository) FindAll(filter ProductFilter) ([]*Product, error) {
	var products []*Product
	query := r.db.Preload("Category")

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	err := query.Find(&products).Error
	return products, err
}

func (r *gormRepository) Update(c *Product) error {
	return r.db.Save(c).Error
}

func (r *gormRepository) Delete(id uint) error {
	return r.db.Delete(&Product{}, id).Error
}
