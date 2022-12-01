package epresencesdto

import "time"

type CreateEpresenceRequest struct {
	Type      string    `json:"type"`
	UserID    int       `json:"user_id"`
	Date      time.Time `json:"date"`
	IsApprove bool      `json:"isapprove"`
}

type UpdateEpresenceRequest struct {
	Type      string    `json:"type"`
	UserID    int       `json:"user_id"`
	Date      time.Time `json:"date"`
	IsApprove bool      `json:"isapprove"`
}
