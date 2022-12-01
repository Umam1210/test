package epresencesdto

import (
	"journey/models"
	"time"
)

type EpresenceResponse struct {
	ID        int                 `json:"id" `
	Type      string              `json:"type"`
	User      models.UserResponse `json:"user"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int                 `json:"user_id"`
	Date      time.Time           `json:"date"`
	IsApprove bool                `json:"isapprove"`
}
