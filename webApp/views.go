package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	title = "StockChain Website"
)

type RegisterFrom struct {
	Name      string `form:"Name"`
	AccountID string `form:"AccountID"`
	Org       string `form:"Organization"`
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
func GetRegister(c *gin.Context)     {}
func PostRegister(c *gin.Context)    {}
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
