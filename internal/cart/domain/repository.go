package domain

type CartRepository interface {
	SaveCart(userId int) (int, error)
	GetCartWithItems(cartId int) (*CartWithItems, error)
	GetActiveCart(userID int) (*Cart, error)
	MarkCartStatusAsTrue(userId, cartId int) error
}

type CartItemRepository interface {
	AddItem(items *CartItems) (int, error)
	GetAll(cartId int) ([]CartItems, error)
	UpdateCartItem(cItemId, quantity int) error
	DeleteProduct(cItemId int) error
}
