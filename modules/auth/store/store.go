package store

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"../../../models"
	"../crypto"
	"../jwt"
)

var (
	ErrEmailDuplication = errors.New("The email is already in the store")
	ErrUserNotFound     = errors.New("User not found")
	ErrWrongPassword    = errors.New("email or password is incorrent")
)

type UserRepository interface {
	ActivateUser(token string, email string, options jwt.Options) (bool, error)
	Authenticate(token string, options jwt.Options) (bool, error)
	ConfirmUserAccount(token string, email string, options jwt.Options) (bool, error)
	DeactivateUser(token string, email string, text string, options jwt.Options) (bool, error)
	GetUser(token string, options jwt.Options) (models.User, error)
	GetUsers(token string, options jwt.Options) ([]models.User, error)
	Login(email string, password string, options jwt.Options) (string, error)
	RefreshToken(token string, options jwt.Options) (string, error)
	Signin(email string, password string, scopes []string) (string, error)
	TokenRemainTime(token string, options jwt.Options) (string, error)
	TokenScopes(token string, options jwt.Options) ([]string, error)
	UpdatePassword(token string, password string, options jwt.Options) (bool, error)
	UpdateUserDOB(token string, dob int64, options jwt.Options) (bool, error)
	UpdateUserFirstname(token string, firstname string, options jwt.Options) (bool, error)
	UpdateUserLanguage(token string, language string, options jwt.Options) (bool, error)
	UpdateUserSurname(token string, surname string, options jwt.Options) (bool, error)
	ValidateScope(token string, scope string, options jwt.Options) (bool, error)
}

func NewUser(userId, email, pass string, scopes []string) (models.User, error) {
	salt := crypto.GenerateRandomKey(128)
	hpass, err := crypto.HashPassword(pass, salt)

	if err != nil {
		return models.User{}, err
	}

	imageHash := strings.TrimSpace(strings.ToLower(email))
	data := []byte(imageHash)
	var b [16]byte
	b = md5.Sum(data)
	imageHash = hex.EncodeToString(b[:])

	return models.User{
		Id:          userId,
		Email:       strings.TrimSpace(strings.ToLower(email)),
		Password:    string(hpass),
		Salt:        string(salt),
		Scopes:      scopes,
		Active:      true,
		Confirmed:   true,
		SignedUp:    time.Now().Unix(),
		PasswordSet: time.Now().Unix(),
		Language:    "en",
		Picture:     "http://s.gravatar.com/avatar/" + imageHash,
	}, nil
}
