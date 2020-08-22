package stockMarket

import (
	"time"
)

// Card is a struct for dividend info and owners
type Card struct {
	// TraderID added for more readable and query
	TraderID int `json:"traderID"`
	// Count how many share owns this traderID
	Count  int    `json:"count"`
	Issuer Issuer `json:"issuer"`
	// Dividend toman/share
	Dividend int `json:"dividend"`
	// mno need for state
	//State           bool              `json:"state"`
	// DividendPayments the plan of paying dividend
	DividendPayments []DividendPayment `json:"dividendPayment"`
}

//type CardStatus struct {
//	Issuer      Issuer    `json:"issuer"`
//	Payment     int       `json:"payment"`
//	PaymentDate date.Date `json:"paymentDate"`
//	Paid        bool      `json:"paid"`
//}
// DividendPayment status of time plan of dividend pays
type DividendPayment struct {
	// Percentage of dividend count pays
	Percentage float32 `json:"percentage"`
	// PDate date of payment
	PDate time.Time `json:"pDate"`
	// for changing status of
	Paid bool `json:"paid"`
}

// addCard calls putState of chaincode to add card maybe create a string to push in worldstate
func (c *Card) addCard(card Card) error {
	//call something like createCar
	return nil
}

//UpdateCard should take allvariables or get card and update it maybe be a func not a method is better
func (c *Card) UpdateCard(traderID int, count int, issuer Issuer, dividend int, dPayment DividendPayment) {
	//queryCard and update it  on worlstate
}
func QueryByTrader(traderID int) []Card {
	//query on worldstate
	return nil
}
func QueryByIssuer(issuer Issuer) []Card {
	//query on worldstate
	return nil
}
