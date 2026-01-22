package category

import "gorm.io/gorm"

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db}
}

func (r *gormRepository) Create(c *Category) error {
	return r.db.Create(c).Error
}

func (r *gormRepository) FindByID(id uint) (*Category, error) {
	var category Category
	err := r.db.First(&category, id).Error
	return &category, err

}

func (r *gormRepository) FindAll() ([]*Category, error) {
	var categories []*Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *gormRepository) Update(c *Category) error {
	return r.db.Save(c).Error
}

func (r *gormRepository) Delete(id uint) error {
	return r.db.Delete(&Category{}, id).Error
}
