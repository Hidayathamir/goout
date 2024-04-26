package goouthttp

type ReqErajolBikeOrderDriver struct {
	CustomerID uint `json:"customer_id"`
	DriverID   uint `json:"driver_id"`
	Price      int  `json:"price"`
}

type ResErajolBikeOrderDriver struct {
	Data  ResDataErajolBikeOrderDriver `json:"data"`
	Error string                       `json:"error"`
}

type ResDataErajolBikeOrderDriver struct {
	OrderID uint `json:"order_id"`
}
