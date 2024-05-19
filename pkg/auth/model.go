package auth

type User struct {
	ID           int    `bson:"id"`
	IsBot        bool   `bson:"is_bot"`
	FirstName    string `bson:"first_name"`
	LastName     string `bson:"last_name"`
	UserName     string `bson:"username"`
	LanguageCode string `bson:"language_code"`
}
