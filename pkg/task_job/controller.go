package taskjob

import (
	"divviup-client/pkg/common/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := r.Group("/taskjob", middleware.TokenAuthMiddleware(db))
    routes.POST("/add", h.AddTaskJob)
    routes.POST("/delete", h.DeleteTaskJob)
    routes.GET("/", h.GetTaskJob)
}