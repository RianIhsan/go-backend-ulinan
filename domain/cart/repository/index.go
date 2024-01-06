package repository

import (
	"gorm.io/gorm"
	"ulinan/domain/cart"
	"ulinan/entities"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) cart.CartRepositoryInterface {
	return &CartRepository{db}
}

func (r *CartRepository) GetCartByUserId(userId int) (*entities.CartEntity, error) {
	carts := &entities.CartEntity{}
	if err := r.db.Where("user_id = ?", userId).First(carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *CartRepository) GetCartById(cartId int) (*entities.CartEntity, error) {
	carts := &entities.CartEntity{}
	if err := r.db.First(carts, cartId).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *CartRepository) GetCartItemByProductID(cartId, productId int) (*entities.CartItemEntity, error) {
	carts := entities.CartItemEntity{}
	if err := r.db.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&carts).Error; err != nil {
		return nil, err
	}
	return &carts, nil
}

func (r *CartRepository) GetCartItemsByCartID(cartId int) ([]*entities.CartItemEntity, error) {
	var cartItems []*entities.CartItemEntity
	if err := r.db.Where("cart_id = ?", cartId).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (r *CartRepository) GetCartItemByID(cartItemId int) (*entities.CartItemEntity, error) {
	var cartItem *entities.CartItemEntity
	if err := r.db.Where("id = ?", cartItemId).First(&cartItem).Error; err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (r *CartRepository) GetCart(userID int) (*entities.CartEntity, error) {
	carts := &entities.CartEntity{}
	if err := r.db.
		Preload("CartItems", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, cart_id, product_id, quantity, total_price, arrival_date").
				Preload("Product", func(db *gorm.DB) *gorm.DB {
					return db.Select("id, name, price").
						Preload("ProductPhotos")
				})
		}).Where("user_id = ?", userID).First(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *CartRepository) CreateCart(newCart *entities.CartEntity) (*entities.CartEntity, error) {
	err := r.db.Create(newCart).Error
	if err != nil {
		return nil, err
	}
	return newCart, nil
}

func (r *CartRepository) CreateCartItem(cartItem *entities.CartItemEntity) (*entities.CartItemEntity, error) {
	err := r.db.Create(cartItem).Error
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (r *CartRepository) UpdateCartItem(cartItem *entities.CartItemEntity) error {
	if err := r.db.Save(&cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (r *CartRepository) UpdateGrandTotal(cartID, grandTotal int) error {
	var carts *entities.CartEntity
	result := r.db.Model(&carts).Where("id = ?", cartID).Update("grand_total", grandTotal)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CartRepository) DeleteCartItem(cartItemID int) error {
	result := r.db.Where("id = ?", cartItemID).Delete(&entities.CartItemEntity{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *CartRepository) IsProductInCart(userID, productID int) bool {
	var count int64
	r.db.Model(&entities.CartItemEntity{}).
		Joins("JOIN carts ON cart_items.cart_id = carts.id").
		Where("carts.user_id = ? AND cart_items.product_id = ?", userID, productID).
		Count(&count)
	return count > 0
}

func (r *CartRepository) RemoveProductFromCart(userID, productID int) error {
	var carts entities.CartEntity
	if err := r.db.Where("user_id = ?", userID).Preload("CartItems").First(&carts).Error; err != nil {
		return err
	}

	var cartItem entities.CartItemEntity
	for _, item := range carts.CartItems {
		if item.ProductId == productID {
			cartItem = *item
			break
		}
	}
	if cartItem.Id == 0 {
		return nil
	}
	if err := r.db.Delete(&cartItem).Error; err != nil {
		return err
	}

	return nil
}
