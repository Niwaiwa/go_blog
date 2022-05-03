package controller

import (
	"context"
	"fmt"
	"go_blog/crypto"
	"go_blog/db"
	"go_blog/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type LoginData struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth header Authorization
		// log.Println(c.Request.Header)
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		mapClaims, err := crypto.ParseValidJwtToken(tokenString)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		log.Println(mapClaims)
		userId := mapClaims.UserId

		ctx := context.Background()
		val, err := db.Rdb.Get(ctx, "_cache_login:"+fmt.Sprint(userId)).Result()
		switch {
		case err == redis.Nil:
			log.Println("cache key does not exist")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		case err != nil:
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		default:
			log.Println(userId, val)
		}

		// log.Println(userId.(type))
		c.Set("userId", userId)
		c.Next()
	}
}

func Login(c *gin.Context) {
	var json LoginData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(json)
	user, err := model.GetUser(json.Account)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	} else if user == nil {
		log.Println("account wrong: ", json.Account)
		c.JSON(http.StatusBadRequest, gin.H{"error": "account or password wrong"})
		return
	}

	err = crypto.CompareHashAndPassword(user.Password, json.Password)
	if err != nil {
		log.Println("password wrong", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "account or password wrong"})
		return
	}

	ctx := context.Background()
	err = db.Rdb.SetEX(ctx, "_cache_login:"+fmt.Sprint(user.ID), time.Now(), time.Second*60).Err()
	if err != nil {
		log.Println("redis error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	}

	tokenString, err := crypto.NewJwtToken(int32(user.ID))
	if err != nil {
		log.Println("jwt error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	} else {
		c.Header("Authorization", tokenString)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	}
}

func Logout(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx := context.Background()
	err := db.Rdb.Del(ctx, "_cache_login:"+fmt.Sprint(userId)).Err()
	if err != nil {
		log.Println("redis error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "you are logged out",
	})
}
