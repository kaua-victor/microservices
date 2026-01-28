package shipping_adapter

import (
	"context"
	"log"
	"time"

	"github.com/kaua-victor/microservices-proto/golang/shipping"
	"github.com/kaua-victor/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client shipping.ShippingClient
}

func NewAdapter(url string) (*Adapter, error) {
	conn, err := grpc.Dial(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		client: shipping.NewShippingClient(conn),
	}, nil
}

func (a *Adapter) CreateShipping(order *domain.Order) error {
	var items []*shipping.ShippingItem

	for _, item := range order.OrderItems {
		items = append(items, &shipping.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    int32(item.Quantity),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := a.client.Create(ctx, &shipping.CreateShippingRequest{
		OrderId: order.ID,
		Items:   items,
	})

	if err != nil {
		return err
	}

	log.Printf("Sucesso! O Shipping calculou o prazo de %d dias para o pedido %d",
		resp.DeliveryDays, order.ID)

	return nil
}
