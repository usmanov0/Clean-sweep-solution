package domain

type Cart struct {
	Id     int
	UserId int
	Status bool
}

type CartItems struct {
	Id        int
	CartId    int
	ProductId int
	Quantity  int
}

type CartWithItems struct {
	Cart
	Items []CartItems
}
