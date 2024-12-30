package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aspirin100/ArticleManager/internal/types"
	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")

}

func getArticle(ctx *gin.Context) {

	articleID, err := strconv.Atoi(ctx.Param("article_id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	article, err := getArticleByID(articleID)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	render(ctx, gin.H{
		"title":   "Home Page",
		"payload": article}, "index.html")
}

func getArticleByID(id int) (*types.Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}

func render(ctx *gin.Context, data gin.H, templateName string) {
	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		ctx.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		ctx.XML(http.StatusOK, data["payload"])
	default:
		ctx.HTML(http.StatusOK, templateName, data)
	}

}
