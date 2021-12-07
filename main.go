package main

import (
	"backend-simple-pos/app"
	"backend-simple-pos/auth"
	"backend-simple-pos/category_product"
	"backend-simple-pos/customer"
	"backend-simple-pos/handler"
	"backend-simple-pos/penjualan"
	"backend-simple-pos/product"
	"backend-simple-pos/role"
	"backend-simple-pos/supplier"
	"backend-simple-pos/user"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func main() {

	db := app.NewDB()

	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService,authService)
	authMiddleware := app.AuthMiddleware(authService,userService)

	roleRepository := role.NewRepository(db)
	roleService := role.NewService(roleRepository)
	roleHandler := handler.NewRoleHandler(roleService)

	customerRepository := customer.NewRepository(db)
	customerService := customer.NewService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService, authService)

	categoryProductRepository := category_product.NewRepository(db)
	categoryProductService := category_product.NewService(categoryProductRepository)
	categoryProductHandler := handler.NewCategoryProductHandler(categoryProductService, authService)

	supplierRepository := supplier.NewRepository(db)
	supplierService := supplier.NewService(supplierRepository, categoryProductRepository)
	supplierHandler := handler.NewSupplierHandler(supplierService, authService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository, supplierRepository)
	productHandler := handler.NewProductHandler(productService, authService)

	penjualanRepository := penjualan.NewRepository(db)
	penjualanService := penjualan.NewService(penjualanRepository)
	penjualanHandler := handler.NewPenjualanHandler(penjualanService, authService)

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/user/register", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.Login)

	// USER
	api.GET("/user/fetch", authMiddleware, userHandler.FetchUser)
	api.GET("/user/all", authMiddleware, userHandler.Users)
	api.GET("/user/:id", authMiddleware, userHandler.GetUserDetail)
	api.PUT("/user/:id", authMiddleware, userHandler.UpdatedUser)
	api.DELETE("/user", authMiddleware, userHandler.DeleteUser)
	api.POST("/upload-avatar", authMiddleware, userHandler.UploadAvatar)

	//ROLE
	api.GET("/role", authMiddleware, roleHandler.GetRoles)

	//CUSTOMER
	api.POST("/customer", authMiddleware, customerHandler.RegisterCustomer)
	api.PUT("/customer/:id", authMiddleware, customerHandler.UpdateCustomer)
	api.DELETE("/customer", authMiddleware, customerHandler.DeleteCustomer)
	api.GET("/customer", authMiddleware, customerHandler.GetAllCustomer)

	//CATEGORY PRODUCT
	api.GET("/category-product", authMiddleware, categoryProductHandler.GetAllProductCategory)
	api.POST("/category-product", authMiddleware, categoryProductHandler.RegisterProductCategory)
	api.PUT("/category-product/:id", authMiddleware, categoryProductHandler.UpdateProductCategory)
	api.DELETE("/category-product", authMiddleware, categoryProductHandler.DeleteProductCategory)

	//SUPPLIER
	api.GET("/supplier", authMiddleware, supplierHandler.GetAllSupliers)
	api.POST("/supplier", authMiddleware, supplierHandler.RegisterSupplier)
	api.PUT("/supplier/:id", authMiddleware, supplierHandler.UpdateSupplier)
	api.DELETE("/supplier", authMiddleware, supplierHandler.DeleteSupplier)

	//PRODUCT
	api.GET("/product", authMiddleware, productHandler.GetAllProduct)
	api.POST("/product", authMiddleware, productHandler.RegisterProduct)
	api.PUT("/product/:id", authMiddleware, productHandler.UpdateProduct)
	api.DELETE("/product", authMiddleware, productHandler.DeleteProduct)

	//PENJUALAN
	api.GET("/pemesanan", authMiddleware, penjualanHandler.GetAllPemesanan)
	api.POST("/pemesanan", authMiddleware, penjualanHandler.CreatePemesanan)
	api.PUT("/pemesanan/:id", authMiddleware, penjualanHandler.UpdatePemesananDetail)
	api.DELETE("/pemesanan/:id", authMiddleware, penjualanHandler.DeletePemesananDetail)

	router.Run()
}
