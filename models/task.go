package models

import "time"

type Task struct {
	ID                    uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name                  string    `gorm:"column:name;type:varchar(100);not null;comment:任务名" json:"name"`
	CauseID               uint64    `gorm:"column:cause_id;not null;index:idx_cause_id;comment:所属的课程id" json:"cause_id"`
	CoreRequirements      string    `gorm:"column:core_requirements;type:varchar(512);comment:任务核心要求" json:"core_requirements"`
	SubmitedCount         int       `gorm:"column:submited_count;default:0;comment:已提交人数" json:"submited_count"`
	AllCount              int       `gorm:"column:all_count;default:0;comment:课程全部人数" json:"all_count"`
	ViewCount             int       `gorm:"column:view_count;default:0;comment:查看次数" json:"view_count"`
	DiscussCount          int       `gorm:"column:discuss_count;default:0;comment:讨论次数" json:"discuss_count"`
	ExcellentSubmissionIDs string   `gorm:"column:excellent_submission_ids;type:varchar(256);comment:优秀作业id" json:"excellent_submission_ids"`
	StartTime             time.Time `gorm:"column:start_time;default:CURRENT_TIMESTAMP;comment:任务开始时间" json:"start_time"`
	ExpirationDate        time.Time `gorm:"column:expiration_date;default:CURRENT_TIMESTAMP;comment:任务截止时间" json:"expiration_date"`
	ScoreCriteria         string    `gorm:"column:score_criteria;type:varchar(512);comment:任务给分标准" json:"score_criteria"`
	CreatedAt             time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	DeletedAt             *time.Time `gorm:"column:deleted_at;index:idx_deleted_at" json:"deleted_at,omitempty"`
	
}

// TableName 指定表名
func (Task) TableName() string {
	return "task"
}