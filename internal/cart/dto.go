package cart


type AddCartDTO struct{
	ProductID uint `json:"product_id"`
	Quantity int `json:"quantity"`
}

type UpdateCartDTO struct{
	Quantity int `json:"quantity"`
}