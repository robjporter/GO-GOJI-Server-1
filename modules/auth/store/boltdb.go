package store

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"../../../models"
	"../crypto"
	"github.com/boltdb/bolt"
)

type BoltStore struct {
	db     *bolt.DB
	bucket []byte
}

func NewBoltStore(db *bolt.DB, userBucket string) (*BoltStore, error) {
	bucket := []byte(userBucket)

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return fmt.Errorf("Creating bucket: %s", err)
		}
		return nil
	})
	return &BoltStore{db: db, bucket: bucket}, err
}

func (bs *BoltStore) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := bs.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket(bs.bucket)
		b.ForEach(func(k, v []byte) error {
			user, err := gobDecode(v)
			if err == nil {
				users = append(users, user)
			}
			return nil
		})
		return err
	})
	if err == nil {
		return users, nil
	} else {
		return users, err
	}
}

func (bs *BoltStore) UserUpdate(user models.User) (bool, error) {
	_, err := bs.UserByEmail(user.Email)
	if err != nil {
		return false, ErrUserNotFound
	}

	err = bs.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(bs.bucket)

		g, err := gobEncode(user)
		if err != nil {
			return err
		}
		err = b.Put([]byte(user.Email), g)
		return err
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

func (bs *BoltStore) UserByID(id string) (models.User, error) {
	var user models.User
	err := bs.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket(bs.bucket)
		gobUser := b.Get([]byte(id))
		if gobUser == nil {
			return ErrUserNotFound
		}
		user, err = gobDecode(gobUser)
		return err
	})
	return user, err
}

func (bs *BoltStore) UserByEmail(email string) (models.User, error) {
	var user models.User
	err := bs.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket(bs.bucket)
		gobUser := b.Get([]byte(email))
		if gobUser == nil {
			return ErrUserNotFound
		}
		user, err = gobDecode(gobUser)
		return err
	})
	return user, err
}

func (bs *BoltStore) Signin(email, pass string, scope []string) (string, error) {
	// check if the user exists
	fmt.Println("HERE")
	_, err := bs.UserByEmail(email)
	fmt.Println("HERE")
	if err == nil {
		return "", ErrEmailDuplication
	}

	err = bs.db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(bs.bucket)

		// email is going to be the Id of the user
		user, err := NewUser(email, email, pass, scope)
		if err != nil {
			return err
		}
		g, err := gobEncode(user)
		if err != nil {
			return err
		}
		err = b.Put([]byte(email), g)
		return err
	})

	if err != nil {
		return "", err
	}
	return email, nil

}

func (bs *BoltStore) Login(email, pass string) (string, error) {
	user, err := bs.UserByEmail(email)
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

func (bs *BoltStore) LoginWithUser(email, pass string) (models.User, error) {
	user, err := bs.UserByEmail(email)
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

func gobEncode(user models.User) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	err := enc.Encode(user)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func gobDecode(b []byte) (models.User, error) {
	reader := bytes.NewReader(b)
	dec := gob.NewDecoder(reader)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
