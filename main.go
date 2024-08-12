package main

import (
	"FirstGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	server := gin.Default() // Engine.HTTP-Server. Ctrl + C. Shut down.
server.GET("/programs", getExercises)

	server.Run(":8080")
}

func getExercises(context *gin.Context) {
	programs := models.GetPrograms()
	context.JSON(http.StatusOK, programs)
}