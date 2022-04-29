package controller

import (
	"go_blog/crypto"
	"go_blog/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserData struct {
	Username string `json:"username" binding:"required"`
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

	encryptPw, err := crypto.PasswordEncrypt(json.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error."})
		return
	}

	user, err := model.CreateUser(json.Username, json.Account, encryptPw)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error."})
	} else {
		log.Println(user)
		c.JSON(http.StatusOK, gin.H{"status": "account created."})
	}
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
