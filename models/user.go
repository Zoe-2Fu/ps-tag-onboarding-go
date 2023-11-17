package models

type User struct {
	ID        string `json:"id" bson:"id"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Age       int    `json:"age" bson:"age"`
}

func NewUser(id string, fn string, ln string, email string, age int) User {
	return User{id, fn, ln, email, age}
}
