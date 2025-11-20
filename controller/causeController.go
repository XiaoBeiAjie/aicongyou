package controller

import (
	"net/http"
	"strconv"

	"github.com/XiaoBeiAjie/aicongyou/service"
	"github.com/gin-gonic/gin"
)

// GetAllCauses 获取所有课程
func GetAllCauses(ctx *gin.Context) {
	causes, err := service.GetAllCauses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "获取课程列表失败",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"data":  causes,
		"count": len(causes),
	})
}

// GetCauseByID 根据ID获取课程
func GetCauseByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "无效的课程ID",
			"details": err.Error(),
		})
		return
	}

	cause, err := service.GetCauseByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "课程不存在",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": cause,
	})
}

// GetCausesByParams 根据参数查询课程（支持分页和多条件查询）
func GetCausesByParams(ctx *gin.Context) {
	params := &service.GetCauseQueryParams{}

	// 绑定查询参数
	if err := ctx.ShouldBindQuery(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "参数错误",
			"details": err.Error(),
		})
		return
	}

	// 设置默认分页参数
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}
	// 限制每页最大数量
	if params.PageSize > 100 {
		params.PageSize = 100
	}

	causes, total, err := service.GetCausesByParams(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "查询课程失败",
			"details": err.Error(),
		})
		return
	}

	// 计算总页数
	totalPages := (total + int64(params.PageSize) - 1) / int64(params.PageSize)

	ctx.JSON(http.StatusOK, gin.H{
		"code":        200,
		"data":        causes,
		"total":       total,
		"page":        params.Page,
		"page_size":   params.PageSize,
		"total_pages": totalPages,
	})
}