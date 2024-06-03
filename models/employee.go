package models

import "time"

type Employee struct {
    ID           uint      `gorm:"primary_key" json:"id"`
    Name         string    `gorm:"not null" json:"name"`
    Email        string    `gorm:"not null;unique" json:"email"`
    PhoneNumber  string    `json:"phone_number"`
    JoinDate     time.Time `json:"join_date"`
    Status       string    `json:"status"`
    Address      string    `json:"address"`
    NIK          string    `gorm:"not null;unique" json:"nik"`
    Gender       string    `json:"gender"`
    Position     string    `json:"position"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
