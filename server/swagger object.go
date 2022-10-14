package server

import "time"

type MenyuJson struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Product struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
type Table struct {
	Id     string `json:"id"`
	Number uint8  `json:"number"`
	Buys   bool   `json:"buys"`
}

type ShopStruct struct {
	TableId string `json:"table_id"`
	FoodId  string `json:"food_id"`
	SaladId string `json:"salad_id"`
	DrinkId string `json:"drink_id"`
}
