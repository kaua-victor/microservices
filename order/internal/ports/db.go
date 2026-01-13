package ports

import "github.com/kaua-victor/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
	Update(order *domain.Order) error
}
