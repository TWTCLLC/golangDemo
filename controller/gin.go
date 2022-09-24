package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func WebServiceStart() {
	router := gin.Default()

	router.Run(":8080")
	fmt.Println("Application Started:")
}
