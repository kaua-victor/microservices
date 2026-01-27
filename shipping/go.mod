module github.com/kaua-victor/microservices/shipping

go 1.25.1

require github.com/kaua-victor/microservices-proto/golang/shipping v0.0.0-00010101000000-000000000000

require (
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.25.12
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.78.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/kaua-victor/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping
