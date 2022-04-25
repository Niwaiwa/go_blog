package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var json GetUserData
	if err := c.ShouldBind(&json); err == nil {
		log.Println(json.Account)
		log.Println(json.Password)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "account created."})
}

func GetUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  "you are logged in",
		"account": "admin",
	})
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
