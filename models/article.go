package models

type Article struct {
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type Articles []Article

func AllsArticles() *Articles {
	articles := Articles {
		Article{
			Title: "Test",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco"
		},
	}

	return &articles
}