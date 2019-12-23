package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	Username string	`json:"username"`
	Password string	`json:"password"`
}
type UserService struct {
	DB *gorm.DB
}

var secretKey = []byte("my_secret_key_here")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (u UserService) Register(ctx context.Context, username string, password string) (uint, error) {
	//find already exist
	var user User
	if u.DB.Where("username = ?", username).First(&user).RecordNotFound() {

		//create new and save
		user = User{
			Username: username,
			Password: password,
		}
		u.DB.Create(&user)

		return user.ID, nil
	}

	return 0, ErrorString{Message:"Username already exist"}
}
func (u UserService) Login(ctx context.Context, username string, password string) (string, error) {
	var user User
	if u.DB.Where("username = ? AND password = ?", username, password).First(&user).RecordNotFound() {
		return "", ErrorString{Message:"Wrong username or password"}
	}

	var exp = time.Now().Add(30 * time.Minute)
	var claim = Claims{
		Username:       user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:exp.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", ErrorString{Message:"can not create jwt token"}
	}

	return jwtToken, nil
}
func (u UserService) Verify(ctx context.Context, token string) (uint, error) {
	index := strings.Index(token, "Bearer ")
	if index != 0 {
		return 0, ErrorString{Message:"invalid Authentication header"}
	}

	token = strings.ReplaceAll(token, "Bearer ", "")

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !tkn.Valid {
		return 0, ErrorString{Message:"unauthorized"}
	}

	//no need to check exp time, right?
	var user User
	if u.DB.Where("username = ?", claims.Username).First(&user).RecordNotFound() {
		return 0, ErrorString{Message:"username not found"}
	}

	return user.ID, nil
}

// errorString is a trivial implementation of error.
type ErrorString struct {
	Message string
}

func (e ErrorString) Error() string {
	return e.Message
}
