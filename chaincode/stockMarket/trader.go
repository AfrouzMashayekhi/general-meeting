package stockMarket

type Trader struct {
	Cards          []Card       `json:"cards"`
	RegistrationID int          `json:"registerationID"`
	CardStatus     []CardStatus `json:"cardstatus"`
}
