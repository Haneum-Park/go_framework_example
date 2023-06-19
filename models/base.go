package models

type BaseEntity struct {
	Id        uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"` // NOTE primary key
	CreatedAt string `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt string `gorm:"type:timestamp;not null" json:"updated_at"`
}
