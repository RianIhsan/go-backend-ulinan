package handler

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"strconv"
	"ulinan/domain/product"
	"ulinan/domain/product/dto"
	"ulinan/entities"
	"ulinan/helper/cloudinary"
	"ulinan/helper/response"
	"ulinan/helper/validator"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
}

func NewProductHandler(productService product.ProductServiceInterface) product.ProductHandlerInterface {
	return &ProductHandler{productService}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}
	var payload dto.TCreateProductRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	_, err := h.productService.CreateProduct(&payload)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create product: "+err.Error())
	}

	return response.SendStatusCreatedResponse(c, "product created successfully")
}

func (h *ProductHandler) GetProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}
	product, err := h.productService.GetProductById(id)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to get product: "+err.Error())
	}
	return response.SendStatusOkWithDataResponse(c, "success get product", dto.GetProductById(product))
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageConv, _ := strconv.Atoi(strconv.Itoa(page))
	limit := c.Query("limit", "10")
	limitInt, _ := strconv.Atoi(limit)
	search := c.Query("search")
	categoryID, _ := strconv.Atoi(c.Query("category_id", "0"))

	var products []*entities.ProductEntity
	var totalItems int64
	var err error

	if search != "" && categoryID > 0 {
		products, totalItems, err = h.productService.GetProductsByCategoryAndNameProduct(pageConv, limitInt, categoryID, search)
	} else if categoryID > 0 {
		products, totalItems, err = h.productService.GetProductByCategory(pageConv, limitInt, categoryID)
	} else if search != "" {
		products, totalItems, err = h.productService.GetProductByName(pageConv, limitInt, search)
	} else if limit != "" {
		products, totalItems, err = h.productService.GetAllProducts(pageConv, limitInt)
	}

	if err != nil {
		return response.SendStatusInternalServerError(c, "failed to get products: "+err.Error())
	}

	currentPage, totalPages := h.productService.CalculatePaginationValues(pageConv, int(totalItems), limitInt)
	nextPage := h.productService.GetNextPage(currentPage, totalPages)
	prevPage := h.productService.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.GetPaginationProducts(products), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success get products")
}

func (h *ProductHandler) CreateProductImage(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}
	var payload dto.CreateProductImage
	file, err := c.FormFile("image")
	var uploadedURL string
	if err == nil {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.SendStatusInternalServerError(c, "Failed open to open file: "+err.Error())
		}
		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)
		uploadedURL, err = cloudinary.Uploader(fileToUpload)
		if err != nil {
			return response.SendStatusInternalServerError(c, "Failed to upload image: "+err.Error())
		}
	}

	payload.Image = uploadedURL
	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}

	productImage, err := h.productService.CreateProductImage(payload)
	if err != nil {
		return response.SendStatusBadRequest(c, "error upload image: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "product image created successfully", dto.CreateImageProductResponse(productImage))
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}
	var payload dto.TUpdateProductRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	err = h.productService.UpdateProduct(id, &payload)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to update product: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "product updated successfully")
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}
	err = h.productService.DeleteProduct(id)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to delete product: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "product deleted successfully")
}

func (h *ProductHandler) DeleteProductImage(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}
	productId, err := strconv.Atoi(c.Params("productId"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid product id")
	}
	imageId, err := strconv.Atoi(c.Params("imageId"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid image id")
	}
	err = h.productService.DeleteProductImage(productId, imageId)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to delete product image: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "product image deleted successfully")
}

func (h *ProductHandler) GetRandomProducts(c *fiber.Ctx) error {
	products, err := h.productService.GetRandomProducts(8)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to get random products: "+err.Error())
	}
	return response.SendStatusOkWithDataResponse(c, "success get random products", dto.GetPaginationProducts(products))
}

func (h *ProductHandler) GetProductByCategory(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageConv, _ := strconv.Atoi(strconv.Itoa(page))
	limit := c.Query("limit", "10")
	limitInt, _ := strconv.Atoi(limit)
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid category id")
	}
	var products []*entities.ProductEntity
	var totalItems int64
	products, totalItems, err = h.productService.GetProductByCategory(pageConv, limitInt, categoryId)

	if err != nil {
		return response.SendStatusInternalServerError(c, "failed to get categories: "+err.Error())
	}

	currentPage, totalPages := h.productService.CalculatePaginationValues(pageConv, int(totalItems), limitInt)
	nextPage := h.productService.GetNextPage(currentPage, totalPages)
	prevPage := h.productService.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.GetPaginationProducts(products), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success get categories")
}
