package dto

// ReqErajolBikeOrderDriver represents the request data structure for ordering a driver.
type ReqErajolBikeOrderDriver struct {
	CustomerID uint
	DriverID   uint
	Price      int
}

// ResErajolBikeOrderDriver represents the response data structure for ordering a driver.
type ResErajolBikeOrderDriver struct {
	OrderID uint
}
