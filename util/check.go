package util

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    "log"
    "net/http"
)

type response struct {
    Message string `json:"message"`
}

func HandleTransaction(c *gin.Context, tx *gorm.DB) {
    r := recover()

    if r == nil {
        return
    }

    err, ok := r.(error)

    if ok {
        c.JSON(http.StatusInternalServerError, response{Message: err.Error()})
    } else {
        c.JSON(http.StatusInternalServerError, response{Message: "Неизвестная ошибка"})
    }

    tx.Rollback()
}

func Handle(c *gin.Context) {
    r := recover()

    if r == nil {
        return
    }

    err, ok := r.(error)

    if ok {
        c.JSON(http.StatusInternalServerError, response{Message: err.Error()})
    } else {
        c.JSON(http.StatusInternalServerError, response{Message: "Неизвестная ошибка"})
    }
}


func Check(err error) {
    if err != nil {
        log.Println(err.Error())
        panic(err)
    }
}