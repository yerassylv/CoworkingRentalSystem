package domain

type User struct {
	ID    string `bson:"_id"`
	Name  string
	Email string
	Role  string
}

type UserRepository interface {
	Register(*User) error
	GetByID(id string) (*User, error)
	ListAll() ([]*User, error)
}
