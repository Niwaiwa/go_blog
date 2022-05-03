package controller

import (
	"go_blog/crypto"
	"go_blog/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserData struct {
	Username string `json:"username" binding:"required"`
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserData struct {
	Username string `json:"username" binding:"required"`
}

func Register(c *gin.Context) {
	var json CreateUserData
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
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := model.GetUserById(userId.(int32))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	} else if user == nil {
		log.Println("userId wrong", userId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "account wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"account":  user.Account,
		"username": user.Username,
	})
}

func UpdateUser(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var json UpdateUserData
	if err := c.ShouldBind(&json); err == nil {
		log.Println(json.Username)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.UpdateUserById(userId.(int32), json.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	} else if user == nil {
		log.Println("userId wrong", userId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "account wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"account":  user.Account,
		"username": user.Username,
	})
}

func DeleteUser(c *gin.Context) {

}
