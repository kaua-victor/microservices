// package grpc_adapter

// import (
// 	"context"

// 	shippingpb "github.com/kaua-victor/microservices-proto/golang/shipping"
// 	"github.com/kaua-victor/microservices/shipping/internal/application/core/api"
// 	"github.com/kaua-victor/microservices/shipping/internal/application/core/domain"
// )

// type Adapter struct {
// 	app shippingpb.UnimplementedShippingServer
// 	api *api.Application
// }

// func NewAdapter(api *api.Application) *Adapter {
// 	return &Adapter{api: api}
// }

// func (a *Adapter) Create(
// 	ctx context.Context,
// 	req *shippingpb.CreateShippingRequest,
// ) (*shippingpb.CreateShippingResponse, error) {

// 	// proto → domain
// 	var items []domain.ShippingItem
// 	for _, item := range req.Items {
// 		items = append(items, domain.ShippingItem{
// 			ProductCode: item.ProductCode,
// 			Quantity:    int32(item.Quantity),
// 		})
// 	}

// 	shipping := domain.Shipping{
// 		OrderID: req.OrderId,
// 		Items:   items,
// 	}

// 	// domain → application
// 	days, err := a.api.CreateShipping(shipping)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// domain → proto
// 	return &shippingpb.CreateShippingResponse{
// 		DeliveryDays: days,
// 	}, nil
// }

package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kaua-victor/microservices-proto/golang/shipping"
	"github.com/kaua-victor/microservices/shipping/internal/application/core/api"
	"github.com/kaua-victor/microservices/shipping/internal/application/core/domain"
)

type Adapter struct {
	api *api.Application
	shipping.UnimplementedShippingServer
}

func NewAdapter(api *api.Application) *Adapter {
	return &Adapter{api: api}
}

func (a Adapter) Create(
	ctx context.Context,
	req *shipping.CreateShippingRequest,
) (*shipping.CreateShippingResponse, error) {

	var items []domain.ShippingItem
	for _, item := range req.Items {
		items = append(items, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	shippingDomain := domain.Shipping{
		OrderID: req.OrderId,
		Items:   items,
	}

	result, err := a.api.CreateShipping(shippingDomain)
	if err != nil {
		return nil, err
	}

	return &shipping.CreateShippingResponse{
		DeliveryDays: result,
	}, nil

}

func (a *Adapter) Run(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	shipping.RegisterShippingServer(server, a)

	reflection.Register(server)

	log.Printf("Shipping service running on port %d", port)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
