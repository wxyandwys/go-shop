package model

type ShopUsers struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sale string `json:"sale"`
	TimeCreate
}