package models

import (
	"gorm.io/gorm"
)

// Struct User
type User struct {
	gorm.Model
	Name       string `form:"name" json: "name" validate:"required"`
	Username   string `form:"username" json: "username" validate:"required"`
	Email      string `form:"email" json: "email" validate:"required"`
	Password   string `form:"password" json: "password" validate:"required"`
	Cart       Cart
	Transaksis []Transaksi
}

// fungsi untuk menambah user baru
func CreateUser(db *gorm.DB, dataUser *User) (err error) {
	err = db.Create(dataUser).Error
	if err != nil {
		return err
	}
	return nil
}

// fungsi untuk memanggil data semua user
func GetUser(db *gorm.DB, dataUser *[]User) (err error) {
	err = db.Find(dataUser).Error
	if err != nil {
		return err
	}
	return nil
}

// fungsi untuk memanggil data user berdasarkan username
func GetUserByUsername(db *gorm.DB, dataUser *User, username string) (err error) {
	err = db.Where("username=?", username).First(dataUser).Error
	if err != nil {
		return err
	}
	return nil
}
