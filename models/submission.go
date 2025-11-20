package models

import "time"

type Submission struct {
	ID            uint64 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Content       string `gorm:"column:content;type:varchar(512);not null;comment:提交内容" json:"content"`
	Score         int    `gorm:"column:score;default:0;comment:得分" json:"score"`
	StudentID     uint64 `gorm:"column:student_id;not null;comment:学生id" json:"student_id"`
	TaskID        uint64 `gorm:"column:task_id;not null;index:idx_task_id;comment:任务id" json:"task_id"`
	IsSubmited    int    `gorm:"column:is_submited;default:0;comment:是否提交" json:"is_submited"`
	TeacherComment string `gorm:"column:teacher_comment;type:varchar(512);comment:教师评语" json:"teacher_comment"`
	IsRecommend   int    `gorm:"column:is_recommend;default:0;index:idx_is_recommend;comment:是否推荐" json:"is_recommend"`
	SubmittedAt    *time.Time `gorm:"column:submitted_at;comment:提交时间" json:"submitted_at"`
	CreatedAt      time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at;index:idx_deleted_at" json:"deleted_at,omitempty"`
}

// TableName 指定表名
func (Submission) TableName() string {
	return "submission"
}