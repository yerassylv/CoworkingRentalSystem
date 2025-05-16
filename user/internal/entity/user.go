package entity

type User struct {
	UserID       string `bson:"user_id" json:"user_id"`
	FullName     string `bson:"full_name" json:"full_name"`
	Email        string `bson:"email" json:"email"`
	Phone        string `bson:"phone" json:"phone"`
	PasswordHash string `bson:"password_hash" json:"-"`
}
