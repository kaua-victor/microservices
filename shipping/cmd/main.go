package main

import (
	"log"

	"github.com/kaua-victor/microservices/shipping/config"
	db_adapter "github.com/kaua-victor/microservices/shipping/internal/adapters/db"
	grpc_adapter "github.com/kaua-victor/microservices/shipping/internal/adapters/grpc"
	"github.com/kaua-victor/microservices/shipping/internal/application/core/api"
)

func main() {
	db, err := db_adapter.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	app := api.NewApplication(db)

	grpcAdapter := grpc_adapter.NewAdapter(app)
	grpcAdapter.Run(config.GetApplicationPort())

}
