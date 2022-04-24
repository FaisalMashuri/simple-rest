package controllers

import (
	"echo_rest/models"
	"net/http"
	"time"

	// "strconv"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	result, err := models.Signup(user.Nama, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "fail",
		})
	}
	// defer c.Request().Body.Close()

	// err := json.NewDecoder(c.Request().Body).Decode(&user)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{
	// 		"message": "can't extract value of json request",
	// 	})
	// }
	// result, err := models.Signup(user)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{
	// 		"message": "fail",
	// 	})
	// }
	return c.JSON(http.StatusOK, result)

}

func Login(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	result, err := models.Signin(user.Nama, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "fail",
		})
	}
	JwtCookie := new(http.Cookie)
	JwtCookie.Name = "JwtCookie"
	JwtCookie.Value = result.Data
	JwtCookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(JwtCookie)
	return c.JSON(http.StatusOK, result)
}

func Logout(c echo.Context) error {
	cookie, err := c.Cookie("JwtCookie")
	if err != nil {
		return c.String(http.StatusInternalServerError, "gagal")
	}

	cookie.Name = "JwtCookie"
	cookie.Value = ""
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Anda berhasil Logout",
	})
}
