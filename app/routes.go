package app

import (
    "github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func CreateEngine() {
    Engine = gin.Default()
}
