package dto

import "time"

type ActivityLogResponse struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	UserName        string    `json:"username"`
	Module          string    `json:"module"`
	Action          string    `json:"action"`
	Description     string    `json:"description"`
	ReferenceID     string    `json:"reference_id"`
	ReferenceNumber string    `json:"reference_number"`
	CreatedAt       time.Time `json:"created_at"`
}

type GetAllActivityLogResponseQuery struct { // pakai tag form, karena sifatnya query param di URL
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"search"`
}
