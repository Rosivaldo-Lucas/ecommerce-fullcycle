package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rosivaldolucas/ecommerce/goapi/internal/database"
	"github.com/rosivaldolucas/ecommerce/goapi/internal/service"
	"github.com/rosivaldolucas/ecommerce/goapi/internal/webserver"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/catalog")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(*categoryService)
	webProductHandler := webserver.NewWebProductHandler(*productService)

	c := chi.NewRouter()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/categories", webCategoryHandler.GetCategories)
	c.Get("/categories/{id}", webCategoryHandler.GetCategory)
	c.Post("/categories", webCategoryHandler.CreateCategory)

	c.Get("/products", webProductHandler.GetProducts)
	c.Get("/products/categories/{categoryID}", webProductHandler.GetProductByCategoryID)
	c.Get("/products/{id}", webProductHandler.GetProduct)
	c.Post("/products", webProductHandler.CreateCategory)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
