package user

type Kind string

const (
	Customer Kind = "CUSTOMER"
	Seller   Kind = "SELLER"
)

func (k Kind) IsSeller() bool {
	return k == Seller
}

func (k Kind) String() string {
	return string(k)
}
