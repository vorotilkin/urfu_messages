package models

type PostRequest struct {
	Message string `json:"message" form:"message" query:"message" validate:"required"`
	UserID  int32  `json:"user_id" form:"user_id" query:"user_id" validate:"required"`
}
