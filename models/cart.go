package models

import (
	"gorm.io/gorm"
)

// Struct Cart
type Cart struct {
	gorm.Model
	UserID   uint
	Products []*Product `gorm:"many2many:cart_products;"`
}

// fungsi untuk membuat cart baru
func CreateCart(db *gorm.DB, dataCart *Cart, userId uint) (err error) {
	dataCart.UserID = userId
	err = db.Create(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// fungsi untuk menambahkan product ke cart
func AddCart(db *gorm.DB, dataCart *Cart, product *Product) (err error) {
	dataCart.Products = append(dataCart.Products, product)
	err = db.Save(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// fungsi untuk memanggil semua data cart
func GetCart(db *gorm.DB, dataCart *Cart, id int) (err error) {
	err = db.Where("user_id=?", id).Preload("Products").Find(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// fungsi untuk memanggil data cart berdasarkan id
func GetCartById(db *gorm.DB, dataCart *Cart, id int) (err error) {
	err = db.Where("user_id=?", id).First(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// fungsi untuk menghapus cart jika cart sudah di checkout
func UpdateCart(db *gorm.DB, dataProduct []*Product, dataCart *Cart, userId uint) (err error) {
	db.Model(&dataCart).Association("Products").Delete(dataProduct)

	return nil
}
