package serializer

type ProductBody struct {
	Name         string `json:"name" validate:"required"`
	SerialNumber string `json:"serial_number" validate:"required"`
}
