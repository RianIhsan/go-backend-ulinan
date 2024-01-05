package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"ulinan/domain/cart"
	"ulinan/domain/cart/dto"
	"ulinan/entities"
	"ulinan/helper/response"
	"ulinan/helper/validator"
)

type CartHandler struct {
	service cart.CartServiceInterface
}

func NewCartHandler(service cart.CartServiceInterface) cart.CartHandlerInterface {
	return &CartHandler{service}
}

func (h *CartHandler) AddCartItem(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}

	req := new(dto.AddCartItemsRequest)
	if err := c.BodyParser(req); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	result, err := h.service.AddCartItems(currentUser.ID, req)
	if err != nil {
		return response.SendStatusInternalServerError(c, "Gagal menambahkan produk ke keranjang: "+err.Error())
	}
	return response.SendStatusCreatedWithDataResponse(c, "Berhasil menambahkan produk ke keranjang", result)
}

func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}

	cartItemSummary, err := h.service.GetCart(currentUser.ID)
	if err != nil {
		return response.SendStatusInternalServerError(c, "Gagal mendapatkan keranjang: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Berhasil mendapatkan keranjang", dto.FormatCart(cartItemSummary))
}

func (h *CartHandler) ReduceQuantity(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}
	req := new(dto.ReduceCartItemsRequest)
	if err := c.BodyParser(req); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	if err := validator.ValidateStruct(req); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	err := h.service.ReduceCartItemQuantity(req.CartItemID, req.Quantity)
	if err != nil {
		return response.SendStatusInternalServerError(c, "Gagal mengurangi jumlah produk di keranjang: "+err.Error())
	}
	return response.SendStatusOkResponse(c, "Berhasil mengurangi kuantittas")

}

func (h *CartHandler) DeleteCartItem(c *fiber.Ctx) error {
	currentUser, _ := c.Locals("CurrentUser").(*entities.UserEntity)
	if currentUser.Role != "user" {
		return response.SendStatusForbidden(c, "Access denied: you are admin, not user")
	}
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendStatusBadRequest(c, "invalid id")
	}
	err = h.service.DeleteCartItem(id)
	if err != nil {
		return response.SendStatusInternalServerError(c, "Gagal menghapus produk di keranjang: "+err.Error())
	}
	return response.SendStatusOkResponse(c, "Berhasil menghapus produk dari keranjang")

}
