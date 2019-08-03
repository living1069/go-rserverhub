package main

import (
    "rserverhub/app"
    "rserverhub/routes"
)
import _ "github.com/jinzhu/gorm/dialects/postgres"

func init() {

}

func main() {
    app.CreateConnection()
    app.CreateEngine()
    routes.Install(app.Engine)
    app.Engine.Run(":8080")
}
