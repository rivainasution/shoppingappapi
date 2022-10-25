package controllers

import (
	"github.com/rivainasution/shoppingappapi/database"
	"github.com/rivainasution/shoppingappapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

// Struct AuthAPIController
type AuthAPIController struct {
	// Declare variables
	Db    *gorm.DB
	store *session.Store
}

// init
func InitAuthAPIController(s *session.Store) *AuthAPIController {
	db := database.InitDb()

	db.AutoMigrate(&models.User{})

	return &AuthAPIController{Db: db, store: s}
}

// Post: /api/login
func (controller *AuthAPIController) Login(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)

	if err != nil {
		panic(err)
	}

	var dataUser models.User
	var dataForm LoginForm

	if err := c.BodyParser(&dataForm); err != nil {
		return c.JSON(fiber.Map{
			"msg": "Form tidak boleh kosong",
		})
	}

	//Mencari username apakah sudah terdaptar
	errs := models.GetUserByUsername(controller.Db, &dataUser, dataForm.Username)
	if errs != nil {
		return c.JSON(fiber.Map{
			"msg": "Form tidak boleh kosong",
		})
	}

	//Hashing
	hash := bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(dataForm.Password))
	if hash == nil {
		sess.Set("username", dataUser.Username)
		sess.Set("userId", dataUser.ID)
		sess.Save()

		return c.JSON(fiber.Map{
			"msg": "Login berhasil",
		})
	}

	return c.JSON(fiber.Map{
		"msg": "Username atau password salah",
	})
}

// Get: /api/listregister
func (controller *AuthAPIController) GetRegister(c *fiber.Ctx) error {
	var dataUser []models.User
	err := models.GetUser(controller.Db, &dataUser)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.JSON(dataUser)
}

// Post: /api/register
func (controller *AuthAPIController) Register(c *fiber.Ctx) error {
	var dataUser models.User
	var dataCart models.Cart

	if err := c.BodyParser(&dataUser); err != nil {
		return c.SendStatus(400)
	}

	//hashing password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(dataUser.Password), 10)
	hash := string(bytes)

	dataUser.Password = hash

	err := models.CreateUser(controller.Db, &dataUser)
	if err != nil {
		return c.SendStatus(500)
	}

	errs := models.GetUserByUsername(controller.Db, &dataUser, dataUser.Username)
	if errs != nil {
		return c.SendStatus(500)
	}

	errCart := models.CreateCart(controller.Db, &dataCart, dataUser.ID)
	if errCart != nil {
		return c.JSON(dataUser)
	}

	return c.JSON(dataUser)
}

// Get: /api/logout
func (controller *AuthAPIController) Logout(c *fiber.Ctx) error {

	sess, err := controller.store.Get(c)

	if err != nil {
		panic(err)
	}

	sess.Destroy()

	return c.JSON(fiber.Map{
		"msg": "Logout Berhasil",
	})
}
