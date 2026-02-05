package transaction

type Repository interface {
	FindByID(id uint) (*Transaction, error)
	FindAll() ([]*Transaction, error)
	Create(transaction *Transaction) error
	Update(transaction *Transaction) error
	Delete(id uint) error
}
