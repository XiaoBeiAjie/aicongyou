package models

type UserRole string

const (
	RoleStudent UserRole = "student"
	RoleTeacher UserRole = "teacher"
	RoleAdmin UserRole = "admin"
)

var RoleMap = map[UserRole]string {
	RoleStudent: "学生",
	RoleTeacher: "教师",
	RoleAdmin: "管理员",
}

func (r UserRole) DisplayName() string {
	if name, exists := RoleMap[r]; exists {
		return name
	}
	return "未知角色"
}