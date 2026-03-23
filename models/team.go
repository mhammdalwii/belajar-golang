package models

import "time"

type TeamMember struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	// Tambahkan binding:"required" agar tidak boleh kosong
	Name      string    `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Role      string    `gorm:"type:varchar(50);not null" json:"role" binding:"required"` 
	// Tambahkan binding:"required,email" agar format email harus valid
	Email     string    `gorm:"type:varchar(100);unique" json:"email" binding:"required,email"`
	Projects  []Project `gorm:"foreignKey:TeamMemberID" json:"projects"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}