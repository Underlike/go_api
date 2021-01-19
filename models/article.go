package models

type Article struct {
	Id int `json:"Id"`
	Title string `json:"Title"`
	Description string `json:"Description"`
	Categories []Category `json:"Categories"`
}

type Articles []Article

func AllsArticles() *Articles {
	articles := Articles {
		Article{
			Id: 1,
			Title: "Article 1",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco",
			Categories: []Category{
				Category{
					Id: 1,
					Title: "Catégorie 1",
					Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco",
				},
			},
		},
	}

	return &articles
}