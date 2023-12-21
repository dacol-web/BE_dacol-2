package src

import (
	"net/http"
	"os"

	"github.com/Hy-Iam-Noval/dacol-2/src/ctrl"
	"github.com/Hy-Iam-Noval/dacol-2/src/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
)

type App = *fiber.App

func Route() App {
	r := fiber.New()
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// config
	r.Use(ctrl.Acceptable, cors.New(cors.Config{
		AllowOrigins: os.Getenv("FE_BASEURL"),
		AllowHeaders: "Content-Type, user",
		AllowMethods: "GET, POST, DELETE",
	}))

	r.Static("/public", "./public")
	r.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./public"),
	}))

	r.Post("/login", ctrl.Login)
	r.Post("/register", ctrl.Register)

	// Group /auth
	auth := r.Group("/auth").Use(helpers.AuthWare)
	{
		auth.Get("/user", ctrl.ParseToken)

		// Post /upload/:path
		// :path path folder where ]file will be put
		auth.Post("/upload/:path", ctrl.FileUpload)

		// Group /auth/product
		auth.Post("/product_add", ctrl.AddProd)
		auth.Post("/product_update", ctrl.UpdateProd)
		auth.Get("/product_all", ctrl.AllProd)

		// Group /auth/product/:id
		prodIdR := auth.Group("/product/:id")
		{
			prodIdR.Get("/", ctrl.FindProd)
			prodIdR.Delete("/delete", ctrl.DeleteProd)
		}

		// Group /auth/selling
		auth.Post("/selling_add", ctrl.AddSelling)
		sellingR := auth.Group("/selling")
		{
			sellingR.Get("/", ctrl.AllSelling)
			sellingR.Get("/:id", ctrl.AllByIDSelling)
		}

	}
	return r
}
