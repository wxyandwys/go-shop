package model
// 商品分类
type ShopCategorys struct {
	Id int	`json:"id"`
	CategoryName string	`json:"name"`
	TimeCreate
}