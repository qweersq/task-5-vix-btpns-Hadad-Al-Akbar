package dto

type PhotoUpdateDTO struct {
	ID      uint64 `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Caption string `json:"caption" form:"caption"`
	Url     string `json:"url" form:"url"`
	UserID  uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

// BOOKUPDATEDTO IS A MODEL THAT CLIENT USE WHEN CREATE A NEW BOOK
type PhotoCreateDTO struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Caption string `json:"caption" form:"caption"`
	Url     string `json:"url" form:"url"`
	UserID  uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
