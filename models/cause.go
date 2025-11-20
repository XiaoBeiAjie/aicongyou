package models

import "time"

type Cause struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(100);not null;comment:课程名" json:"name"`
	Info      string    `gorm:"column:info;type:varchar(512);default:暂无课程简介;comment:课程简介" json:"info"`
	Semester  string    `gorm:"column:semester;type:varchar(256);default:2026-9;comment:开设学期" json:"semester"`
	TeacherID uint64    `gorm:"column:teacher_id;not null;comment:授课教师id" json:"teacher_id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index:idx_deleted_at" json:"deleted_at,omitempty"`
	
}

// TableName 指定表名
func (Cause) TableName() string {
	return "cause"
}