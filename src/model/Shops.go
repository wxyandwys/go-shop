package model

// 商品
type Shops struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Img string	`json:"img"`
	Price float64	`json:"price"`
	Num int `json:"num"`
	Sale int	`json:"sale"`
	CategoryId int	`json:"cid"`
	Quote int `json:"quote"`
	NoneSku bool `json:"none_sku"`
	HideStock bool `json:"hide_stock"`
	Picture string `json:"picture"`
	TimeCreate
}

//规格类目
type ShopTrees struct {
	Id int `json:"id"`
	K string `json:"k"`
	Ks string `json:"k_s"`
	LargeImageMode bool `json:"largeImageMode"`
	ShopsId int `json:"shop_id"`
	//Vs map[string]interface{}
	Vs []ShopTreeVs	`json:"v"`
}

// 规格
type ShopTreeVs struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ImgUrl string `json:"imgUrl"`
	PreviewImgUrl string `json:"previewImgUrl"`
	TreeId int `json:"tree_id"`
}

//所有 sku 的组合列表
type ShopLists struct {
	Id int `json:"id"`
	Price float32 `json:"price"`
	StockNum int `json:"stock_num"`
	ShopId	int `json:"shop_id"`
	Slsu [] ShopListSkus `json:"slsu"`
}

// 规格组合
type ShopListSkus struct {
	Id int `json:"id"`
	K string `json:"k"`
	V string `json:"v"`
	Ks string `json:"k_s"`
}