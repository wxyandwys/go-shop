package model

type Login struct {
	username string `json:"username"`
	password string `json:"password"`
	value int `json:"value"`
}