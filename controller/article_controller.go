package controller

import (
	"go_blog/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetArticlesQuery struct {
	Limit int `form:"limit,default=10"`
	Page  int `form:"page,default=0"`
}

type CreateArticleData struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type GetArticleData struct {
	ID int32 `uri:"id" binding:"required"`
}

type UpdateArticleData struct {
	Title   *string `json:"title" binding:"required,omitempty"`
	Content *string `json:"content" binding:"required,omitempty"`
}

func GetArticles(c *gin.Context) {
	var qs GetArticlesQuery
	if err := c.ShouldBindQuery(&qs); err == nil {
		log.Println(qs.Limit)
		log.Println(qs.Page)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articles, err := model.GetArticles(qs.Limit, qs.Page)
	log.Println(articles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	} else if articles == nil {
		log.Println("no articles")
		c.JSON(http.StatusBadRequest, gin.H{"error": "account wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"articles": articles,
	})
}

func CreateArticle(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var json CreateArticleData
	if err := c.ShouldBind(&json); err == nil {
		log.Println(json.Title)
		log.Println(json.Content)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := model.CreateArticle(userId.(int32), json.Title, json.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error."})
	} else {
		log.Println(article)
		c.JSON(http.StatusOK, gin.H{
			"status": "Article created.",
			"id":     article.ID,
		})
	}
}

func GetArticle(c *gin.Context) {
	var data GetArticleData
	if err := c.ShouldBindUri(&data); err == nil {
		log.Println(data.ID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := model.GetArticleById(data.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	} else if article == nil {
		log.Println("no article")
		c.JSON(http.StatusNotFound, gin.H{"error": "article wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"article": article,
	})
}

func UpdateArticle(c *gin.Context) {
	var data GetArticleData
	if err := c.ShouldBindUri(&data); err == nil {
		log.Println(data.ID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateData UpdateArticleData
	if err := c.ShouldBindJSON(&updateData); err == nil {
		log.Println(*updateData.Title)
		log.Println(*updateData.Content)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := model.UpdateArticleById(data.ID, *updateData.Title, *updateData.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	} else if article == nil {
		log.Println("no article")
		c.JSON(http.StatusNotFound, gin.H{"error": "article wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"id":     article.ID,
	})
}

func DeleteArticle(c *gin.Context) {
	var data GetArticleData
	if err := c.ShouldBindUri(&data); err == nil {
		log.Println(data.ID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := model.GetArticleById(data.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	} else if article == nil {
		log.Println("no article")
		c.JSON(http.StatusNotFound, gin.H{"error": "article wrong"})
		return
	}

	err = model.DeleteArticleById(data.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "please retry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
