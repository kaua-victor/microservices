// package api

// import "github/kaua-victor/microservices/shipping/internal/application/core/domain"

// type Application struct{}

// func NewApplication() *Application {
// 	return &Application{}
// }

// func (a Application) CreateShipping(shipping domain.Shipping) (int32, error) {
// 	return shipping.DeliveryDays(), nil
// }

package api

import (
	"github.com/kaua-victor/microservices/shipping/internal/application/core/domain"
	"github.com/kaua-victor/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) CreateShipping(shipping domain.Shipping) (int32, error) {
	err := a.db.Save(&shipping)
	if err != nil {
		return 0, err
	}
	return shipping.DeliveryDays(), nil
}
