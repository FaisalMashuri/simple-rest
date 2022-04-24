package models

import (
	"net/http"
	"time"

	"echo_rest/db"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	// "github.com/labstack/echo/v4"
	// "golang.org/x/text/date"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type User struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
}

func Signup(nama, password string) (Response, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := User{
		Nama:     nama,
		Password: string(hashPass),
	}
	db := db.GetDBInstance()
	var res Response
	err = db.Create(&user).Error
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "failed"
		res.Data = []Pegawai{}
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "pendaftaran berhasil"
	res.Data = user
	return res, nil
}

func Signin(nama, password string) (ResponseLogin, error) {
	var res ResponseLogin
	db := db.GetDBInstance()
	user := User{}
	// Unfound
	db.FirstOrInit(&user, User{Nama: nama})
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return res, err
	}

	token, err := createJwtToken(string(rune(user.ID)))
	res.Status = http.StatusOK
	res.Message = "login berhasil"
	res.Data = token

	return res, nil

}

func createJwtToken(id string) (string, error) {
	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			Id:        "user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("MySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

// func writeCooke(token string,c echo.Context) error {
// 	cookie := new(http.Cookie)
// 	cookie.Name = "username"
// 	cookie.Value = "jon"
// 	cookie.Expires = time.Now().Add(24 * time.Hour)
// 	c.SetCookie(cookie)
// 	return
// }
