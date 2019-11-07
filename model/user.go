package model

type User struct {
	Email      string
	PostsCount int
	Name       string
	Avatar     string
}

type ApiAccessToken struct {
	UserID int
	User   User
	Key    string
	Enable string
}

// GetUser 用 token 找用戶
func GetUser(Token interface{}) (User, error) {
	var apiAccessToken ApiAccessToken
	var user User

	DB.First(&apiAccessToken, "key = ?", Token)
	result := DB.First(&user, "id = ?", apiAccessToken.UserID)

	return user, result.Error
}
