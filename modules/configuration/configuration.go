package configuration

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "rserverhub/app"
    "rserverhub/models"
)

func One(c *gin.Context) {
    var host models.Configuration
    app.DB.Where("id = ?", c.Param("id")).Find(&host)
    c.JSON(http.StatusOK, host)
}

func Update(c *gin.Context) {
    var host models.Configuration
    _ = c.BindJSON(&host)
    app.DB.Save(&host)
    c.JSON(http.StatusOK, host)
}

func Delete(c *gin.Context) {
    app.DB.Delete(models.Configuration{}, "id = ?", c.Param("id"))
}
