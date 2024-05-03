package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	r := gin.Default()

	r.POST("/jobs/enqueue", Enqueue)
	r.POST("/jobs/dequeue", Dequeue)
	r.POST("/jobs/:job_id/conclude", Conclude)

	r.GET("/jobs/:job_id", GetJob)

	err := r.Run(":" + PORT)
	if err != nil {
		fmt.Println("error starting server")
	}

	fmt.Println("Server running on port:", PORT)
}
