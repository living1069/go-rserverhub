package routes

import (
    "github.com/gin-gonic/gin"
    "rserverhub/modules/configuration"
    "rserverhub/modules/host"
)

func Install(e *gin.Engine) {

    hostGroup := e.Group("/api/host")
    {
        hostGroup.GET("/", host.List)
        hostGroup.GET("/:id", host.One)
        hostGroup.POST("/", host.Create)
        hostGroup.DELETE("/:id", host.Delete)
        hostGroup.PATCH("/", host.Update)
        hostGroup.POST("/update", host.AgentUpdate)
        hostGroup.GET("/:id/servers", host.Servers)
    }

    configurationGroup := e.Group("/api/configuration")
    {
        configurationGroup.GET("/:id", configuration.One)
        configurationGroup.POST("/", configuration.Update)
    }
}
