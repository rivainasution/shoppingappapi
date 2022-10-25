package controllers

import (
	"strconv"

	"github.com/rivainasution/shoppingappapi/database"
	"github.com/rivainasution/shoppingappapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type CartApiController struct {
	// Declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitCartAPIController(s *session.Store) *CartApiController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Cart{})

	return &CartApiController{Db: db, store: s}
}

// GET /addtocart/:cartid/product/:productid
func (controller *CartApiController) AddCart(c *fiber.Ctx) error {
	params := c.AllParams()

	intCartId, _ := strconv.Atoi(params["cartid"])
	intProductId, _ := strconv.Atoi(params["productid"])

	var cart models.Cart
	var product models.Product

	err := models.GetProductById(controller.Db, &product, intProductId)
	if err != nil {
		return c.SendStatus(500)
	}

	errs := models.GetCartById(controller.Db, &cart, intCartId)
	if errs != nil {
		return c.SendStatus(500)
	}

	errss := models.AddCart(controller.Db, &cart, &product)
	if errss != nil {
		return c.SendStatus(500)
	}

	if intCartId != 1 {
		return c.SendStatus(400)
	}

	return c.SendString("success")
}

// Get: /shoppingcart/:cartid
func (controller *CartApiController) GetCart(c *fiber.Ctx) error {
	params := c.AllParams()

	intCartId, _ := strconv.Atoi(params["cartid"])

	var dataCart models.Cart
	err := models.GetCart(controller.Db, &dataCart, intCartId)

	if err != nil {
		return c.SendStatus(400)
	}

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}

	userId := sess.Get("userId")

	if intCartId != 1 {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"Title":    "Detail Product",
		"Products": dataCart.Products,
		"UserId":   userId,
	})
}
