package model

type ShopUsers struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sale string `json:"sale"`
	Img string `json:"img"`
	TimeCreate
}