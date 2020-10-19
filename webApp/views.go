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

func PostHome(c *gin.Context) {
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
		trader := gmSDK.RegisterTrader(ccName, client, registerForm.AccountID)
		traderList = append(traderList, trader)
	}
	if registerForm.Org == "issuer" {
		issuer := gmSDK.RegisterIssuer(ccName, client, registerForm.Name, registerForm.AccountID)
		issuerList = append(issuerList, issuer)

	}
	////todo: return what
	//c.HTML(http.StatusOK, "register.tmpl", gin.H{
	//	"title": title,
	//})
	c.Redirect(http.StatusFound, "/home")
}
func GetView(c *gin.Context)         {}
func PostView(c *gin.Context)        {}
func GetViewCompany(c *gin.Context)  {}
func PostViewCompany(c *gin.Context) {}
func GetViewTrader(c *gin.Context)   {}
func PostViewTrader(c *gin.Context)  {}
func GetTrader(c *gin.Context)       {}
func PostTrader(c *gin.Context)      {}
func GetComapny(c *gin.Context)      {}
func PostCompany(c *gin.Context)     {}
