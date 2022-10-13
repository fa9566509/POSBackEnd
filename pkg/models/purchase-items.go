package models

type PurchaseItems struct {
	PurchaseStaus string
	TotalPrice    float64
	Discount      float32
	Quantity      int
	LineNo        int
	// TODO: Implement Relation
}
