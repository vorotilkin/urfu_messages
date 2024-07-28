package models

type DeleteRequest struct {
	ID int32 `json:"id" param:"id" validate:"required"`
}
