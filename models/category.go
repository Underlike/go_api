package models

type Category struct {
	Id int `json:"Id"`
	Title string `json:"Title"`
	Description string `json:"Description"`
}

type Categories []Category

func AllsCategories() *Categories {
	categories := Categories {
		Category{
			Id: 1,
			Title: "Catégorie 1",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco",
		},
	}

	return &categories
}