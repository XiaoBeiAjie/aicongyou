package service

import (
	"github.com/XiaoBeiAjie/aicongyou/db"
	"github.com/XiaoBeiAjie/aicongyou/models"
)

// GetCauseQueryParams 课程查询参数
type GetCauseQueryParams struct {
	ID        uint64 `json:"id,omitempty"`         // 课程ID
	Name      string `json:"name,omitempty"`       // 课程名称（模糊查询）
	Semester  string `json:"semester,omitempty"`   // 开设学期
	TeacherID uint64 `json:"teacher_id,omitempty"` // 教师ID
	Page      int    `json:"page,omitempty"`       // 页码
	PageSize  int    `json:"page_size,omitempty"`  // 每页数量
}

// GetAllCauses 获取所有课程
func GetAllCauses() ([]models.Cause, error) {
	var causes []models.Cause
	db := db.GetDB()
	if err := db.Find(&causes).Error; err != nil {
		return nil, err
	}
	return causes, nil
}

// GetCauseByID 根据ID获取课程
func GetCauseByID(id uint64) (*models.Cause, error) {
	var cause models.Cause
	db := db.GetDB()
	if err := db.First(&cause, id).Error; err != nil {
		return nil, err
	}
	return &cause, nil
}

// GetCausesByParams 根据参数查询课程
func GetCausesByParams(params *GetCauseQueryParams) ([]models.Cause, int64, error) {
	var causes []models.Cause
	var total int64

	db := db.GetDB()
	query := db.Model(&models.Cause{})

	// 根据ID查询
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}

	// 根据课程名称模糊查询
	if params.Name != "" {
		query = query.Where("name LIKE ?", "%"+params.Name+"%")
	}

	// 根据学期查询
	if params.Semester != "" {
		query = query.Where("semester = ?", params.Semester)
	}

	// 根据教师ID查询
	if params.TeacherID != 0 {
		query = query.Where("teacher_id = ?", params.TeacherID)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页处理
	if params.Page > 0 && params.PageSize > 0 {
		offset := (params.Page - 1) * params.PageSize
		query = query.Offset(offset).Limit(params.PageSize)
	}

	// 执行查询
	if err := query.Find(&causes).Error; err != nil {
		return nil, 0, err
	}

	return causes, total, nil
}

// CreateCause 创建课程
func CreateCause(cause *models.Cause) error {
	db := db.GetDB()
	return db.Create(cause).Error
}

// UpdateCause 更新课程
func UpdateCause(cause *models.Cause) error {
	db := db.GetDB()
	return db.Save(cause).Error
}

// DeleteCause 删除课程（软删除）
func DeleteCause(id uint64) error {
	db := db.GetDB()
	return db.Delete(&models.Cause{}, id).Error
}