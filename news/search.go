package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
const apiKey = "bf9a4c68daf4453c9edd9fe70fbd3b6e"
const newsUrl = "https://newsapi.org/v1"
const sourceUrl = newsUrl + "/sources?language=en&category=%s"
const articleUrl = newsUrl + "/articles?source=%s&sortBy=latest&apiKey=%s"

type source struct {
	ID string `json:"id"`
}

type sourcesAPI struct {
	Sources []source `json:"sources"`
}

type articlesAPI struct {
	Articles []Article `json:"articles"`
}

func getSources(category string) []source {
	body := getData(sourceURL(category))
	var sourceAPI sourcesAPI
	json.Unmarshal(body, &sourceAPI)

	return sourceAPI.Sources
}

func getArticles(sources []source) []Article {
	var articles []Article

	for _, source := range sources {
		body := getData(articleURL(source.ID))
		var articleAPI articlesAPI
		json.Unmarshal(body, &articleAPI)

		articles = append(articles, articleAPI.Articles...)
	}

	return articles
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func sourceURL(category string) string {
	return fmt.Sprintf(sourceUrl, category)
}

func articleURL(source string) string {
	return fmt.Sprintf(articleUrl, source, apiKey)
}
