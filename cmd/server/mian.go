package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/database"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
	"awesomeProject/transport/http"
	_ "gorm.io/gorm"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	// 假设配置文件在项目根目录的/configs/config.yaml
	cfg, err := config.Load("./configs/config.dev.yaml")
	if err != nil {
		log.Fatalf("FATAL: Failed to load config: %v", err)
	}

	// 2. 初始化GORM数据库连接
	// 我们将GORM的初始化逻辑也封装到了database包中，使main.go更整洁
	db, err := database.NewGormConnection(cfg.Database)
	if err != nil {
		log.Fatalf("FATAL: Failed to connect to database: %v", err)
	}
	log.Println("Database connection established successfully.")

	// 可以在这里运行数据库迁移
	// db.AutoMigrate(&models.Product{})
	// log.Println("Database migration completed.")

	// 3. 依赖注入：将所有组件连接起来
	// 数据流向: Handler -> Service -> Repository -> Database
	// a. 创建 Repository 实例，它依赖 GORM 的数据库连接(db)
	productRepo := repository.NewProductRepository(db)

	// b. 创建 Service 实例，它依赖 Repository 层的接口(productRepo)
	productService := service.NewProductService(productRepo)

	// c. 创建 Handler 实例，它依赖 Service 层的接口(productService)
	productHandler := http.NewProductHandler(productService)

	// 4. 初始化 Gin 路由引擎
	router := gin.Default()

	// 5. 注册产品相关的路由
	// 创建一个API分组，方便管理版本，例如 /api/v1
	apiV1 := router.Group("/api/v1")
	{
		products := apiV1.Group("/products")
		{
			products.POST("", productHandler.CreateProduct)       // 创建产品
			products.GET("", productHandler.GetAllProducts)       // 获取所有产品
			products.GET("/:id", productHandler.GetProduct)       // 获取单个产品
			products.PUT("/:id", productHandler.UpdateProduct)    // 更新产品
			products.DELETE("/:id", productHandler.DeleteProduct) // 删除产品
		}
	}

	// 6. 启动服务器
	log.Println("Starting server on port :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}
