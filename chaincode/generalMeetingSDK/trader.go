package generalMeetingSDK

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
	"fmt"
	sm "github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// Trader is a shareholder or someone with TraderID who but dividend
type Trader struct {
	// Cards is  a list of Trader's cards which have info of dividend and its pays
	Cards []sm.Card `json:"cards"`
	// TraderID is a unique ID for every trader who register in this market
	TraderID string `json:"traderID"`
}

var reservedStockSymbol []byte = []byte("afrouz")
var reservedTraderID []byte = []byte("afrouz")

// RegisterTrader func is called when someone want to join the market
func RegisterTrader(ccName string, sdk *fabsdk.FabricSDK, client *channel.Client) *Trader {
	//todo: add trader id
	// todo: add cards to worldstate for every trader in market

	response, err := client.Query(channel.Request{
		ChaincodeID: ccName,
		Fcn:         "QueryByTrader",
		Args:        [][]byte{reservedTraderID},
		IsInit:      false,
	})
	if err != nil {
		fmt.Errorf("couldn't query cards for")
	}
	cards := sm.QueryCard{}

	//trader := Trader{TraderID: traderID}
	//setup.trader= &trader
	//return &trader
	return nil
}

// AddCards func add cards for trader of issuer validate it return true
func (t *Trader) AddCards(client *channel.Client, cards []sm.Card) bool {
	//todo: call validateCard
	//todo: if validated call transaction add card
	return true
}

// Trading trade from seller to buyer the buyCount mount if succeeded return true
// should it be function or method for seller? no it should be func it calls from outside
func Trading(client *channel.Client, seller string, buyer string, buyCount int, stockSymbol string) bool {
	//todo:execute Trade
	return true
}
