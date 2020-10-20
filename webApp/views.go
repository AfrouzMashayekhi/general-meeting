package main

import (
	"fmt"
	gmSDK "github.com/afrouzMashaykhi/general-meeting/chaincode/generalMeetingSDK"
	_ "github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.HTML(http.StatusOK, "trader_profile.tmpl", gin.H{
		"title": title,
	})
}
func PostTraderAddCard(c *gin.Context) {
	traderID := c.Param("trader")
}
func PostTraderTrade(c *gin.Context) {
	traderID := c.Param("trader")
}
func GetComapny(c *gin.Context) {
	c.HTML(http.StatusOK, "issuer_profile.tmpl", gin.H{
		"title": title,
	})
}
func PostCompanyGenralMeeting(c *gin.Context) {
	stockSymbol := c.Param("compnay")

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
