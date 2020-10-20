package main

import (
	"fmt"
	gmSDK "github.com/afrouzMashaykhi/general-meeting/chaincode/generalMeetingSDK"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

var (
	port        = ":8080"
	userSDK     = "user1"
	orgSDK      = "trader"
	channelName = "mychannel"
	secret      = "user1pw"
	ccName      = "stock"
	client      *channel.Client
	sdk         *fabsdk.FabricSDK
)

func main() {
	//setup sdk
	fmt.Println("setting up...")
	var err error
	sdk, client, err = gmSDK.Setup(userSDK, orgSDK, channelName, secret)
	if err != nil {
		fmt.Println("can't setup chaincode %+v , %+v", sdk, client)
	}
	// Creates a router without any middleware by default
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())
	// Serve static files
	app.Static("/assets", "./assets")
	app.GET("/view/t/:user", GetTrader)
	app.GET("/view/c/:company", GetComapny)
	app.GET("/home", GetHome)
	app.GET("/view/t", GetViewTrader)
	app.GET("/view/c", GetViewCompany)
	app.GET("/register", GetRegister)
	app.POST("/register", PostRegister)
	// Listen and serve on 0.0.0.0:8080
	app.Run(port)
}
