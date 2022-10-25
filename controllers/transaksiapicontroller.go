package controllers

import (
	"strconv"

	"github.com/rivainasution/shoppingappapi/database"
	"github.com/rivainasution/shoppingappapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type TransaksiAPIController struct {
	// Declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitTransaksiAPIController(s *session.Store) *TransaksiAPIController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Transaksi{})

	return &TransaksiAPIController{Db: db, store: s}
}

// GET /checkout/:userid
func (controller *TransaksiAPIController) AddTransaksi(c *fiber.Ctx) error {
	params := c.AllParams()
	intUserId, _ := strconv.Atoi(params["userid"])

	var dataTransaksi models.Transaksi
	var dataCart models.Cart

	// Find the product first,
	err := models.GetCart(controller.Db, &dataCart, intUserId)
	if err != nil {
		return c.SendStatus(500)
	}

	errs := models.CreateTransaksi(controller.Db, &dataTransaksi, uint(intUserId), dataCart.Products)
	if errs != nil {
		return c.SendStatus(500)
	}

	// Delete products in cart
	errss := models.UpdateCart(controller.Db, dataCart.Products, &dataCart, uint(intUserId))

	if errss != nil {
		return c.SendStatus(500)
	}

	if intUserId != 1 {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"Title": "History Transaksi",
		"Tid":   intUserId,
	})
}

// GET /historytransaksi/:userid
func (controller *TransaksiAPIController) GetTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() 

	intUserId, _ := strconv.Atoi(params["userid"])

	var dataTransaksi []models.Transaksi
	err := models.GetTransaksiById(controller.Db, &dataTransaksi, intUserId)
	if err != nil {
		return c.SendStatus(500) 
	}

	if intUserId != 1 {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"Title":      "History Transaksi",
		"Transaksis": dataTransaksi,
	})

}

// GET /history/detail/:transaksiid
func (controller *TransaksiAPIController) DetailTransaksi(c *fiber.Ctx) error {
	params := c.AllParams()

	intTransaksiId, _ := strconv.Atoi(params["transaksiid"])

	var dataTransaksi models.Transaksi
	err := models.GetTransaksi(controller.Db, &dataTransaksi, intTransaksiId)
	if err != nil {
		return c.SendStatus(500) 

	return c.JSON(fiber.Map{
		"Title":    "History Transaksi",
		"Products": dataTransaksi.Products,
	})
}
