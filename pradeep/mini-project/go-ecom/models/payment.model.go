package models

type Payment struct {
	IsDigital        bool `json:"is_digital"`
	IsCashOnDelivery bool `json:"is_cash_on_delivery"`
}
