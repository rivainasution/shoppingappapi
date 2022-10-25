package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rivainasution/shoppingappapi/controllers"
)

func main() {
	// session
	store := session.New()

	app := fiber.New()

	// static
	app.Static("/", "./public", fiber.Static{
		Index: "",
	})

	app.Static("/public", "./public")

	// Fungsi Check Login
	CheckLogin := func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		temp := sess.Get("username")
		if temp != nil {
			return c.Next()
		}

		return c.SendString("Belum Login, silahkan login terlebih dahulu")
	}

	// controllers
	authApiController := controllers.InitAuthAPIController(store)
	productsApiController := controllers.InitProductAPIController()
	cartApiController := controllers.InitCartAPIController(store)
	transaksiApiController := controllers.InitTransaksiAPIController(store)

	//API group
	auth := app.Group("/api")
	//auth
	auth.Post("/login", authApiController.Login)
	auth.Get("/listregister", authApiController.GetRegister)
	auth.Post("/register", authApiController.Register)
	auth.Post("/logout", authApiController.Logout)

	//products
	auth.Get("/products", CheckLogin, productsApiController.GetProduct)
	auth.Post("/products", CheckLogin, productsApiController.CreateProduct)
	auth.Get("/products/detail/:id", CheckLogin, productsApiController.GetDetailProduct)
	auth.Put("/products/:id", CheckLogin, productsApiController.UpdateProduct)
	auth.Delete("/products/:id", CheckLogin, productsApiController.DeleteProduct)

	//cart
	auth.Get("/addcart/:cartid/product/:productid", CheckLogin, cartApiController.AddCart)
	auth.Get("/shoppingcart/:cartid", CheckLogin, cartApiController.GetCart)

	//transaksi
	auth.Get("/listtransaksi/:userid", CheckLogin, transaksiApiController.GetTransaksi)
	auth.Get("/detail/:transaksiid", CheckLogin, transaksiApiController.DetailTransaksi)
	auth.Get("/checkout/:userid", CheckLogin, transaksiApiController.AddTransaksi)

	app.Listen(":3000")
}
