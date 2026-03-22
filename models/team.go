package models

import "time"

// Struct ini akan diubah menjadi tabel 'team_members' oleh GORM
type TeamMember struct {
ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Role      string    `gorm:"type:varchar(50);not null" json:"role"` 
	Email     string    `gorm:"type:varchar(100);unique" json:"email"`
	// Tambahkan baris ini untuk relasi One-to-Many:
	Projects  []Project `gorm:"foreignKey:TeamMemberID" json:"projects"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}