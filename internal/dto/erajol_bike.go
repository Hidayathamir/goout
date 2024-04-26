package dto

type ReqOrderDriver struct {
	CustomerID uint
	DriverID   uint
	Price      int
}

type ResOrderDriver struct {
	OrderID uint
}
