package app

import (
	"strconv"

	"github.com/aloysZy/gin_web/global"
	"github.com/gin-gonic/gin"
)

// GetPage 根据请求 URL 的相关key设置页码
func GetPage(c *gin.Context) int {
	// page := convert.StrTo(c.Query("page")).MustInt()
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		return 1
	}
	return page
}

// GetPageSize 根据请求 URL 和默认最大显示数量设置每页展示最大值
func GetPageSize(c *gin.Context) int {
	// pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize <= 0 {
		return global.AppSetting.Page.DefaultPageSize
	}
	if pageSize > global.AppSetting.Page.MaxPageSize {
		return global.AppSetting.Page.MaxPageSize
	}
	return pageSize
}

// GetPageOffset 计算开始显示数据位置
// pageSize 是每页显示的最大数量
// page 是页数
// 如果page 是 0，那么重头开始显示，page 是 1，计算后也是重头开始显示，page 是 2，那么应该去除第一页的数据
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	// 返回的是数据库查询时候，数据开始位置
	return result
}
