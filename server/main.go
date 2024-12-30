package main

import (
	"log"

	"github.com/aspirin100/ArticleManager/internal/types"
	"github.com/gin-gonic/gin"
)

var articleList = []types.Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}

func getAllArticles() []types.Article {
	return articleList
}
