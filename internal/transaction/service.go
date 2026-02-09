package transaction

import (
	"crud-api/internal/product"
	"time"
)

type Service struct {
	repo        Repository
	productRepo product.Repository
}

func NewService(repo Repository, productRepo product.Repository) *Service {
	return &Service{repo: repo, productRepo: productRepo}
}

func (s *Service) Checkout(items []CheckoutItem) (*Transaction, error) {
	var details []TransactionDetail
	var totalAmount int

	for _, item := range items {
		product, err := s.productRepo.FindByID(uint(item.ProductID))
		if err != nil {
			return nil, err
		}

		subtotal := product.Price * item.Quantity

		details = append(details, TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
		totalAmount += subtotal
	}

	transaction := &Transaction{
		TotalAmount: totalAmount,
		CreatedAt:   time.Now(),
		Details:     details,
	}

	if err := s.repo.Create(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
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
