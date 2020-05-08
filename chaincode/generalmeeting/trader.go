package generalmeeting

type Trader interface {
	/*buy sale or aniything that a trader do*/
	func () BuyShare()
	func () SellShare()
	// update share data set
	func () SetShare ()
	func () GetShare ()
	//query specific share of trader
	func () QueryTraderShare ()
	//query all shares of trader
	func () QueryAllTraderShares ()
}