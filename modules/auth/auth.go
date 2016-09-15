package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"../../models"
	"./crypto"
	"./jwt"
	"./store"
)

var invalidated map[string]int = map[string]int{}

func ActivateUser(token string, email string, bs *store.BoltStore, options jwt.Options) (bool, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	user, err := bs.UserByID(userId)
	if err == nil {
		user, err = bs.UserByID(email)
		if err == nil {
			if user.Active {
				return false, errors.New("User already activated!")
			} else {
				user.Active = true
				user.LastUpdated = time.Now().Unix()
				success, err := bs.UserUpdate(user)
				if success {
					return true, nil
				} else {
					return false, err
				}
			}
		} else {
			return false, err
		}
	}
	return false, err
}

func Authenticate(token string, options jwt.Options) (bool, error) {
	if token == "" {
		return false, errors.New("Error no token was provided")
	}
	jti, _ := jwt.AuthenticateMethodJTI(token, options.PublicKey)
	if invalidated[jti] == 1 {
		return false, errors.New("Token has been revoked.")
	} else {
		_, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}

func AuthenticateUserId(token string, bs *store.BoltStore, options jwt.Options) (string, error) {
	if token == "" {
		return "", errors.New("Error no token was provided")
	}
	jti, _ := jwt.AuthenticateMethodJTI(token, options.PublicKey)
	if invalidated[jti] == 1 {
		return "", errors.New("Token has been revoked.")
	} else {
		userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
		if err != nil {
			return "", err
		}
		return userId, nil
	}
}

func ConfirmUserAccount(token string, email string, bs *store.BoltStore, options jwt.Options) (bool, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	user, err := bs.UserByID(userId)
	if err == nil {
		user, err = bs.UserByID(email)
		if err == nil {
			if user.Confirmed {
				return false, errors.New("User account has already been confirmed!")
			} else {
				user.Confirmed = true
				user.LastUpdated = time.Now().Unix()
				success, err := bs.UserUpdate(user)
				if success {
					return true, nil
				} else {
					return false, err
				}
			}
		} else {
			return false, err
		}
	}
	return false, err
}

func DeactivateUser(token string, email string, text string, bs *store.BoltStore, options jwt.Options) (bool, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	user, err := bs.UserByID(userId)
	if err == nil {
		user, err = bs.UserByID(email)
		if err == nil {
			if !user.Active {
				return false, errors.New("User has already been deactivated!")
			} else {
				user.Active = false
				user.DisabledText = text
				user.LastUpdated = time.Now().Unix()
				success, err := bs.UserUpdate(user)
				if success {
					return true, nil
				} else {
					return false, err
				}
			}
		} else {
			return false, err
		}
	}
	return false, err
}

func GetUser(token string, bs *store.BoltStore, options jwt.Options) (models.User, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	if err == nil {
		user, err := bs.UserByID(userId)
		if err == nil {
			return user, nil
		} else {
			return models.User{}, err
		}
	}
	return models.User{}, err
}

func GetUsers(token string, bs *store.BoltStore, options jwt.Options) ([]models.User, error) {
	_, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	if err == nil {
		tmp, err2 := bs.GetAllUsers()
		if err2 == nil {
			return tmp, nil
		}
		return []models.User{}, err
	}
	return []models.User{}, err
}

func Impersonate(userId string, bs *store.BoltStore, options jwt.Options) (string, error) {
	user, err := bs.UserByID(userId)

	token, err := jwt.GenerateJWTToken(user.Id, user.Scopes, options)
	if err != nil {
		return "", errors.New("Error while Signing Token :S")
	}

	return token, nil
}

func ValidateBolt() {

}

func ValidateLedis() {

}

func ValidateMongo() {

}

func LoginBolt(email string, password string, bs *store.BoltStore, options jwt.Options) (string, error) {
	user, err := bs.LoginWithUser(email, password)
	if err != nil {
		return "", err
	}
	user, token, err := login(user, options)
	if err == nil {
		bs.UserUpdate(user)
	}
	return token, err
}

func LoginLedis(email string, password string, ls *store.LedisStore, options jwt.Options) (string, error) {
	user, err := ls.LoginWithUser(email, password)
	if err != nil {
		return "", err
	}
	user, token, err := login(user, options)
	if err == nil {
		ls.UserUpdate(user)
	}
	return token, err
}

func LoginMongo(email string, password string, ms *store.MongoStore, options jwt.Options) (string, error) {
	user, err := ms.LoginWithUser(email, password)
	if err != nil {
		return "", err
	}
	user, token, err := login(user, options)

	if err == nil {
		ms.UserUpdate(user)
	}
	return token, err
}

func login(user models.User, options jwt.Options) (models.User, string, error) {
	if user.Active && user.Confirmed {
		user.LastLogin = time.Now().Unix()
		token, err := jwt.GenerateJWTToken(user.Id, user.Scopes, options)

		if err != nil {
			return user, "", errors.New("Error while Signing Token :S")
		} else {
			return user, token, nil
		}

		//check confirmed
		//return disabled text

	} else {
		if user.Active && !user.Confirmed {
			return user, "", errors.New("Your account still needs to be confirmed.  Please check your email and click confirm account.")
		}
		if !user.Active && user.Confirmed {
			return user, user.DisabledText, errors.New("User has been disabled by your system administrator.  Please contact them to proceed.")
		}
		return user, "", errors.New("Your account is blocked.  Please contact your system administrator")
	}
	return user, "", errors.New("User is not active.  Please contact your system administrator")

}

func RefreshToken(token string, bs *store.BoltStore, options jwt.Options) (string, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	if err != nil {
		return "", err
	}

	b, err := bs.UserByID(userId)

	token2, err := jwt.GenerateJWTToken(userId, b.Scopes, options)
	if err != nil {
		return "", err
	}

	jtoken, err := json.Marshal(map[string]string{"token": token2})
	if err != nil {
		return "", err
	}

	fmt.Println(jtoken)

	return token2, nil
}

func Signin(email string, password string, scopes []string, bs *store.BoltStore) (string, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	userId, err := bs.Signin(email, password, scopes)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func TokenRemainTime(token string, options jwt.Options) (string, error) {
	tim, err := jwt.TimeRemainOnTokenRaw(token, options.PublicKey)
	if err != nil {
		return "", err
	}
	return tim, nil
}

func TokenScopes(token string, options jwt.Options) ([]string, error) {
	tim, err := jwt.ScopesOnTokenRaw(token, options.PublicKey)
	if err != nil {
		return []string{}, err
	}
	return tim, nil
}

func UpdatePassword(token string, password string, bs *store.BoltStore, options jwt.Options) (bool, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	if err == nil {
		user, err := bs.UserByID(userId)
		if err == nil {
			hpass, err := crypto.HashPassword(password, []byte(user.Salt))
			if err == nil {
				user.Password = string(hpass)
				user.PasswordSet = time.Now().Unix()
				user.LastUpdated = time.Now().Unix()
				success, err := bs.UserUpdate(user)
				if success {
					return true, nil
				} else {
					return false, err
				}
			} else {
				return false, err
			}
		} else {
			return false, err
		}
	}
	return false, err
}

func UpdateUserDOB(token string, dob int64, bs *store.BoltStore, options jwt.Options) (bool, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	if err == nil {
		user, err := bs.UserByID(userId)
		if err == nil {
			if dob != 0 && dob != user.Dob {
				user.Dob = dob
				user.LastUpdated = time.Now().Unix()
				success, err := bs.UserUpdate(user)
				if success {
					return true, nil
				} else {
					return false, err
				}
			} else {
				return false, err
			}
		} else {
			return false, err
		}
	}
	return false, err
}

func UpdateUserFirstname(token string, firstname string, bs *store.BoltStore, options jwt.Options) (bool, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	if err == nil {
		user, err := bs.UserByID(userId)
		if err == nil {
			if firstname != "" && firstname != user.Firstname {
				user.Firstname = firstname
				user.LastUpdated = time.Now().Unix()
				success, err := bs.UserUpdate(user)
				if success {
					return true, nil
				} else {
					return false, err
				}
			} else {
				return false, err
			}
		} else {
			return false, err
		}
	}
	return false, err
}

func UpdateUserLanguage(token string, language string, bs *store.BoltStore, options jwt.Options) (bool, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	if err == nil {
		user, err := bs.UserByID(userId)
		if err == nil {
			if language != "" && language != user.Language {
				user.Language = language
				user.LastUpdated = time.Now().Unix()
				success, err := bs.UserUpdate(user)
				if success {
					return true, nil
				} else {
					return false, err
				}
			} else {
				return false, err
			}
		} else {
			return false, err
		}
	}
	return false, err
}

func UpdateUserSurname(token string, surname string, bs *store.BoltStore, options jwt.Options) (bool, error) {
	userId, _, err := jwt.ValidateTokenRaw(token, options.PublicKey)
	if err == nil {
		user, err := bs.UserByID(userId)
		if err == nil {
			if surname != "" && surname != user.Surname {
				user.Surname = surname
				user.LastUpdated = time.Now().Unix()
				success, err := bs.UserUpdate(user)
				if success {
					return true, nil
				} else {
					return false, err
				}
			} else {
				return false, err
			}
		} else {
			return false, err
		}
	}
	return false, err
}

func ValidateScope(token string, scope string, options jwt.Options) (bool, error) {
	tim, err := jwt.ScopesOnTokenRaw(token, options.PublicKey)
	if err != nil {
		return false, err
	}
	for i := 0; i < len(tim); i++ {
		if scope == tim[i] {
			return true, nil
		}
	}
	return false, nil
}
