package sdk

import (
	"time"
)

// Issuer is a company that validate and pays Dividend and holds General Meetings
type Issuer struct {
	CompanyName string `json:"companyName"`
	StockSymbol string `json:"stockSymbol"`
}

// ValidateCard Func Validates traders cards if they own this company share or not
func (i *Issuer) ValidateCard(card Card) bool {
	//todo:validate Card
	return false
}

// GeneralMeeting Func took place and add card to shareholders for dividend
func (i *Issuer) GeneralMeeting() {
	cards := QueryByIssuer(*i)
	for _, card := range cards {
		//updateCard(nil,nil)
		//get trader and update cardlist
		// do we need cardstatus?
	}
}

// PayCard Func at the time of payDate
func (i *Issuer) PayCard(payDate time.Time) {
	cards := QueryByIssuer(*i)
	for _, card := range cards {
		for _, dDate := range card.DividendPayments {
			if dDate.PDate.Equal(payDate) {
				dDate.Paid = true

			}
		}
		//card.UpdateCard()
	}

}
