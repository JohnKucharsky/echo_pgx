package serializer

type OrderBody struct {
	ProductId int `json:"product_id" validate:"required"`
	UserId    int `json:"user_id" validate:"required"`
}
