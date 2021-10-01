package model

type Shops struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Img string	`json:"img"`
	Price float64	`json:"price"`
	Sale int	`json:"sale"`
	CategoryId int	`json:"cid"`
	TimeCreate
}
