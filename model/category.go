package model

type CategoryResult struct {
	Status string         `json:"status"`
	Item   []CategoryItem `json:"item"`
}

type CategoryItem struct {
	Url          string       `json:"url"`
	SecondaryUrl string       `json:"secondaryUrl"`
	Categories   CategoryData `json:"categories"`
}

type CategoryData struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Parent  string `json:"parent"`
	Label   string `json:"label"`
	Entries string `json:"entries"`
}
