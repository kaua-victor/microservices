package ports

import "github.com/kaua-victor/microservices/order/internal/application/core/domain"

type ShippingPort interface {
	CreateShipping(order *domain.Order) error
}
