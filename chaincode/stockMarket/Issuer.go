package stockMarket

import "google.golang.org/genproto/googleapis/type/date"

type Issuer struct {
	CompanyName string `json:"companyName"`
	StockSymbol string `json:"stockSymbol"`
}

func (i *Issuer) ValidateCard(card Card) bool {
	//validate Card
	return false
}

func (i *Issuer) GeneralMeeting() {
	cards := QueryCardByIssuer(*i)
	for _, card := range cards {
		//updateCard(nil,nil)
		//get trader and update cardlist
		// do we need cardstatus?
	}
}

func (i *Issuer) PayCard(payDate date.Date) {
	cards := QueryCardByIssuer(*i)
	for _, card := range cards {
		for _, dDate := range card.DividendPayment {
			if dDate.PDate.String() == payDate.String() {
				dDate.Paid = true

			}
		}
		//card.UpdateCard()
	}

}
