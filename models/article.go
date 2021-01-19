package models

type Article struct {
	Title string `json:"Title"`
}

type Articles []Article

func AllsArticles() *Articles {
	articles := Articles {
		Article{Title: "Test"},
	}

	return &articles
}