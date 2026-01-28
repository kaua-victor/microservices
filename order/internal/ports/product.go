package ports

type ProductPort interface {
	Exists(productCode string) (bool, error)
}
