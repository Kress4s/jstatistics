package common

import (
	"time"

	"gorm.io/gorm"
)

// Base 基础模型定义
type Base struct {
	CreateAt time.Time `gorm:"column:create_at;type:timestamptz;not null;comment:创建时间" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;type:timestamptz;not null;comment:最后一次更新时间" json:"update_at"`
	CreateBy string    `gorm:"column:create_by;type:varchar(40);not null;comment:创建者ID" json:"create_by"`
	UpdateBy string    `gorm:"column:update_by;type:varchar(40);not null;comment:最后一次更新者ID" json:"update_by"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreateAt = now
	b.UpdateAt = now
	return nil
}
