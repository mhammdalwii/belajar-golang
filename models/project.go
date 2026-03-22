package models

import "time"

// ini akan menjadi tabel 'projects'
type Project struct {
ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"type:varchar(150);not null" json:"title"`
	Category     string    `gorm:"type:varchar(100)" json:"category"`
	Status       string    `gorm:"type:varchar(50)" json:"status"`
	TeamMemberID uint       `json:"team_member_id"` 
	TeamMember   TeamMember `gorm:"foreignKey:TeamMemberID" json:"team_member"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}