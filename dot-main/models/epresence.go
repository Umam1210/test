package models

import "time"

type Epresence struct {
	ID        int          `json:"id" `
	Type      string       `json:"type"`
	IsApprove bool         `json:"isapprove"`
	User      UserResponse `json:"user"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    int          `json:"user_id"`
	Date      time.Time    `json:"date"`
}

type EpresenceResponse struct {
	ID     int          `json:"id" `
	Type   string       `json:"type"`
	User   UserResponse `json:"user"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID int          `json:"user_id"`
	Date   time.Time    `json:"date"`
}

func (EpresenceResponse) TableName() string {
	return "eprsence"
}
