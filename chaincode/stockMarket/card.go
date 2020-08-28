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
	//// for changing status of NOT NEEDED I THINK IF  PAYING IS OUT OF BLOCKCHAIN JUST QUERY ISSUER
	//Paid bool `json:"paid"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Card
}

// InitLedger create all cards with TraderID and Issuer and other attr nil
func (sc *StockContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// todo: get all trader id and issuer id and make card calls AddCard function
	// todo: for now just add some static issuer trader
	cards := []Card{}
	for _, card := range cards {
		err := sc.AddCard(ctx, card)
		if err != nil {
			return fmt.Errorf("failed to init Cards %s", err.Error())
		}
	}
	return nil

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

func (sc *StockContract) QueryByTrader(ctx contractapi.TransactionContextInterface, traderID string) ([]QueryResult, error) {
	traderIterator, err := ctx.GetStub().GetStateByPartialCompositeKey("trader~stocksymbol", []string{traderID})
	if err != nil {
		return nil, fmt.Errorf("failed to get Cards of trader %s", err.Error())
	}
	defer traderIterator.Close()
	cards := []QueryResult{}
	for traderIterator.HasNext() {
		response, err := traderIterator.Next()
		if err != nil {
			return nil, err
		}
		card := new(Card)
		_ = json.Unmarshal(response.Value, card)
		result := QueryResult{Key: response.Key, Record: card}
		cards = append(cards, result)

	}
	return cards, nil

}

func (sc *StockContract) QueryByStockSymbol(ctx contractapi.TransactionContextInterface, stockSymbol string) ([]QueryResult, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"stockSymbol\":\"%s\"}}", stockSymbol)
	ssymbolIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, fmt.Errorf("failed to get Cards of stock symbol %s", err.Error())
	}
	defer ssymbolIterator.Close()
	cards := []QueryResult{}
	for ssymbolIterator.HasNext() {
		response, err := ssymbolIterator.Next()
		if err != nil {
			return nil, err
		}
		card := new(Card)
		_ = json.Unmarshal(response.Value, card)
		result := QueryResult{Key: response.Key, Record: card}
		cards = append(cards, result)

	}
	return cards, nil

}

func (sc *StockContract) Trade(ctx contractapi.TransactionContextInterface, seller string, buyer string, count int, stockSymbol string) error {
	// todo: get seller+ stock
	indexName := "trader~stocksymbol"
	sellerCardKey, _ := ctx.GetStub().CreateCompositeKey(indexName, []string{seller, stockSymbol})
	sellerResponse, err := ctx.GetStub().GetState(sellerCardKey)
	if err != nil {
		return fmt.Errorf("failed to get Card from world state %s", err.Error())
	}
	if sellerResponse == nil {
		return fmt.Errorf("not such a card in worldstate")
	}
	sellerResponseCard := new(Card)
	_ = json.Unmarshal(sellerResponse, sellerResponseCard)
	// add it here cause its refrence and maybe deleted in calling update coutn
	dividendPaymentCard := sellerResponseCard.DividendPayments
	// negative the number
	err = sc.updateCount(ctx, sellerResponseCard, -count)
	if err != nil {
		return fmt.Errorf("can't sell card update count ")
	}
	//todo: if updateCard count  and dividend payment do I need to update dividend payment if zero so not added?
	// but maybe again she buy it again so clear it up? if zero?
	//todo:add count and dividentPayment to buyer
	buyerCardKey, _ := ctx.GetStub().CreateCompositeKey(indexName, []string{buyer, stockSymbol})
	buyerResponse, err := ctx.GetStub().GetState(buyerCardKey)
	if err != nil {
		return fmt.Errorf("failed to get Card from world state %s", err.Error())
	}
	buyerResponseCard := new(Card)
	_ = json.Unmarshal(buyerResponse, buyerResponseCard)
	// negative the number
	err = sc.updateCount(ctx, buyerResponseCard, count)
	if err != nil {
		return fmt.Errorf("can't buy card update count")
	}
	err = sc.addDividendPayment(ctx, buyerResponseCard, dividendPaymentCard)
	if err != nil {
		return fmt.Errorf("can't buy card update dpayment")
	}
	return nil
}

func (sc *StockContract) updateCount(ctx contractapi.TransactionContextInterface, card *Card, countChange int) error {
	indexName := "trader~stocksymbol"
	cardKey, _ := ctx.GetStub().CreateCompositeKey(indexName, []string{card.TraderID, card.StockSymbol})
	response, err := ctx.GetStub().GetState(cardKey)
	if err != nil {
		return fmt.Errorf("failed to get Card from world state %s", err.Error())
	}
	if response == nil {
		return fmt.Errorf("not such a card in worldstate")
	}
	responseCard := new(Card)
	_ = json.Unmarshal(response, responseCard)
	responseCard.Count += countChange
	if responseCard.Count <= 0 {
		return fmt.Errorf("can't update count the count will be negative ")
	}
	if responseCard.Count == 0 {
		sc.deleteDividendPayment(ctx, responseCard)
	} else {
		cardAsByte, _ := json.Marshal(responseCard)
		err = ctx.GetStub().PutState(cardKey, cardAsByte)
		if err != nil {
			return fmt.Errorf("failed to put Card to world state %s", err.Error())
		}
	}
	return nil
}

func (sc *StockContract) addDividendPayment(ctx contractapi.TransactionContextInterface, card *Card, dPayments []DividendPayment) error {
	indexName := "trader~stocksymbol"
	cardKey, _ := ctx.GetStub().CreateCompositeKey(indexName, []string{card.TraderID, card.StockSymbol})
	response, err := ctx.GetStub().GetState(cardKey)
	if err != nil {
		return fmt.Errorf("failed to get Card from world state %s", err.Error())
	}
	if response == nil {
		return fmt.Errorf("not such a card in worldstate")
	}
	responseCard := new(Card)
	_ = json.Unmarshal(response, responseCard)
	for _, dPayment := range dPayments {
		responseCard.DividendPayments = append(responseCard.DividendPayments, dPayment)
	}
	cardAsByte, _ := json.Marshal(responseCard)
	err = ctx.GetStub().PutState(cardKey, cardAsByte)
	if err != nil {
		return fmt.Errorf("failed to put Card to world state %s", err.Error())
	}
	return nil
}

func (sc *StockContract) deleteDividendPayment(ctx contractapi.TransactionContextInterface, card *Card) error {
	indexName := "trader~stocksymbol"
	cardKey, _ := ctx.GetStub().CreateCompositeKey(indexName, []string{card.TraderID, card.StockSymbol})
	response, err := ctx.GetStub().GetState(cardKey)
	if err != nil {
		return fmt.Errorf("failed to get Card from world state %s", err.Error())
	}
	if response == nil {
		return fmt.Errorf("not such a card in worldstate")
	}
	responseCard := new(Card)
	_ = json.Unmarshal(response, responseCard)
	responseCard.DividendPayments = nil
	cardAsByte, _ := json.Marshal(responseCard)
	err = ctx.GetStub().PutState(cardKey, cardAsByte)
	if err != nil {
		return fmt.Errorf("failed to put Card to world state %s", err.Error())
	}
	return nil
}
