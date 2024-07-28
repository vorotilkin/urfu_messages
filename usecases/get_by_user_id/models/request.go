package models

type GetByUserIDRequest struct {
	ID int32 `json:"id" param:"id" validate:"required"`
}
