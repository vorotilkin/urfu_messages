package models

type GetByIDRequest struct {
	ID int32 `json:"id" param:"id" validate:"required"`
}
