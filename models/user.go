package models

type User struct {
	Id           string
	Email        string
	Password     string
	Salt         string
	Scopes       []string
	SignedUp     int64
	LastLogin    int64
	LastUpdated  int64
	PasswordSet  int64
	Picture      string
	Active       bool
	Confirmed    bool
	DisabledText string
	Firstname    string
	Surname      string
	Dob          int64
	Language     string
}
