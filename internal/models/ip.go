package models

import (
	"time"
)

type IP struct {
	ID        uint      `gorm:"primarykey"`
	Data      string    `gorm:"column:data;uniqueIndex:uk_data;comment:'IP:PORT'"`
	Source    string    `gorm:"column:source;comment:'数据来源'"`
	CreatedAt time.Time `gorm:"column:created_at;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:'更新时间'"`
}

func (*IP) TableName() string {
	return "ips"
}
