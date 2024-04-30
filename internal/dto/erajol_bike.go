package dto

// ReqErajolBikeOrderDriver represents the request data structure for ordering a driver.
type ReqErajolBikeOrderDriver struct {
	CustomerID uint `validate:"required,nefield=DriverID"`
	DriverID   uint `validate:"required"`
	Price      int  `validate:"required"`
}

// ResErajolBikeOrderDriver represents the response data structure for ordering a driver.
type ResErajolBikeOrderDriver struct {
	OrderID uint
}
