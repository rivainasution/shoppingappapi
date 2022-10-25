package models

import (
	"gorm.io/gorm"
)

// Struct Transaksi
type Transaksi struct {
	gorm.Model
	Id       int `form:"id" json: "id" validate:"required"`
	UserID   uint
	Products []*Product `gorm:"many2many:transaksi_products;"`
}

// Fungsi untuk membuat transaksi baru
func CreateTransaksi(db *gorm.DB, dataTransaksi *Transaksi, userId uint, dataProduct []*Product) (err error) {
	dataTransaksi.UserID = userId
	dataTransaksi.Products = dataProduct

	err = db.Create(dataTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk memanggil semua data transaksi
func GetTransaksi(db *gorm.DB, dataTransaksi *Transaksi, id int) (err error) {
	err = db.Where("id=?", id).Preload("Products").Find(dataTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk memanggil data transaksi berdasarkan id
func GetTransaksiById(db *gorm.DB, dataTransaksi *[]Transaksi, id int) (err error) {
	err = db.Where("user_id=?", id).Find(dataTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk menambahkan produk yang di cart ke transaksi
func AddTransaksi(db *gorm.DB, dataCart *Cart, dataProduct *Product) (err error) {
	dataCart.Products = append(dataCart.Products, dataProduct)
	err = db.Save(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}
