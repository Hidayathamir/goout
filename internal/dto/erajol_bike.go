package dto

// ReqOrderDriver represents the request data structure for ordering a driver.
type ReqOrderDriver struct {
	CustomerID uint
	DriverID   uint
	Price      int
}

// ResOrderDriver represents the response data structure for ordering a driver.
type ResOrderDriver struct {
	OrderID uint
}
