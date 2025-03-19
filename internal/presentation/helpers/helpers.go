package helpers

import (
	"app/internal/application"
	"errors"
	"github.com/gin-gonic/gin"
)

func JsonError(c *gin.Context, err error, status int) {
	var appErr *application.Err
	var parameters *map[string]string
	if errors.As(err, &appErr) {
		status = appErr.StatusCode
		parameters = &appErr.Parameters
	}
	c.JSON(status, gin.H{
		"code":   status,
		"error":  err.Error(),
		"data":   nil,
		"params": parameters,
	})
}

func JsonOk(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":  200,
		"error": nil,
		"data":  data,
	})
}

func JsonList(c *gin.Context, data []interface{}, limit int, offset int, total int) {
	if data == nil {
		data = []interface{}{}
	}
	c.JSON(200, gin.H{
		"code":   200,
		"error":  nil,
		"data":   data,
		"limit":  limit,
		"offset": offset,
		"total":  total,
	})
}
