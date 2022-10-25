package controllers

import (
	"fmt"
	"strconv"

	"github.com/rivainasution/shoppingappapi/database"
	"github.com/rivainasution/shoppingappapi/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductAPIController struct {
	// declare variables
	Db *gorm.DB
}

func InitProductAPIController() *ProductAPIController {
	db := database.InitDb()

	db.AutoMigrate(&models.Product{})

	return &ProductAPIController{Db: db}
}

// Get: /api/products
func (controller *ProductAPIController) GetProduct(c *fiber.Ctx) error {
	var dataProduct []models.Product

	err := models.GetProducts(controller.Db, &dataProduct)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(dataProduct)
}

// Post: /api/products
func (controller *ProductAPIController) CreateProduct(c *fiber.Ctx) error {
	var dataProduct models.Product

	if err := c.BodyParser(&dataProduct); err != nil {
		return c.SendStatus(400)
	}

	if form, err := c.MultipartForm(); err == nil {

		files := form.File["image"]

		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			dataProduct.Image = fmt.Sprintf("public/upload/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/upload/%s", file.Filename)); err != nil {
				return err
			}
		}
	}
	err := models.CreateProduct(controller.Db, &dataProduct)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.JSON(dataProduct)
}

// Get: /products/detail/:id
func (controller *ProductAPIController) GetDetailProduct(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, _ := strconv.Atoi(idParams)

	var dataProduct models.Product

	err := models.GetProductById(controller.Db, &dataProduct, id)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(dataProduct)
}

// Put: /products/:id
func (controller *ProductAPIController) UpdateProduct(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, _ := strconv.Atoi(idParams)

	var dataProduct models.Product
	err := models.GetProductById(controller.Db, &dataProduct, id)
	if err != nil {
		return c.SendStatus(500)
	}
	var updateProduct models.Product

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.SendStatus(400)
	}
	dataProduct.Name = updateProduct.Name
	dataProduct.Shop = updateProduct.Shop
	dataProduct.Price = updateProduct.Price

	if form, err := c.MultipartForm(); err == nil {

		files := form.File["image"]

		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			dataProduct.Image = fmt.Sprintf("public/upload/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/upload/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	models.UpdateProduct(controller.Db, &dataProduct)

	return c.JSON(dataProduct)
}

// Delete: /products/:id
func (controller *ProductAPIController) DeleteProduct(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, _ := strconv.Atoi(idParams)

	var dataProduct models.Product
	models.DeleteProduct(controller.Db, &dataProduct, id)

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}
