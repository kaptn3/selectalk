package router

import (
	"github.com/gin-gonic/gin"

	_ "github.com/rodsher/selectel/api/swagger/docs"
	"github.com/rodsher/selectel/pkg/controller"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

// @title Hack Innovation
// @version 1.0
// @description Awesome platform.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email isaevm2@yandex.ru

// @host localhost:8080
// @BasePath /api/v1

// Init instantiates and configures a router for REST API.
func Init() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controller.Ping)
		v1.GET("/users/:id", controller.GetUser)
		v1.GET("/tasks/:user_id", controller.GetTasks)
		v1.GET("/achievements", controller.GetAchievements)
		v1.GET("/courses", controller.GetCourses)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
