package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	articles := getAllArticles()

	if len(articles) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, val := range articles {
		if val.Content != articleList[i].Content ||
			val.ID != articleList[i].ID ||
			val.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := io.ReadAll(w.Body)
		if err != nil {
			log.Fatal(err)
		}

		pageOK := strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

func TestGetArticleJSON(t *testing.T) {
	r := getRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	// Create a request to send to the above route
	req, err := http.NewRequest("GET", "/article/view/1", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		response := map[string]any{}

		err = json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			log.Fatal(err)
		}

		_, responseOK := response["content"]

		return statusOK && responseOK
	})
}
