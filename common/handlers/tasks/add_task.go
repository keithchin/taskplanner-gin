package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keithchin/taskplanner-gin/common/models"
)

type AddTaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Comments    string `json:"comments"`
}

func (h handler) AddTask(c *gin.Context) {
	body := AddTaskRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	task.Title = body.Title
	task.Description = body.Description
	task.Status = body.Status
	task.Comments = body.Comments

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &task)
}
