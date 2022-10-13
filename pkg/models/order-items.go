package models

type OrderItems struct {
	OrderStatus string
	TotalPrice  float64
	Discount    float32
	Quantity    int
	LineNo      int
}
