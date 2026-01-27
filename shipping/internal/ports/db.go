package ports

import "github.com/kaua-victor/microservices/shipping/internal/application/core/domain"

type DBPort interface {
	Save(*domain.Shipping) error
}
