package controllers

import (
	"echo_rest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPegawai(c echo.Context) error {
	result, err := models.ReadAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed",
		})
	}
	return c.JSON(http.StatusOK, result)
}

func GetPegawaiById(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.ReadOneById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed",
		})
	}
	return c.JSON(http.StatusOK, result)
}

func CreatePegawai(c echo.Context) error {
	pegawai := models.Pegawai{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&pegawai)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "can't extract value of json request",
		})
	}
	result, err := models.StorePegawai(pegawai)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, result)

}

func EditPegawai(c echo.Context) error {
	id := c.Param("id")
	pegawai := models.Pegawai{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&pegawai)
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "can't extract value of json request",
		})
	}
	result, err := models.UpdatePegawai(conv_id, pegawai)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, result)
}

func DeletePegawai(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "fail",
		})
	}
	result, err := models.DeletePegawai(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "fail",
		})
	}
	return c.JSON(http.StatusOK, result)

}
