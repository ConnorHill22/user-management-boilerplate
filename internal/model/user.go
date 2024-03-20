package model

type User struct {
	Id    string
	Email string
}

type UserAttribute struct {
	Id         string
	UserID     string
	Type       UserAttributeType
	Value      string
	CreatedAt  string
	ModifiedAt string
}
