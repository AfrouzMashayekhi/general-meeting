package main

import (
	"fmt"
	gmSDK "github.com/afrouzMashaykhi/general-meeting/chaincode/generalMeetingSDK"
	sm "github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var (
	title      = "StockChain Website"
	traderList []*gmSDK.Trader
	issuerList []*gmSDK.Issuer
)

type RegisterForm struct {
	Name      string `form:"name" binding:"required"`
	AccountID string `form:"accountID" binding:"required"`
	Org       string `form:"org" binding:"required`
}
type GeneralMeetingForm struct {
	StockSymbol string      `form:"issuerName" binding:"required"`
	Dividend    int         `form:"dividend" binding:"required"`
	Percentages []float32   `form:"percentage[]" binding:"required`
	Times       []time.Time `form:"time[]" binding:"required`
}
type TraderCardForm struct {
	TraderID    string      `form:"traderID" binding:"required"`
	StockSymbol string      `form:"stockSymbol" binding:"required"`
	Dividend    int         `form:"dividend" binding:"required"`
	Count       int         `form:"count" binding:"required"`
	Percentages []float32   `form:"percentage[]" binding:"required`
	Times       []time.Time `form:"time[]" binding:"required`
}
type TradeForm struct {
	Seller      string `form:"seller" binding:"required"`
	Buyer       string `form:"buyer" binding:"required"`
	StockSymbol string `form:"tradeStockSymbol" binding:"required"`
	Count       int    `form:"sellCount" binding:"required"`
}

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": title,
	})
}

func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{
		"title": title,
	})
}
func PostRegister(c *gin.Context) {
	var registerForm RegisterForm
	if err := c.ShouldBind(&registerForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("in the post")
	// if bind was ok detect org and register
	if registerForm.Org == "trader" {
		trader := gmSDK.RegisterTrader(ccName, client, registerForm.Name, registerForm.AccountID)
		traderList = append(traderList, trader)
		c.Redirect(http.StatusFound, "/view/t")

	}
	if registerForm.Org == "issuer" {
		issuer := gmSDK.RegisterIssuer(ccName, client, registerForm.Name, registerForm.AccountID)
		issuerList = append(issuerList, issuer)
		c.Redirect(http.StatusFound, "/view/c")

	}

}
func GetViewCompany(c *gin.Context) {
	getViewList(c, "issuer")
}
func GetViewTrader(c *gin.Context) {
	getViewList(c, "trader")
}
func GetTrader(c *gin.Context) {
	queriedTID := c.Param("trader")
	queriedTID = strings.Replace(queriedTID, ":", "", -1)
	indexTrader := -1
	for i, traders := range traderList {
		traders.TraderName = queriedTID
		indexTrader = i
		break
	}
	items, err := traderList[indexTrader].GetCards(ccName, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "trader_profile.tmpl", gin.H{
		"title":    title,
		"traderID": queriedTID,
		"items":    items.Cards,
	})
}
func PostTrader(c *gin.Context) {
	queriedTID := c.Param("trader")
	queriedTID = strings.Replace(queriedTID, ":", "", -1)
	//POST ADD CARD
	var traderCardForm TraderCardForm
	//khob alan man 2 ta form ro check mikonam chejor i joda   na joda
	indexTrader := -1
	if c.PostForm("formName") == "Add" {
		if err := c.ShouldBind(&traderCardForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, traders := range traderList {
			traders.TraderName = queriedTID
			indexTrader = i
			break
		}
		//validate DividendPayments
		var inPayments []sm.DividendPayment
		for j, percentage := range traderCardForm.Percentages {
			dpay := sm.DividendPayment{
				Percentage: percentage,
				PDate:      traderCardForm.Times[j],
			}
			inPayments = append(inPayments, dpay)
		}
		addingCard := sm.Card{
			TraderID:         traderCardForm.TraderID,
			Count:            traderCardForm.Count,
			StockSymbol:      traderCardForm.StockSymbol,
			Dividend:         traderCardForm.Dividend,
			DividendPayments: inPayments,
		}
		err := traderList[indexTrader].AddCards(ccName, client, addingCard)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	//POST TRADE

	var tradeForm TradeForm
	if c.PostForm("formName") == "Trade" {
		if err := c.ShouldBind(&tradeForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := gmSDK.Trading(ccName, client, tradeForm.Seller, tradeForm.Buyer, tradeForm.Count, tradeForm.StockSymbol)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	// GET CARD
	items, err := traderList[indexTrader].GetCards(ccName, client)
	//yani for?
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "trader_profile.tmpl", gin.H{
		"title":    title,
		"traderID": queriedTID,
		"items":    items.Cards,
	})
}

func GetComapny(c *gin.Context) {
	queriedSS := c.Param("company")
	queriedSS = strings.Replace(queriedSS, ":", "", -1)
	indexIssuer := -1
	for i, issuers := range issuerList {
		issuers.StockSymbol = queriedSS
		indexIssuer = i
		break
	}
	items, err := issuerList[indexIssuer].GetCards(ccName, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "issuer_profile.tmpl", gin.H{
		"title":      title,
		"issuerName": queriedSS,
		"items":      items.Cards,
	})
}
func PostCompanyGenralMeeting(c *gin.Context) {
	queriedSS := c.Param("company")
	queriedSS = strings.Replace(queriedSS, ":", "", -1)
	var gMeetingForm GeneralMeetingForm
	//POST GENERAL MEETING Part
	if err := c.ShouldBind(&gMeetingForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	indexIssuer := -1
	for i, issuers := range issuerList {
		issuers.StockSymbol = queriedSS
		indexIssuer = i
		break
	}
	//validate DividendPayments
	var inPayments []sm.DividendPayment
	for j, percentage := range gMeetingForm.Percentages {
		dpay := sm.DividendPayment{
			Percentage: percentage,
			PDate:      gMeetingForm.Times[j],
		}
		inPayments = append(inPayments, dpay)
	}
	err := issuerList[indexIssuer].GeneralMeeting(ccName, client, gMeetingForm.Dividend, inPayments)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//GET CARDS PART
	items, err := issuerList[indexIssuer].GetCards(ccName, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "issuer_profile.tmpl", gin.H{
		"title":      title,
		"issuerName": queriedSS,
		"items":      items.Cards,
	})

}

func getViewList(c *gin.Context, listType string) {
	if listType == "trader" {
		listtype_title := "Traders"
		idfieldname := "TraderID"
		items := traderList
		c.HTML(http.StatusOK, "listview.tmpl", gin.H{
			"listtype":    listtype_title,
			"title":       "Manage Entities",
			"idfieldname": idfieldname,
			"items":       items,
		})
	} else {
		listtype_title := "Companies"
		idfieldname := "StockSymbol"
		items := issuerList
		c.HTML(http.StatusOK, "listview.tmpl", gin.H{
			"listtype":    listtype_title,
			"title":       "Manage Entities",
			"idfieldname": idfieldname,
			"items":       items,
		})
	}
}
