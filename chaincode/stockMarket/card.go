package stockMarket

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"time"
)

// each object should contains its own contract means its own invoke and init ledger
// StockContract contract for handling writing and reading from the world state
type StockContract struct {
	contractapi.Contract
}

// Card is a struct for dividend info and owners
type Card struct {
	// TraderID added for more readable and query
	// make it string for compositekey
	TraderID string `json:"traderID"`
	// Count how many share owns this traderID
	Count       int    `json:"count"`
	StockSymbol string `json:"stockSymbol"`
	// Dividend toman/share
	Dividend int `json:"dividend"`
	// mno need for state
	//State           bool              `json:"state"`
	// DividendPayments the plan of paying dividend
	DividendPayments []DividendPayment `json:"dividendPayment"`
}

// DividendPayment status of time plan of dividend pays
type DividendPayment struct {
	// Percentage of dividend count pays
	Percentage float32 `json:"percentage"`
	// PDate date of payment
	PDate time.Time `json:"pDate"`
	// for changing status of
	Paid bool `json:"paid"`
}

func (sc *StockContract) InitLedger(ctx contractapi.TransactionContextInterface) {
	// todo: get all trader id and issuer id and make card calls AddCard function
}

// AddCard calls putState of chaincode to add card maybe create a string to push in worldstate
func (sc *StockContract) AddCard(ctx contractapi.TransactionContextInterface, card Card) error {
	indexName := "trader~stocksymbol"
	cardAsByte, _ := json.Marshal(card)
	// todo: error handling
	// todo: validate by issuer is handeled here?
	cardKey, _ := ctx.GetStub().CreateCompositeKey(indexName, []string{card.TraderID, card.StockSymbol})
	err := ctx.GetStub().PutState(cardKey, cardAsByte)
	if err != nil {
		return fmt.Errorf("failed to put Card to world state %s", err.Error())
	}
	return nil
}

// not needed if we zero things
//func (sc *StockContract) DeleteCard(ctx contractapi.TransactionContextInterface, card Card) error {
//	indexName := "trader~stocksymbol"
//	cardKey, _ := ctx.GetStub().CreateCompositeKey(indexName, []string{card.TraderID, card.StockSymbol})
//	err := ctx.GetStub().DelState(cardKey)
//	if err != nil {
//		return fmt.Errorf("failed to delete Card from world state %s", err.Error())
//	}
//	return nil
//}

func (sc *StockContract) QueryByTrader(ctx contractapi.TransactionContextInterface, traderID string) []Card {

}

func (sc *StockContract) QueryByStockSymbol(ctx contractapi.TransactionContextInterface, stockSymbol string) []Card {

}

func (sc *StockContract) Trade(ctx contractapi.TransactionContextInterface, seller string, buyer string, count int, stockSymbol string) error {
	// todo: get seller+ stock
	//todo: if updateCard count  and dividend payment
	//todo:add count and dividentPayment to buyer
}

func (sc *StockContract) UpdateCount(ctx contractapi.TransactionContextInterface, card Card, countChange int) error {
	// todo:get card
	// todo:change count
	// todo:putcard

}
func (sc *StockContract) UpdateDividendPayment(ctx contractapi.TransactionContextInterface, card Card, dPayment DividendPayment) error {
	// todo:get card
	// change edpayment attributes
}
func (sc *StockContract) AddDividendPayment(ctx contractapi.TransactionContextInterface, card Card, dPayment DividendPayment) error {
	// todo:get card
	// add dpayment
}
