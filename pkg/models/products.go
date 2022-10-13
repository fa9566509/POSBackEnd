package models

type Products struct {
	LineNo    int
	Barcode   string
	ALU       string
	ProdName  string
	BuyPrice  float32
	SalePrice float32
	// TODO: Implement Supplier relationship many to many
}
