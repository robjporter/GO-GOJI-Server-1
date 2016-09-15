package store

import (
	"log"
	"time"

	"../../../models"
	"../crypto"

	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoStore struct {
	db *mgo.Collection
}

func NewMongoStore(db *mgo.Collection) *MongoStore {
	return &MongoStore{db: db}
}

func (ms *MongoStore) AdminExists(email string) (bool, string, error) {
	_, err := ms.UserByEmail(email)
	if err == nil {
		return true, email, err
	} else {
		return false, "", err
	}
}

func (ms *MongoStore) Signin(email, pass string, scope []string) (string, error) {
	// check if the user exists
	_, err := ms.UserByEmail(email)
	if err == nil {
		return "", ErrEmailDuplication
	}

	user, err := NewUser(email, email, pass, scope)

	if err != nil {
		return "", err
	}

	err = ms.db.Insert(&user)

	if err != nil {
		return "", err
	}

	return email, nil
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (ms *MongoStore) LoginWithUser(email, pass string) (models.User, error) {
	defer TimeTrack(time.Now(), "LoginWithUser")
	user, err := ms.UserByEmail(email)

	if err != nil {
		return models.User{}, ErrWrongPassword
	}

	passStored := user.Password
	hpass, err := encryptWithCrypto(user.Salt, pass)
	//hpass, err := encryptWithBcrypt(user.Salt, pass)

	if err != nil {
		return models.User{}, err
	}

	passOk := crypto.SecureCompare(hpass, []byte(passStored))

	if !passOk {
		return models.User{}, ErrWrongPassword
	}

	return user, nil
}

func encryptWithCrypto(salt string, pass string) ([]byte, error) {
	defer TimeTrack(time.Now(), "encryptWithCrypto")
	return crypto.HashPassword(pass, []byte(salt))
}

func encryptWithBcrypt(salt string, pass string) ([]byte, error) {
	defer TimeTrack(time.Now(), "encryptWithBcrypt")
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func (ms *MongoStore) UserUpdate(user models.User) (bool, error) {
	defer TimeTrack(time.Now(), "UserUpdate")
	_, err := ms.UserByEmail(user.Email)

	if err != nil {
		return false, ErrUserNotFound
	}

	err = ms.db.Update(bson.M{"email": user.Email}, &user)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ms *MongoStore) UserByEmail(email string) (models.User, error) {
	defer TimeTrack(time.Now(), "UserByEmail")
	var user models.User

	err := ms.db.Find(bson.M{"email": email}).One(&user)

	if err == nil {

		return user, err
	}
	return user, err
}
