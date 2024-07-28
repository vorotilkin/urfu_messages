package models

type PutRequest struct {
	ID      int32  `json:"id" form:"id" query:"id" validate:"required"`
	Message string `json:"message" form:"message" query:"message" validate:"required"`
}
