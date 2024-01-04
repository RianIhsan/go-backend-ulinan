package handler

import (
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"strconv"
	"ulinan/domain/category"
	"ulinan/domain/category/dto"
	"ulinan/entities"
	"ulinan/helper/cloudinary"
	"ulinan/helper/response"
	"ulinan/helper/validator"
)

type CategoryHandler struct {
	categoryService category.CategoryServiceInterface
}

func NewCategoryHandler(categoryService category.CategoryServiceInterface) category.CategoryHandlerInterface {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}

	var payload dto.TCreateCategoryRequest

	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}

	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

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
	categoryCreated, err := h.categoryService.CreateCategory(&payload)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to create category: "+err.Error())
	}

	return response.SendStatusCreatedWithDataResponse(c, "category created successfully", dto.CreateCategoryResponse(categoryCreated))

}

func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageConv, _ := strconv.Atoi(strconv.Itoa(page))
	limit := c.Query("limit", "10")
	limitInt, _ := strconv.Atoi(limit)
	search := c.Query("search")

	var categories []*entities.CategoryEntity
	var totalItems int64
	var err error

	if search != "" {
		categories, totalItems, err = h.categoryService.GetCategoryByName(pageConv, limitInt, search)
	} else if limit != "" {
		categories, totalItems, err = h.categoryService.GetAllCategories(pageConv, limitInt)
	}

	if err != nil {
		return response.SendStatusInternalServerError(c, "failed to get categories: "+err.Error())
	}

	currentPage, totalPages := h.categoryService.CalculatePaginationValues(pageConv, int(totalItems), limitInt)
	nextPage := h.categoryService.GetNextPage(currentPage, totalPages)
	prevPage := h.categoryService.GetPrevPage(currentPage)

	return response.SendPaginationResponse(c, dto.GetPaginationCategories(categories), currentPage, totalPages, int(totalItems), nextPage, prevPage, "success get categories")
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}

	id, _ := strconv.Atoi(c.Params("id"))
	var payload dto.TUpdateCategoryRequest

	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}

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

	err = h.categoryService.UpdateCategory(id, &payload)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to update category: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "category updated successfully")
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "admin" {
		return response.SendStatusUnauthorized(c, "Access denied: you are not admin")
	}

	id, _ := strconv.Atoi(c.Params("id"))

	err := h.categoryService.DeleteCategory(id)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to delete category: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "category deleted successfully")
}

func (h *CategoryHandler) GetCategoryById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	category, err := h.categoryService.GetCategoryById(id)
	if err != nil {
		return response.SendStatusBadRequest(c, "failed to get category: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "success get category", dto.GetCategoryResponse(category))
}
