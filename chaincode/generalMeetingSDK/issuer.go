package generalMeetingSDK

import (
	"encoding/json"
	"fmt"
	sm "github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"strconv"
)

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
var reservedStockSymbol []byte = []byte("afrouz")

// Issuer is a company that validate and pays Dividend and holds General Meetings
type Issuer struct {
	CompanyName string `json:"companyName"`
	StockSymbol string `json:"stockSymbol"`
}

// RegisterIssuer is called for new company to join stock market
func RegisterIssuer(ccName string, client *channel.Client, companyName string, stockSymbol string) *Issuer {
	// add cards to worldstate for every issuer in market
	response, err := client.Query(channel.Request{
		ChaincodeID: ccName,
		Fcn:         "QueryByStockSymbol",
		Args:        [][]byte{reservedStockSymbol},
		IsInit:      false,
	})
	if err != nil {
		fmt.Errorf("couldn't query cards for%s", stockSymbol)
	}
	cards := sm.QueryCard{}
	_ = json.Unmarshal(response.Payload, &cards)
	for _, card := range cards.Cards {
		invokeArgs := [][]byte{[]byte(card.TraderID), []byte(stockSymbol), []byte("0"), []byte("0")}
		_, err := client.Execute(channel.Request{
			ChaincodeID: ccName,
			Fcn:         "AddCard",
			Args:        invokeArgs,
		}, channel.WithRetry(retry.DefaultChannelOpts))

		if err != nil {
			fmt.Printf("Failed to invoke AddCard for issuer: %+v\n", err)
		}

	}

	issuer := Issuer{
		CompanyName: companyName,
		StockSymbol: stockSymbol,
	}
	return &issuer
}

// ValidateCard Func Validates traders cards if they own this company share or not
func (i *Issuer) ValidateCard(card sm.Card) bool {
	//todo: do we have a list or anything should ask from Analoui
	return true
}

// GeneralMeeting Func took place and add card to shareholders for dividend and dividendPayment
func (i *Issuer) GeneralMeeting(ccName string, client *channel.Client, dividend int, payments []sm.DividendPayment) error {
	response, err := client.Query(channel.Request{
		ChaincodeID: ccName,
		Fcn:         "QueryByStockSymbol",
		Args:        [][]byte{[]byte(i.StockSymbol)},
		IsInit:      false,
	})
	if err != nil {
		return fmt.Errorf("couldn't query cards for general meeting %s\n", i.StockSymbol)
	}
	cards := sm.QueryCard{}
	_ = json.Unmarshal(response.Payload, &cards)
	for _, card := range cards.Cards {
		//if count be zero it means the trader doesn't have any dividend
		if card.Count != 0 {
			dividendString := strconv.Itoa(dividend)
			paymentAsByte, _ := json.Marshal(payments)
			invokeArgs := [][]byte{[]byte(card.TraderID), []byte(i.StockSymbol), []byte(dividendString), paymentAsByte}
			_, err := client.Execute(channel.Request{
				ChaincodeID: ccName,
				Fcn:         "UpdateDividend",
				Args:        invokeArgs,
			}, channel.WithRetry(retry.DefaultChannelOpts))

			if err != nil {
				return fmt.Errorf("Failed to updateDividend: %+v for card: %+v\n", err, card)
			}
		}
	}

	return nil
}

// GetCards return all cards of issuer
func (i *Issuer) GetCards(ccName string, client *channel.Client) (sm.QueryCard, error) {
	response, err := client.Query(channel.Request{
		ChaincodeID: ccName,
		Fcn:         "QueryByStockSymbol",
		Args:        [][]byte{[]byte(i.StockSymbol)},
		IsInit:      false,
	})
	if err != nil {
		return sm.QueryCard{}, fmt.Errorf("couldn't query cards for issuer %s\n", i.StockSymbol)
	}
	cards := sm.QueryCard{}
	_ = json.Unmarshal(response.Payload, &cards)
	return cards, nil
}

//// PayCard Func at the time of payDate
//func (i *Issuer) PayCard(payDate time.Time) {
//	cards := QueryByIssuer(*i)
//	for _, card := range cards {
//		for _, dDate := range card.DividendPayments {
//			if dDate.PDate.Equal(payDate) {
//				dDate.Paid = true
//
//			}
//		}
//		//card.UpdateCard()
//	}
//
//}
