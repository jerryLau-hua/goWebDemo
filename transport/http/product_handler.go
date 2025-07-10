package http

import (
	"awesomeProject/internal/service" // 确保这是你的服务包路径
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ProductHandler 负责处理产品相关的HTTP请求
type ProductHandler struct {
	productService service.ProductService // 它依赖于 ProductService 接口
}

// NewProductHandler 是 ProductHandler 的构造函数
func NewProductHandler(svc service.ProductService) *ProductHandler {
	return &ProductHandler{productService: svc}
}

// CreateProduct godoc
// @Summary      创建一个新产品
// @Description  根据传入的JSON数据创建一个新产品
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      CreateProductRequest  true  "创建产品请求"
// @Success      201      {object}  models.Product
// @Failure      400      {object}  gin.H
// @Failure      500      {object}  gin.H
// @Router       /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	// 定义一个临时的结构体来绑定请求的JSON body
	var req struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"gt=0"`
		Stock int     `json:"stock" binding:"gte=0"`
	}

	// 解析并验证JSON请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: " + err.Error()})
		return
	}

	// 调用Service层来创建产品
	product, err := h.productService.CreateProduct(c.Request.Context(), req.Name, req.Price, req.Stock)
	if err != nil {
		// 根据Service层返回的错误类型，可以返回更具体的HTTP状态码
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回201 Created状态码和创建成功的产品信息
	c.JSON(http.StatusCreated, product)
}

// GetProduct godoc
// @Summary      获取单个产品
// @Description  根据产品ID获取产品详情
// @Tags         Products
// @Produce      json
// @Param        id   path      int  true  "产品ID"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  gin.H
// @Failure      404  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	// 从URL路径中获取ID参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// 调用Service层获取产品
	product, err := h.productService.GetProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 如果Service层返回nil，说明产品不存在
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetAllProducts godoc
// @Summary      获取所有产品列表
// @Description  获取数据库中所有产品的列表
// @Tags         Products
// @Produce      json
// @Success      200  {array}   models.Product
// @Failure      500  {object}  gin.H
// @Router       /products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary      更新一个产品
// @Description  根据ID和传入的JSON数据更新一个已存在的产品
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id       path      int                   true  "产品ID"
// @Param        product  body      UpdateProductRequest  true  "更新产品请求"
// @Success      200      {object}  models.Product
// @Failure      400      {object}  gin.H
// @Failure      404      {object}  gin.H
// @Failure      500      {object}  gin.H
// @Router       /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var req struct {
		Name  string  `json:"name" binding:"required"`
		Price float64 `json:"price" binding:"gt=0"`
		Stock int     `json:"stock" binding:"gte=0"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: " + err.Error()})
		return
	}

	product, err := h.productService.UpdateProduct(c.Request.Context(), id, req.Name, req.Price, req.Stock)
	if err != nil {
		// 这里可以根据service返回的错误类型判断是404还是500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary      删除一个产品
// @Description  根据ID删除一个产品
// @Tags         Products
// @Produce      json
// @Param        id   path      int  true  "产品ID"
// @Success      204  {object}  nil
// @Failure      500  {object}  gin.H
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.productService.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 对于删除操作，成功后通常返回 204 No Content
	c.Status(http.StatusNoContent)
}
