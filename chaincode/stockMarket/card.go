package stockMarket

import "google.golang.org/genproto/googleapis/type/date"

type Card struct {
	Count           int             `json:"count"`
	Issuer          Issuer          `json:"issuer"`
	Dividend        int             `json:"dividend"`
	State           bool            `json:"state"`
	DividendPayment DividendPayment `json:"dividendPayment"`
}
type CardStatus struct {
	Issuer      Issuer    `json:"issuer"`
	Payment     int       `json:"payment"`
	PaymentDate date.Date `json:"paymentDate"`
}
type DividendPayment struct {
	PaymentRate float32   `json:"paymentRate"`
	PaymentDate date.Date `json:"paymentDate"`
}

//func () ChangeShareOwner() {
//
//}
//
//// // initial public offering or
////func () CreateShare ()
//func () ChangeStockOpeningPrice() {
//
//}
//
//func () ChangeStockClosingPrince() {
//
//}
//
//func () ChangeStockLowestPrice() {}
//
//func () ChangeStockHighestPRice() {}
//
//func () ChangeStockState() {}
//func () CreateStock()      {}
