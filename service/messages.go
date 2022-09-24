package service

import (
	"demo/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetMessageList(c *gin.Context) {
	userID := c.PostForm("userID")
	fmt.Printf("userIDaaaa: %v\n", userID)
	messageList := database.GetMessage(userID)
	c.JSON(200, messageList)
}
