package stockMarket

import (
	"google.golang.org/genproto/googleapis/type/date"
)

// each object should contains its own contract means its own involke and initledger

type Share struct {
	owner Trader `json:"owner"`
	//issueDate should be time.Time not supported in
	//https://github.com/hyperledger/fabric-contract-api-go/blob/master/tutorials/getting-started.md
	issueDate    date.Date `json:"issueDate"`
	serialNumber string    `json:"serialNumber"`
	// maybe better name or really do I need diffrent struct for stock my reason is for global state and array of share for traders
	stockData Stock `json:"stockData"`
}

// stock it's unnecessary
type Stock struct {
	stockSymbol string
	//issueDate should be time.Time not supported in
	//https://github.com/hyperledger/fabric-contract-api-go/blob/master/tutorials/getting-started.md
	ipoDate      date.Date `json:"ipoDate"`
	openingPrice int       `"json:openingPrice"`
	closingPrice int       `"json:closingPrince"`
	//stockSymbol  string          `"json:stockSymbol"`
	lowestPrice  int `"json:lowestPrice"`
	highestPrice int `"json:highestPrince"`
	// state i dont know the type of state
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
