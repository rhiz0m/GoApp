package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	server := gin.Default() // Engine.HTTP-Server. Ctrl + C. Shut down.
server.GET("/programs", getExercises)

	server.Run(":8080")
}

func getExercises(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H { 
		"Strong Man": []string{"Atlas Stones, Deadlift, Other"},
		"Powerlifting": []string{"Bench Press, Back Squat, Deadlift"},
		"Grip Strength": "Blob"})
}