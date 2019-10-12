package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodsher/selectel/pkg/service"
	"strconv"
)

// Ping godoc
// @Summary Ping
// @Description ping
// @Success 200
// @Failure 400
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"pong": "ok",
	})
}

// GetUser godoc
// @Summary Get user.
// @Description Retrieve a user info.
// @ID get-user
// @Produce json
// @Success 200 {object} model.User
// @Failure 400
// @Param id path int true "User ID"
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	s := service.NewUserService()
	user, err := s.GetByID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"balance":    user.Balance,
		"skill":      user.Skill,
		"position":   user.Position,
		"department": user.Department,
	})
}

type TaskParam struct {
	UserID int `json:"user_id" binding:"required;gte=0"`
}

// Task godoc
// @Summary Get list of tasks.
// @Description Retrieve list of tasks.
// @ID get-tasks
// @Accept json
// @Produce json
// @Param user_id body TaskParam true "User ID"
// @Success 200 {array} model.Task
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	var param TaskParam
	c.BindJSON(&param)

	s := service.NewTaskService()
	tasks, err := s.GetByUserID(param.UserID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}
