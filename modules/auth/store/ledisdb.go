package store

import (
	"bytes"
	"encoding/gob"
	"errors"

	"../../../models"
	"../crypto"

	"github.com/siddontang/ledisdb/ledis"
)

type LedisStore struct {
	db *ledis.DB
}

func NewLedisStore(db *ledis.DB) *LedisStore {
	return &LedisStore{db: db}
}

func (ls *LedisStore) AdminExists(email string) (bool, string, error) {
	found, err := ls.db.Exists([]byte(email))
	if found == 1 {
		return true, email, err
	} else {
		return false, "", err
	}
}

func (ls *LedisStore) Login(email, pass string) (string, error) {
	user, err := ls.UserByEmail(email)

	if err != nil {
		return "", ErrWrongPassword
	}

	passStored := user.Password
	salt := user.Salt
	hpass, err := crypto.HashPassword(pass, []byte(salt))

	if err != nil {
		return "", err
	}

	passOk := crypto.SecureCompare(hpass, []byte(passStored))

	if !passOk {
		return "", ErrWrongPassword
	}

	return user.Id, nil
}

func (ls *LedisStore) UserUpdate(user models.User) (bool, error) {
	_, err := ls.UserByEmail(user.Email)

	if err != nil {
		return false, ErrUserNotFound
	}

	g, err := gobEncode(user)

	if err != nil {
		return false, err
	}

	err = ls.db.Set([]byte(user.Email), g)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ls *LedisStore) GetAllUsers() ([]models.User, error) {
	var users []models.User

	return users, errors.New("")
}

func (ls *LedisStore) LoginWithUser(email, pass string) (models.User, error) {
	user, err := ls.UserByEmail(email)

	if err != nil {
		return models.User{}, ErrWrongPassword
	}

	passStored := user.Password
	salt := user.Salt
	hpass, err := crypto.HashPassword(pass, []byte(salt))

	if err != nil {
		return models.User{}, err
	}

	passOk := crypto.SecureCompare(hpass, []byte(passStored))

	if !passOk {
		return models.User{}, ErrWrongPassword
	}

	return user, nil
}

func (ls *LedisStore) Signin(email, pass string, scope []string) (string, error) {
	// check if the user exists
	_, err := ls.UserByEmail(email)
	if err == nil {
		return "", ErrEmailDuplication
	}

	user, err := NewUser(email, email, pass, scope)
	if err != nil {
		return "", err
	}
	g, err := gobEncodeLedis(user)
	if err != nil {
		return "", err
	}
	err = ls.db.Set([]byte(email), g)

	if err != nil {
		return "", err
	}
	return email, nil
}

func (ls *LedisStore) UserByEmail(email string) (models.User, error) {
	var user models.User
	gobUser, err := ls.db.Get([]byte(email))
	if err == nil {
		if gobUser == nil {
			return user, ErrUserNotFound
		}

		user, err = gobDecodeLedis(gobUser)

		return user, err
	}
	return user, err
}

func gobEncodeLedis(user models.User) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(user)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func gobDecodeLedis(b []byte) (models.User, error) {
	reader := bytes.NewReader(b)
	dec := gob.NewDecoder(reader)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
