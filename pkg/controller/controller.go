package controller

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodsher/selectel/pkg/service"
)

// Ping godoc
// @Summary Ping
// @Description ping
// @Success 200
// @Failure 400
// @Router /api/v1/ping [get]
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
// @Router /api/v1/users/{id} [get]
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

	c.JSON(http.StatusOK, user)
}

// Task godoc
// @Summary Get list of tasks by user identifier.
// @Description Retrieve list of tasks using user identifier.
// @ID get-tasks
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} model.Task
// @Failure 400
// @Router /api/v1/tasks/{user_id} [get]
func GetTasks(c *gin.Context) {
	param := c.Param("user_id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	s := service.NewTaskService()
	tasks, err := s.GetByUserID(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Achievement godoc
// @Summary Get list of achievements by user id.
// @Description Retrieve list of achievements using user identifier.
// @ID get-achievements
// @Produce json
// @Success 200 {array} model.Achievement
// @Failure 400
// @Router /api/v1/achievements [get]
func GetAchievements(c *gin.Context) {
	s := service.NewAchievementService()
	achievements, err := s.Get()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, achievements)
}

// Course godoc
// @Summary Get list of courses.
// @Description Retrieve list of courses.
// @ID get-courses
// @Produce json
// @Success 200 {array} model.Course
// @Failure 400
// @Router /api/v1/achievements [get]
func GetCourses(c *gin.Context) {
	s := service.NewCourseService()
	courses, err := s.Get()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, courses)
}
