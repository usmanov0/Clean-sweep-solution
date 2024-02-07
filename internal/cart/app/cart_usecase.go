package app

import (
	"example.com/m/internal/cart/domain"
)

type CartService interface {
	CreateCart(userID int) (int, error)
	AddItems(userId int, items *domain.CartItems) (int, error)
	GetCartWithItems(cartId int) (*domain.CartWithItems, error)
	GetActiveCarts(userId int) (*domain.Cart, error)
	GetAllItems(cartId int) ([]domain.CartItems, error)
	UpdateCart(cItemId, quantity int) error
	MarkCartAsPurchased(userId, cartId int) error
	Delete(cItemId int) error
}

type cartService struct {
	cartRepo     domain.CartRepository
	cartItemRepo domain.CartItemRepository
}

func NewCartService(cartRepo domain.CartRepository, cartItemRepo domain.CartItemRepository) CartService {
	return &cartService{
		cartRepo:     cartRepo,
		cartItemRepo: cartItemRepo}
}

func (c *cartService) CreateCart(userID int) (int, error) {
	cartId, err := c.cartRepo.SaveCart(userID)

	if err != nil {
		return 0, err
	}

	return cartId, nil
}

func (c *cartService) AddItems(userId int, items *domain.CartItems) (int, error) {
	activeCart, err := c.cartRepo.GetActiveCart(userId)
	if err != nil {
		return 0, err
	}
	if !activeCart.Status {
		id, err := c.cartItemRepo.AddItem(&domain.CartItems{
			CartId:    activeCart.Id,
			ProductId: items.ProductId,
			Quantity:  items.Quantity,
		})
		if err != nil {
			return 0, err
		}
		return id, nil
	}

	newCartId, err := c.cartRepo.SaveCart(userId)
	if err != nil {
		return 0, err
	}
	id, err := c.cartItemRepo.AddItem(&domain.CartItems{
		CartId:    newCartId,
		ProductId: items.ProductId,
		Quantity:  items.Quantity,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *cartService) GetCartWithItems(cartId int) (*domain.CartWithItems, error) {
	cWithItems, err := c.cartRepo.GetCartWithItems(cartId)
	if err != nil {
		return nil, err
	}
	return cWithItems, nil
}

func (c *cartService) GetActiveCarts(userId int) (*domain.Cart, error) {
	cart, err := c.cartRepo.GetActiveCart(userId)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (c *cartService) GetAllItems(cartId int) ([]domain.CartItems, error) {
	items, err := c.cartItemRepo.GetAll(cartId)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *cartService) UpdateCart(cItemId, quantity int) error {
	err := c.cartItemRepo.UpdateCartItem(cItemId, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (c *cartService) MarkCartAsPurchased(userId, cartId int) error {
	err := c.cartRepo.MarkCartStatusAsTrue(userId, cartId)
	if err != nil {
		return err
	}
	return nil
}

func (c *cartService) Delete(cItemId int) error {
	err := c.cartItemRepo.DeleteProduct(cItemId)
	if err != nil {
		return err
	}
	return nil
}
