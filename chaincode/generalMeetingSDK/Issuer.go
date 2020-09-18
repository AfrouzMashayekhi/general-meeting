package generalMeetingSDK

import sm "github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket"

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

// Issuer is a company that validate and pays Dividend and holds General Meetings
type Issuer struct {
	CompanyName string `json:"companyName"`
	StockSymbol string `json:"stockSymbol"`
}

// RegisterIssuer is called for new company to join stock market
func RegisterIssuer(companyName string, stockSymbol string) *Issuer {
	// todo:should we enroll new user for company
	// todo: add cards to worldstate for every trader in market
	return &Issuer{CompanyName: companyName, StockSymbol: stockSymbol}
}

// ValidateCard Func Validates traders cards if they own this company share or not
func (i *Issuer) ValidateCard(card sm.Card) bool {
	//todo: do we have a list or anything should ask from Analoui
	return true
}

// GeneralMeeting Func took place and add card to shareholders for dividend
func (i *Issuer) GeneralMeeting() error {
	// todo: execute QueryByStockSymbol create a list and loop and update dividend and dividendPayment
	return nil
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
