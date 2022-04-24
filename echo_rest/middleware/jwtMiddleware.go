package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetJWTMiddleware(g *echo.Group) {
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("MySecret"),
		TokenLookup:   "cookie:JwtCookie",
	}))
}
