package main

import (
	gmSDK "github.com/afrouzMashaykhi/general-meeting/chaincode/generalMeetingSDK"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	title = "StockChain Website"
)

type RegisterForm struct {
	Name      string `form:"name" binding:"required"`
	AccountID string `form:"accountID" binding:"required"`
	Org       string `form:"organization" binding:"required`
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
	// if bind was ok detect org and register
	if registerForm.Org == "trader" {
		gmSDK.RegisterTrader()
	}
	if registerForm.Org == "issuer" {

	}
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
