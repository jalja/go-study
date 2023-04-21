package model

import "time"

type BaseModel struct {
	ID          uint `gorm:"primaryKey"`
	CreatedName string
	UpdatedName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     int8 `gorm:"embedded;comment:逻辑删除0=有效 null=无效;default:0"`
}
