package models

type Category struct {
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type Categories []Category

func AllsCategories() *Categories {
	categories := Categories {
		Category{
			Title: "Test",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco",
		},
	}

	return &categories
}