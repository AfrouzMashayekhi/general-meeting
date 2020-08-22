package stockMarket

import "errors"

// Trader is a shareholder or someone with TraderID who but dividend
type Trader struct {
	// Cards is  a list of Trader's cards which have info of dividend and its pays
	Cards []Card `json:"cards"`
	// TraderID si a unique ID for every trader who register in this market
	TraderID int `json:"traderID"`
	//CardStatus []CardStatus `json:"cardstatus"`
}

// TraderRegistration func is called when someone want to join the market
func TraderRegistration(TraderID int) Trader {

	return Trader{nil, TraderID}
}

// InitiateTraderCards func is called when a Trader wants to claim she/he owns these dividends in joining market
func (t *Trader) InitiateTraderCards(cards []Card) error {
	for _, card := range cards {
		if card.Issuer.ValidateCard(card) == true {
			//AddCard(card)
		} else {
			return errors.New("card's owner didn't validate by issuer")
		}
	}
	return nil
}

// sellCard when someone sell card to a buyer? for now you can work on whole of a card not a portion of dividend
func (t *Trader) sellCard(card Card, buyer int) error {
	// where should I set buyer name in buy or sell and update both sides cardstatus
	if QueryByTrader(t.TraderID) != nil {
		//card.UpdateCard(buyer, nil, nil, nil, nil)
		//t.UpdateCardStatuslist(card)
	}
	return nil

}

//func (t *Trader) UpdateCardStatuslist(card Card) {
//	// change card status of trader
//}

func (t *Trader) BuyCard(card Card) {
	// updateCard
	//t.UpdateCardStatuslist(card)
}
