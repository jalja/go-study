package entity

import "time"

type BaseEntity struct {
	ID          uint64 `gorm:"primaryKey"`
	CreatedName string
	UpdatedName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     int8 `gorm:"embedded;comment:逻辑删除0=有效 null=无效;default:0"`
}
