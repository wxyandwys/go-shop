package model

type ShopCategorys struct {
	Id int	`json:"id"`
	CategoryName string	`json:"name"`
	CreateAt int	`json:"create"`
	UpdateAt int	`json:"update"`
	DeleteAt int	`json:"delete"`
}