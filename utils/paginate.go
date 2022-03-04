package utils

import (
	"goProject3/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limitQuery, _ := strconv.Atoi(c.DefaultQuery("limit", models.DefaultLimit))
	pageQuery, _ := strconv.Atoi(c.DefaultQuery("page", models.DefaultPage))
	sortQuery := c.DefaultQuery("sort", models.DefaultSort)

	return models.Pagination{
		Limit: limitQuery,
		Page:  pageQuery,
		Sort:  sortQuery,
	}
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limitQuery, _ := strconv.Atoi(c.DefaultQuery("limit", models.DefaultLimit))
		pageQuery, _ := strconv.Atoi(c.DefaultQuery("page", models.DefaultPage))
		//sortQuery := c.DefaultQuery("sort", models.DefaultSort)

		offset := (pageQuery - 1) * limitQuery
		return db.Offset(offset).Limit(limitQuery)
	}
}
