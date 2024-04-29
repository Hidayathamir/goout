package goouthttp

// ReqErajolBikeOrderDriver represents the request structure for ordering a driver.
type ReqErajolBikeOrderDriver struct {
	CustomerID uint `json:"customer_id"`
	DriverID   uint `json:"driver_id"`
	Price      int  `json:"price"`
}

// ResErajolBikeOrderDriver represents the response structure for ordering a driver.
type ResErajolBikeOrderDriver struct {
	Data  ResDataErajolBikeOrderDriver `json:"data"`
	Error string                       `json:"error"`
}

// ResDataErajolBikeOrderDriver represents the response data structure for ordering a driver.
type ResDataErajolBikeOrderDriver struct {
	OrderID uint `json:"order_id"`
}
