package domain

type ShippingItem struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	OrderID int64
	Items   []ShippingItem `gorm:"-"`
}

func (s *Shipping) DeliveryDays() int32 {
	var total int32

	for _, item := range s.Items {
		total += item.Quantity
	}

	days := total / 5
	if total%5 != 0 {
		days++
	}

	if days < 1 {
		days = 1
	}

	return days
}
