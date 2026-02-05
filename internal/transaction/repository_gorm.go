package transaction

import "gorm.io/gorm"

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db}
}

func (r *gormRepository) Create(c *Transaction) error {
	return r.db.Create(c).Error
}

func (r *gormRepository) FindByID(id uint) (*Transaction, error) {
	var transaction Transaction
	err := r.db.First(&transaction, id).Error
	return &transaction, err

}

func (r *gormRepository) FindAll() ([]*Transaction, error) {
	var transactions []*Transaction
	err := r.db.Preload("Transaction").Find(&transactions).Error
	return transactions, err
}

func (r *gormRepository) Update(c *Transaction) error {
	return r.db.Save(c).Error
}

func (r *gormRepository) Delete(id uint) error {
	return r.db.Delete(&Transaction{}, id).Error
}
