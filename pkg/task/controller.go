package task

import (
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

    routes := r.Group("/task")
    routes.GET("/sync", h.SyncTasks)
    routes.GET("/starred", h.GetStarredTasks)
    routes.GET("/events", h.GetTaskEvents)
}