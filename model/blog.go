package model

type BlogResult struct {
	Status string   `json:"status"`
	Item   BlogItem `json:"item"`
}

type BlogItem struct {
	Id     string             `json:"id"`
	UserId string             `json:"userId"`
	Blogs  []BlogItemListData `json:"blogs"`
}

type BlogItemListData struct {
	Name                     string             `json:"name"`
	Url                      string             `json:"url"`
	SecondaryUrl             string             `json:"secondaryUrl"`
	Nickname                 string             `json:"nickname"`
	Title                    string             `json:"title"`
	Description              string             `json:"description"`
	Default                  string             `json:"default"`
	BlogIconUrl              string             `json:"blogIconUrl"`
	FaviconUrl               string             `json:"faviconUrl"`
	ProfileThumbnailImageUrl string             `json:"profileThumbnailImageUrl"`
	ProfileImageUrl          string             `json:"profileImageUrl"`
	Role                     string             `json:"role"`
	BlogId                   string             `json:"blogId"`
	Statistics               BlogStatisticsData `json:"statistics"`
}

type BlogStatisticsData struct {
	Post       string `json:"post"`
	Comment    string `json:"comment"`
	Trackback  string `json:"trackback"`
	Guestbook  string `json:"guestbook"`
	Invitation string `json:"invitation"`
}
