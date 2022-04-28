package controller

import (
	"go_blog/crypto"
	"go_blog/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// key := os.Getenv("SECRET_KEY")
		// id := model_redis.GetKey(c, key)
		// id := key
		// if id == nil {
		// 	c.Redirect(http.StatusFound, "/login")
		// 	c.Abort()
		// } else {
		// 	c.Next()
		// }
		c.Next()
	}
}

func Login(c *gin.Context) {
	var json LoginData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.GetUser(json.Account)
	if err != nil || user == nil {
		log.Println("account wrong", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "account or password wrong"})
	}

	err = crypto.CompareHashAndPassword(user.Password, json.Password)
	if err != nil {
		log.Println("password wrong", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "account or password wrong"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "you are logged out",
	})
}
