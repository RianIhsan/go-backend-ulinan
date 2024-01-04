package service

import (
	"errors"
	"ulinan/domain/cart"
	"ulinan/domain/cart/dto"
	"ulinan/domain/product"
	"ulinan/entities"
)

type CartService struct {
	repo           cart.CartRepositoryInterface
	productService product.ProductServiceInterface
}

func NewCartService(repo cart.CartRepositoryInterface, productService product.ProductServiceInterface) cart.CartServiceInterface {
	return &CartService{
		repo:           repo,
		productService: productService,
	}
}

func (s *CartService) GetCart(userID int) (*entities.CartEntity, error) {
	carts, err := s.repo.GetCart(userID)
	if err != nil {
		return nil, errors.New("cart not found")
	}
	return carts, nil
}

func (s *CartService) AddCartItems(userId int, request *dto.AddCartItemsRequest) (*entities.CartItemEntity, error) {
	carts, err := s.repo.GetCartByUserId(userId)
	if err != nil {
		if carts == nil {
			newCart := &entities.CartEntity{
				UserId: userId,
			}
			_, err := s.repo.CreateCart(newCart)
			if err != nil {
				return nil, errors.New("failed create new cart")
			}
			carts = newCart
		}
	}

	existingCartItem, err := s.repo.GetCartItemByProductID(carts.Id, request.ProductID)
	if err == nil && existingCartItem != nil {
		existingCartItem.Quantity += request.Quantity
		existingCartItem.TotalPrice = existingCartItem.Quantity * existingCartItem.Price

		err := s.repo.UpdateCartItem(existingCartItem)
		if err != nil {
			return nil, errors.New("failed to change the number of products in the carte")
		}
		err = s.RecalculateGrandTotal(carts)
		if err != nil {
			return nil, errors.New("failed to recalculate the grand total")
		}
		return existingCartItem, nil
	}

	getProductByID, err := s.productService.GetProductById(request.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	cartItem := &entities.CartItemEntity{
		CartId:     carts.Id,
		ProductId:  request.ProductID,
		Quantity:   request.Quantity,
		Price:      getProductByID.Price,
		TotalPrice: getProductByID.Price * request.Quantity,
	}
	result, err := s.repo.CreateCartItem(cartItem)
	if err != nil {
		return nil, errors.New("failed to add product to cart")
	}
	err = s.RecalculateGrandTotal(carts)
	if err != nil {
		return nil, errors.New("failed to recalculate the grand total")
	}
	return result, nil

}

func (s *CartService) RecalculateGrandTotal(cart *entities.CartEntity) error {
	cartItems, err := s.repo.GetCartItemsByCartID(cart.Id)
	if err != nil {
		return err
	}
	var grandTotal int
	for _, item := range cartItems {
		grandTotal += item.TotalPrice
	}

	cart.GrandTotal = grandTotal

	err = s.repo.UpdateGrandTotal(cart.Id, grandTotal)
	if err != nil {
		return err
	}
	return nil
}
func (s *CartService) DeleteCartItem(cartItemID int) error {
	cartItem, err := s.repo.GetCartItemByID(cartItemID)
	if err != nil {
		return errors.New("item in cart not found")
	}

	carts, err := s.repo.GetCartById(cartItem.CartId)
	if err != nil {
		return errors.New("cart not found")
	}

	err = s.repo.DeleteCartItem(cartItem.Id)
	if err != nil {
		return errors.New("failed to delete item in cart")
	}

	err = s.RecalculateGrandTotal(carts)
	if err != nil {
		return errors.New("failed to recalculate the grand total")
	}

	return nil
}

func (s *CartService) ReduceCartItemQuantity(cartItemID, quantity int) error {
	cartItem, err := s.repo.GetCartItemByID(cartItemID)
	if err != nil {
		return errors.New("item in cart not found")
	}
	if quantity > cartItem.Quantity {
		return errors.New("the quantity demanded exceeds the quantity in the cart")
	}

	cartItem.Quantity -= quantity
	cartItem.TotalPrice = cartItem.Quantity * cartItem.Price

	if cartItem.Quantity == 0 {
		err := s.repo.DeleteCartItem(cartItemID)
		if err != nil {
			return errors.New("failed to delete item in cart")
		}
	} else {
		err = s.repo.UpdateCartItem(cartItem)
		if err != nil {
			return errors.New("failed to update item in cart")
		}
	}
	carts, err := s.repo.GetCartById(cartItem.CartId)
	if err != nil {
		return errors.New("cart not found")
	}

	err = s.RecalculateGrandTotal(carts)
	if err != nil {
		return errors.New("failed to recalculate the grand total")
	}

	return nil
}

func (s *CartService) IsProductInCart(userID, productID int) bool {
	isInCart := s.repo.IsProductInCart(userID, productID)
	return isInCart
}

func (s *CartService) RemoveProductFromCart(userID, productID int) error {
	isProductInCart := s.repo.IsProductInCart(userID, productID)
	if !isProductInCart {
		return errors.New("the product is not in the user's cart")
	}

	err := s.repo.RemoveProductFromCart(userID, productID)
	if err != nil {
		return errors.New("failed remove product on cart")
	}

	return nil
}

func (s *CartService) GetCartItems(cartItem int) (*entities.CartItemEntity, error) {
	cartItems, err := s.repo.GetCartItemByID(cartItem)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
