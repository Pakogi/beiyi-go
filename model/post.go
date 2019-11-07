package model

type Post struct {
	UserID   int    `json:"user_id"`
	Category string `json:"category"`
	Context  string `json:"context"`
	Location string `json:"location"`
	IsMask   string `json:"is_mask"`
}

func GetPosts() ([]Post, error) {
	var posts []Post
	result := DB.Where("").Find(&posts)

	return posts, result.Error
}
