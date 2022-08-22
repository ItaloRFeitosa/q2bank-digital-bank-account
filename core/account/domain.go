package account

const (
	BRL string = "BRL"
)

func OpenAccount(userID uint) NewAccountInput {
	return NewAccountInput{
		Balance:  10000, // for test purpose
		Currency: BRL,
		UserID:   userID,
	}
}
