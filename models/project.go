package models

import "time"

// ini akan menjadi tabel 'projects'
type Project struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"type:varchar(150);not null" json:"title"`
	Category     string    `gorm:"type:varchar(100)" json:"category"` // Contoh: Web Development, IoT & Automation, UI/UX
	Status       string    `gorm:"type:varchar(50)" json:"status"`    // Contoh: On Progress, Maintenance, Completed
	
	//  Menghubungkan proyek ini dengan ID anggota tim (PIC)
	TeamMemberID uint      `json:"team_member_id"` 
	
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}