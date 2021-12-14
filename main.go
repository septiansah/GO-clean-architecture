package main

import (
	"log"
	"svc-mahasiswa/config"
	"svc-mahasiswa/controllers"

	"github.com/gin-gonic/gin"
)


func main(){
	conn, errConn := config.ConnectSQL()
	if errConn != nil {
		log.Fatal(errConn)
	}
	studentController := controllers.NewDriverController(conn)
	app := gin.Default()
	studentController.Route(app)
	err := app.Run(":8000"); if err != nil {
		panic(err)
	}
}