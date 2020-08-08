package stockMarket

import "google.golang.org/genproto/googleapis/type/date"

type Card struct {
	TraderID int    `json:"traderID"`
	Count    int    `json:"count"`
	Issuer   Issuer `json:"issuer"`
	Dividend int    `json:"dividend"`
	//State           bool              `json:"state"`
	DividendPayment []DividendPayment `json:"dividendPayment"`
}

//type CardStatus struct {
//	Issuer      Issuer    `json:"issuer"`
//	Payment     int       `json:"payment"`
//	PaymentDate date.Date `json:"paymentDate"`
//	Paid        bool      `json:"paid"`
//}
type DividendPayment struct {
	Percentage float32   `json:"percentage"`
	PDate      date.Date `json:"pDate"`
	Paid       bool      `json:"paid"`
}

func (c *Card) addCard(card Card) error {
	//call something like createCar
	return nil
}

//UpdateCard should take allvariables or get card and update it maybe be a func not a method is better
func (c *Card) UpdateCard(traderID int, count int, issuer Issuer, dividend int, dPayment DividendPayment) {
	//queryCard and update it  on worlstate
}
func QueryCardByTrader(traderId int) []Card {
	//query on worldstate
	return nil
}
func QueryCardByIssuer(issuer Issuer) []Card {
	//query on worldstate
	return nil
}
