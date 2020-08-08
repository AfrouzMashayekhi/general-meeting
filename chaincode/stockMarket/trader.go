package stockMarket

import "errors"

type Trader struct {
	Cards    []Card `json:"cards"`
	TraderID int    `json:"traderID"`
	//CardStatus []CardStatus `json:"cardstatus"`
}

func TraderRegistration(TraderID int) Trader {

	return Trader{nil, TraderID}
}

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

func (t *Trader) sellCard(card Card, buyer int) error {
	// where should I set buyer name in buy or sell and update both sides cardstatus
	if QueryCardByTrader(t.TraderID) != nil {
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
