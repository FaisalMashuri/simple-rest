package models

import (
	"net/http"

	"echo_rest/db"
	// "github.com/labstack/echo/v4"
	// "golang.org/x/text/date"
)

type Pegawai struct {
	ID      int    `json:"id"`
	Nama    string `json:"nama" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Telepon string `json:"telepon" validate:"required"`
}

func ReadAll() (Response, error) {
	db := db.GetDBInstance()
	pegawai := []Pegawai{}
	err := db.Find(&pegawai).Error
	var res Response
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "failed"
		res.Data = []Pegawai{}
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = pegawai
	return res, nil
}

func ReadOneById(id int) (Response, error) {
	db := db.GetDBInstance()
	pegawai := Pegawai{}
	err := db.Find(&pegawai, id).Error
	var res Response
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "failed"
		res.Data = []Pegawai{}
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = pegawai
	return res, nil
}

func StorePegawai(pegawai Pegawai) (Response, error) {
	db := db.GetDBInstance()
	// var pegawai = Pegawai{
	// 	Nama:    name,
	// 	Alamat:  alamat,
	// 	Telepon: telpon,
	// }
	var res Response
	err := db.Create(&pegawai).Error
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "failed"
		res.Data = []Pegawai{}
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Data berhasil ditambahkan"
	res.Data = pegawai
	return res, nil
}

func UpdatePegawai(id int, pegawai Pegawai) (Response, error) {
	db := db.GetDBInstance()
	// pegawaiStruct := Pegawai{}
	var res Response
	err := db.Exec("UPDATE pegawais SET nama=?, alamat=?, telepon=? WHERE id=?", pegawai.Nama, pegawai.Alamat, pegawai.Telepon, id).Error
	// user := db.Where("id", id).First(&pegawaiStruct)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "failed"
		res.Data = []Pegawai{}
		return res, err
	}
	// db.Model(&user).Update(pegawai)
	res.Status = http.StatusOK
	res.Message = "Data Berhasil diubah"
	res.Data = pegawai
	return res, nil

}

func DeletePegawai(id int) (Response, error) {
	db := db.GetDBInstance()
	var res Response
	err := db.Delete(&Pegawai{}, id).Error
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "failed"
		res.Data = []Pegawai{}
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Data Berhasil dihapus"
	return res, nil
}
