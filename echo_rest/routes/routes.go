package routes

import (
	"net/http"

	"echo_rest/controllers"

	"echo_rest/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	p := e.Group("/api/v1/pegawai")
	p.GET("", controllers.GetPegawai)
	p.GET("/:id", controllers.GetPegawaiById)
	p.POST("", controllers.CreatePegawai)
	p.PUT("/:id", controllers.EditPegawai)
	p.DELETE("/:id", controllers.DeletePegawai)

	u := e.Group("/api/v1")
	u.POST("/register", controllers.Register)
	u.POST("/login", controllers.Login)
	u.POST("/logout", controllers.Logout)
	gJWT := e.Group("/api/v1")
	middleware.SetJWTMiddleware(gJWT)
	gJWT.GET("/dashboard", AdminDashboard)
	// e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAuthenticated)
	// e.POST("/pegawai", controllers.StorePegawai, middleware.IsAuthenticated)
	// e.PUT("/pegawai", controllers.UpdatePegawai, middleware.IsAuthenticated)
	// e.DELETE("/pegawai", controllers.DeletePegawai, middleware.IsAuthenticated)

	// e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	// e.POST("/login", controllers.CheckLogin)

	// e.GET("/test-struct-validation", controllers.TestStructValidation)
	// e.GET("/test-variable-validation", controllers.TestVariableValidation)

	return e
}

func AdminDashboard(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"page": "admin dashboard",
	})
}
