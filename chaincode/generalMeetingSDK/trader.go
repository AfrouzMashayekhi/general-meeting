package main

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
import (
	"encoding/json"
	"fmt"
	sm "github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"strconv"
)

// Trader is a shareholder or someone with TraderID who but dividend
type Trader struct {
	// Cards is  a list of Trader's cards which have info of dividend and its pays
	Cards []sm.Card `json:"cards"`
	// TraderID is a unique ID for every trader who register in this market
	TraderID string `json:"traderID"`
}

var reservedTraderID []byte = []byte("afrouz")

// RegisterTrader func is called when someone want to join the market
func RegisterTrader(ccName string, client *channel.Client, traderID string) *Trader {

	// add cards to worldstate for every trader in market
	response, err := client.Query(channel.Request{
		ChaincodeID: ccName,
		Fcn:         "QueryByTrader",
		Args:        [][]byte{reservedTraderID},
		IsInit:      false,
	})
	if err != nil {
		fmt.Errorf("couldn't query cards for%s %s", reservedTraderID, err)
	}
	cards := sm.QueryCard{}
	_ = json.Unmarshal(response.Payload, &cards)
	for _, card := range cards.Cards {
		invokeArgs := [][]byte{[]byte(traderID), []byte("0"), []byte(card.StockSymbol), []byte("0")}
		_, err := client.Execute(channel.Request{
			ChaincodeID: ccName,
			Fcn:         "AddCard",
			Args:        invokeArgs,
		}, channel.WithRetry(retry.DefaultChannelOpts))

		if err != nil {
			fmt.Errorf("Failed to invoke AddCard : %+v\n", err)
		}

	}

	trader := Trader{TraderID: traderID}
	return &trader
}

// AddCards func add cards for trader of issuer validate it return true
func (t *Trader) AddCards(ccName string, client *channel.Client, cards []sm.Card) error {

	for _, card := range cards {
		// todo: it's a distributed app how can I find issuer create new one?
		issuer := Issuer{
			StockSymbol: card.StockSymbol,
		}
		if issuer.ValidateCard(card) {
			//fmt.Println(card.DividendPayments)
			paymentAsByte, _ := json.Marshal(card.DividendPayments)
			//fmt.Println(paymentAsByte)
			countString := strconv.Itoa(card.Count)
			dividendString := strconv.Itoa(card.Dividend)
			invokeArgs := [][]byte{[]byte(card.TraderID), []byte(card.StockSymbol), []byte(countString), []byte(dividendString), paymentAsByte}
			_, err := client.Execute(channel.Request{
				ChaincodeID: ccName,
				Fcn:         "UpdateFields",
				Args:        invokeArgs,
			}, channel.WithRetry(retry.DefaultChannelOpts))

			if err != nil {
				return fmt.Errorf("Failed to validate and update card: %+v\n", err)
			}
		} else {
			return fmt.Errorf("the Card : %+v is not Validated by issuer", card)
		}
	}
	return nil
}

// Trading trade from seller to buyer the buyCount mount if succeeded return true
// should it be function or method for seller? no it should be func it calls from outside
func Trading(ccName string, client *channel.Client, seller string, buyer string, buyCount int, stockSymbol string) error {
	countString := strconv.Itoa(buyCount)
	invokeArgs := [][]byte{[]byte(seller), []byte(buyer), []byte(countString), []byte(stockSymbol)}
	_, err := client.Execute(channel.Request{
		ChaincodeID: ccName,
		Fcn:         "Trade",
		Args:        invokeArgs,
	}, channel.WithRetry(retry.DefaultChannelOpts))

	if err != nil {
		return fmt.Errorf("Failed to trade card: %+v\n", err)
	}
	return nil
}

// GetCards return all cards of trader
func (t *Trader) GetCards(ccName string, client *channel.Client) (sm.QueryCard, error) {
	response, err := client.Query(channel.Request{
		ChaincodeID: ccName,
		Fcn:         "QueryByTrader",
		Args:        [][]byte{[]byte(t.TraderID)},
		IsInit:      false,
	})
	if err != nil {
		return sm.QueryCard{}, fmt.Errorf("couldn't query cards for trader  %s %s\n", t.TraderID, err)
	}
	cards := sm.QueryCard{}
	_ = json.Unmarshal(response.Payload, &cards)
	return cards, nil
}
