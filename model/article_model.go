package model

import (
	"errors"
	"go_blog/db"
	"log"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserId    int32     `gorm:"index" json:"user_id"`
	CreatedAt time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetArticles(limit, page int) (*[]Article, error) {
	var articles []Article
	result := db.DBm.Offset(10 * page).Limit(limit).Order("id desc").Find(&articles)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("article not found")
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	log.Println("raw: ", result.RowsAffected)
	return &articles, nil
}

func CreateArticle(userId int32, title, content string) (*Article, error) {
	article := Article{Title: title, Content: content, UserId: userId}
	result := db.DBm.Create(&article)
	if err := result.Error; err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("article id: ", article.ID, "raw: ", result.RowsAffected)
	return &article, nil
}

func GetArticleById(id int32) (*Article, error) {
	var article Article
	result := db.DBm.Where("id = ?", id).First(&article)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("article not found")
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	log.Println("article id: ", article.ID, "raw: ", result.RowsAffected)
	return &article, nil
}

func UpdateArticleById(id int32, title, content string) (*Article, error) {
	var article Article
	result := db.DBm.Where("id = ?", id).First(&article)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("article not found")
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	article.Title = title
	article.Content = content
	result = db.DBm.Save(&article)
	if err := result.Error; err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("article id: ", article.ID, "raw: ", result.RowsAffected)
	return &article, nil
}

func DeleteArticle(id int32) error {
	var article Article
	result := db.DBm.Where("id = ?", id).Delete(&article)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("article not found")
			return err
		}
		log.Println(err)
		return err
	}
	return nil
}
