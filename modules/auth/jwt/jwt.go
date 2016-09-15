package jwt

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"../crypto"
	jwt "gopkg.in/dgrijalva/jwt-go.v2"
)

var (
	ErrTokenExpired    = errors.New("Token Expired, get a new one")
	ErrTokenValidation = errors.New("JWT Token ValidationError")
	ErrTokenParse      = errors.New("JWT Token Error Parsing the token or empty token")
	ErrTokenInvalid    = errors.New("JWT Token is not Valid")

	logOn = true
)

type Options struct {
	SigningMethod string
	PublicKey     string
	PrivateKey    string
	Expiration    time.Duration
}

// Generates a JSON Web Token given an userId (typically an id or an email), and the JWT options
// to set SigningMethod and the keys you can check
// http://github.com/dgrijalva/jwt-go
//
// In case you use an symmetric-key algorithm set PublicKey and PrivateKey equal to the SecretKey ,
func GenerateJWTToken(userId string, scope []string, op Options) (string, error) {
	t := jwt.New(jwt.GetSigningMethod(op.SigningMethod))

	now := time.Now()
	// set claims
	t.Claims["iat"] = now.Unix()
	t.Claims["exp"] = now.Add(op.Expiration).Unix()
	t.Claims["sub"] = userId
	t.Claims["jti"] = crypto.GenerateRandomKey(32)
	t.Claims["scopes"] = scope

	tokenString, err := t.SignedString([]byte(op.PrivateKey))
	if err != nil {
		logError("ERROR: GenerateJWTToken: %v\n", err)
	}
	return tokenString, err

}

func TokenJTI(r *http.Request, publicKey string) (string, error) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return "", ErrTokenParse
	}

	if !token.Valid {
		return "", ErrTokenInvalid
	}

	// otherwise is a valid token
	JTI := token.Claims["jti"].(string)
	return JTI, nil

}

func ScopeToken(r *http.Request, publicKey string) (string, error) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return "", ErrTokenParse
	}

	if !token.Valid {
		return "", ErrTokenInvalid
	}

	// otherwise is a valid token
	scope := token.Claims["scopes"].(string)
	return scope, nil

}

func TimeRemainOnToken(r *http.Request, publicKey string) (string, error) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return "", ErrTokenParse
	}
	if !token.Valid {

		return "", ErrTokenInvalid
	}

	a := strconv.FormatFloat(token.Claims["exp"].(float64), 'f', 0, 64)
	b, _ := strconv.ParseUint(a, 10, 64)
	c := strconv.FormatInt(int64(b)-int64(time.Now().Unix()), 10)

	return c, nil
}

func RevokeToken(r *http.Request, publicKey string) (bool, error) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return false, ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return false, ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return false, ErrTokenParse
	}

	if !token.Valid {
		return false, ErrTokenInvalid
	}

	// otherwise is a valid token
	token.Valid = false

	return true, nil
}

// Returns the userId, token (base64 encoded), error
func ValidateToken(r *http.Request, publicKey string) (string, string, error) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return "", "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return "", "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return "", "", ErrTokenParse
	}

	if !token.Valid {
		return "", "", ErrTokenInvalid
	}

	// otherwise is a valid token
	userId := token.Claims["sub"].(string)

	return userId, token.Raw, nil

}

func logError(format string, err interface{}) {
	if logOn && err != nil {
		log.Printf(format, err)
	}
}

func ValidateTokenRaw(toke string, publicKey string) (string, string, error) {
	token, err := jwt.Parse(toke, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return "", "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return "", "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return "", "", ErrTokenParse
	}

	if !token.Valid {
		return "", "", ErrTokenInvalid
	}

	// otherwise is a valid token
	userId := token.Claims["sub"].(string)

	return userId, token.Raw, nil

}

func AuthenticateMethodJTI(toke string, publicKey string) (string, error) {
	token, err := jwt.Parse(toke, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return "", ErrTokenParse
	}

	if !token.Valid {
		return "", ErrTokenInvalid
	}

	// otherwise is a valid token
	JTI := token.Claims["jti"].(string)
	return JTI, nil
}

func AuthenticateMethodUser(toke string, publicKey string) (string, error) {
	token, err := jwt.Parse(toke, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return "", ErrTokenParse
	}

	if !token.Valid {
		return "", ErrTokenInvalid
	}

	// otherwise is a valid token
	user := token.Claims["sub"].(string)
	return user, nil
}

func TimeRemainOnTokenRaw(toke string, publicKey string) (string, error) {
	token, err := jwt.Parse(toke, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return "", ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return "", ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return "", ErrTokenParse
	}
	if !token.Valid {

		return "", ErrTokenInvalid
	}

	a := strconv.FormatFloat(token.Claims["exp"].(float64), 'f', 0, 64)
	b, _ := strconv.ParseUint(a, 10, 64)
	c := strconv.FormatInt(int64(b)-int64(time.Now().Unix()), 10)

	return c, nil
}

func ScopesOnTokenRaw(toke string, publicKey string) ([]string, error) {
	token, err := jwt.Parse(toke, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				logError("ERROR: JWT Token Expired: %+v\n", vErr.Errors)
				return []string{}, ErrTokenExpired
			default:
				logError("ERROR: JWT Token ValidationError: %+v\n", vErr.Errors)
				return []string{}, ErrTokenValidation
			}
		}
		logError("ERROR: Token parse error: %v\n", err)
		return []string{}, ErrTokenParse
	}
	if !token.Valid {
		return []string{}, ErrTokenInvalid
	}

	scopes := getStringSlice(token.Claims["scopes"])
	return scopes, nil
}

func getStringSlice(t interface{}) []string {
	scopes := []string{}
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			scopes = append(scopes, s.Index(i).Interface().(string))
		}
	}
	return scopes
}
