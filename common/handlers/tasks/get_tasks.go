package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keithchin/taskplanner-gin/common/models"
)

func (h handler) GetTasks(c *gin.Context) {
	var tasks []models.Task

	if result := h.DB.Find(&tasks); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tasks)
}
