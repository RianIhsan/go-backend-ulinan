package cart

import (
	"github.com/gofiber/fiber/v2"
	"ulinan/domain/cart/dto"
	"ulinan/entities"
)

type CartRepositoryInterface interface {
	CreateCart(newCart *entities.CartEntity) (*entities.CartEntity, error)
	CreateCartItem(cartItem *entities.CartItemEntity) (*entities.CartItemEntity, error)
	GetCartByUserId(userId int) (*entities.CartEntity, error)
	GetCartById(cartId int) (*entities.CartEntity, error)
	GetCartItemByProductID(cartId, productId int) (*entities.CartItemEntity, error)
	GetCartItemsByCartID(cartId int) ([]*entities.CartItemEntity, error)
	GetCartItemByID(cartItemId int) (*entities.CartItemEntity, error)
	GetCart(userID int) (*entities.CartEntity, error)
	UpdateCartItem(cartItem *entities.CartItemEntity) error
	UpdateGrandTotal(cartID, grandTotal int) error
	DeleteCartItem(cartItemID int) error
	IsProductInCart(userID, productID int) bool
	RemoveProductFromCart(userID, productID int) error
}

type CartServiceInterface interface {
	GetCart(userID int) (*entities.CartEntity, error)
	AddCartItems(userId int, request *dto.AddCartItemsRequest) (*entities.CartItemEntity, error)
	RecalculateGrandTotal(cart *entities.CartEntity) error
	DeleteCartItem(cartItemID int) error
	ReduceCartItemQuantity(cartItemID, quantity int) error
	IsProductInCart(userID, productID int) bool
	RemoveProductFromCart(userID, productID int) error
	GetCartItems(cartItem int) (*entities.CartItemEntity, error)
}

type CartHandlerInterface interface {
	AddCartItem(c *fiber.Ctx) error
	GetCart(c *fiber.Ctx) error
	ReduceQuantity(c *fiber.Ctx) error
	DeleteCartItem(c *fiber.Ctx) error
}
