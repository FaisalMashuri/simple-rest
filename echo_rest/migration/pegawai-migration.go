package migration

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Pegawai struct {
	gorm.Model
	Nama    string
	Alamat  string
	Telepon string
}
