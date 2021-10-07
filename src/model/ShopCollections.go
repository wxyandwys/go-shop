package model
// 收藏
type ShopCollections struct {
	Id int `json:"id" form:"id"`
	ShopId int `json:"shop_id" form:"shop_id"`
	UserId int `json:"user_id"`
}