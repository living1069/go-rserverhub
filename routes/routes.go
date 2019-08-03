package routes

import (
    "github.com/gin-gonic/gin"
    "rserverhub/modules/configuration"
    "rserverhub/modules/host"
    "rserverhub/modules/server"
)

func Install(e *gin.Engine) {

    hostGroup := e.Group("/api/host")
    {
        hostGroup.GET("", host.List)
        hostGroup.GET("/", host.List)
        hostGroup.GET("/:id", host.One)
        hostGroup.POST("/", host.Create)
        hostGroup.DELETE("/:id", host.Delete)
        hostGroup.PATCH("", host.Update)
        hostGroup.PATCH("/", host.Update)
        hostGroup.POST("/update", host.AgentUpdate)
        hostGroup.GET("/:id/servers", host.Servers)
    }

    configurationGroup := e.Group("/api/configuration")
    {
        configurationGroup.GET("/:id", configuration.One)
        configurationGroup.POST("", configuration.Update)
        configurationGroup.POST("/", configuration.Update)
    }

    serverGroup := e.Group("/api/server")
    {
        serverGroup.GET("/", server.All)
        serverGroup.GET("", server.All)
        serverGroup.POST("/", server.Update)
        serverGroup.POST("", server.Update)
        serverGroup.PATCH("/", server.Update)
        serverGroup.PATCH("", server.Update)
        serverGroup.GET("/info/:id", server.Get)
        serverGroup.GET("/logs/:id", server.GetLogs)
        serverGroup.GET("/sessions/:id", server.GetSessions)

        serverGroup.POST("/start/:id", server.StartServer)
        serverGroup.POST("/stop/:id", server.StopServer)
        serverGroup.POST("/command/:id", server.SendCommand)

        serverGroup.POST("/update", server.AgentUpdate)
        serverGroup.POST("/logs", server.SaveLogs)
        serverGroup.POST("/stats", server.SaveStats)

        serverGroup.DELETE("/sessions/:id", server.DeleteSessions)
        serverGroup.DELETE("/delete/:id", server.Delete)
    }

    indexHandler := func(c *gin.Context) {
        c.File("./assets/index.html")
    }

    e.GET("/", indexHandler)
    e.GET("/servers", indexHandler)
    e.GET("/server/:id", indexHandler)
    e.GET("/hosts", indexHandler)
    e.GET("/host/:id", indexHandler)
    e.GET("/configurations", indexHandler)
}
