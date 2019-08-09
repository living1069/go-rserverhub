package app

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "os"
)

var Engine *gin.Engine

func CreateEngine() {
    Engine = gin.Default()    
    Engine.Static("/assets/", os.Getenv("ASSETS_PATH") + "/")
    Engine.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"*"},
        AllowHeaders:     []string{"*"},
    }))
}
