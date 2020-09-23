package generalMeetingSDK

import "github.com/afrouzMashaykhi/general-meeting/chaincode/stockMarket"

// Trader is a shareholder or someone with TraderID who but dividend
type Trader struct {
	// Cards is  a list of Trader's cards which have info of dividend and its pays
	Cards []Card `json:"cards"`
	// TraderID is a unique ID for every trader who register in this market
	TraderID int `json:"traderID"`
	//CardStatus []CardStatus `json:"cardstatus"`
}

// traderRegistration func is called when someone want to join the market
func traderRegister(cards []Card, traderID int) *Trader {
	return &Trader{Cards: cards, TraderID: traderID}
}
