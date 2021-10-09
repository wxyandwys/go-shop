package model

type ShopAddresses struct {
	Id int `json:"id" form:"id"`
	UserId int `json:"user_id" form:"user_id"`
	Name string `json:"name" form:"name"`
	Tel string `json:"tel" form:"tel"`
	Province string `json:"province" form:"province"`
	City string `json:"city" form:"city"`
	County string `json:"county" form:"county"`
	AddressDetail string `json:"addressDetail" form:"addressDetail"`
	AreaCode string `json:"areaCode" form:"areaCode"`
	PostalCode string `json:"postalCode" form:"postalCode"`
	IsDefault bool `json:"isDefault" form:"isDefault"`
	TimeCreate
}