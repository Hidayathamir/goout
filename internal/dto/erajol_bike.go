package dto

// ReqOrderDriver -.
type ReqOrderDriver struct {
	CustomerID uint
	DriverID   uint
	Price      int
}

// ResOrderDriver -.
type ResOrderDriver struct {
	OrderID uint
}
