package generalmeeting

import (
	"google.golang.org/genproto/googleapis/type/date"
)

type Share struct {
	owner   Trader `json:"owner"`
	company  CorporateTrader `json:"company"`
	issueDate  date.Date `json:"issueDate"`
	shareID string `json:"shareID"`
	// maybe better name or really do I need diffrent struct for stock my reason is for global state and array of share for traders
	stockData Stock `json:"stockData"`
}
 type Stock struct {
	company  CorporateTrader `json:"company"`
	ipoDate  date.Date `json:"ipoDate"`
	openingPrice int `"json:openingPrice"`
	closingPrice int `"json:closingPrince"`
	stockSymbol string `"json:stockSymbol"`
	lowestPrice int `"json:lowestPrice"`
	highestPrice int `"json:highestPrince"`
	// state i dont know the type of state
 }

func () ChangeShareOwner () {

}
// // initial public offering or
//func () CreateShare ()
func () ChangeStockOpeningPrice () {

}

func () ChangeStockClosingPrince () {

}

func () ChangeStockLowestPrice () {}

func () ChangeStockHighestPRice () {}

func () ChangeStockState () {}
func () CreateStock () {}
