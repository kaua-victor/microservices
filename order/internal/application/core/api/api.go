package api

import (
	"github.com/kaua-victor/microservices/order/internal/application/core/domain"
	"github.com/kaua-victor/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db       ports.DBPort
	payment  ports.PaymentPort
	shipping ports.ShippingPort
	product  ports.ProductPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort, product ports.ProductPort) *Application {
	return &Application{
		db:       db,
		payment:  payment,
		shipping: shipping,
		product:  product,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {

	for _, item := range order.OrderItems {
		exists, err := a.product.Exists(item.ProductCode)
		if err != nil {
			return domain.Order{}, status.Errorf(
				codes.Internal,
				"failed to validate product %s",
				item.ProductCode,
			)
		}

		if !exists {
			return domain.Order{}, status.Errorf(
				codes.InvalidArgument,
				"product %s does not exist",
				item.ProductCode,
			)
		}
	}

	order.Status = "Pending"

	err := a.db.Save(&order)
	if err != nil {
		order.Status = "Canceled"
		a.db.Update(&order)
		return domain.Order{}, err
	}

	// Validate total items
	var totalItems int32
	for _, item := range order.OrderItems {
		totalItems += item.Quantity
	}

	if totalItems > 50 {
		order.Status = "Canceled"
		a.db.Update(&order)
		return domain.Order{}, status.Errorf(
			codes.InvalidArgument,
			"Order with more than 50 items is not allowed",
		)
	}

	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		order.Status = "Canceled"
		a.db.Update(&order)
		return domain.Order{}, paymentErr
	}

	shippingErr := a.shipping.CreateShipping(&order)
	if shippingErr != nil {
		order.Status = "Canceled"
		a.db.Update(&order)
		return domain.Order{}, shippingErr
	}

	order.Status = "Paid"
	err = a.db.Update(&order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
