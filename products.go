package project_go

type Products struct {
  Id int `json:"id"`
  Product_title string `json:"product_title"`
  Price uint `json:"price"`
  Quantity int `json:"quantity"`
  Product_description string `json:"product_description"`
}