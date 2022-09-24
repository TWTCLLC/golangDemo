package controller

import (
	"demo/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func WebServiceStart() {
	router := gin.Default()

	router.POST("/callback", service.Callback)
	router.POST("/getMessageList", service.GetMessageList)

	router.Run(":8080")
	fmt.Println("Application Started:")
}
