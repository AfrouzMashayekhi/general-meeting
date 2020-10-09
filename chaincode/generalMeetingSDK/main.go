package main

import (
	"fmt"
	sm "github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
	"time"
)

func setup(user string, org string, channelName string) (*fabsdk.FabricSDK, *channel.Client, error) {
	//if len(os.Args) != 5 {
	//	fmt.Printf("Usage: main.go <user-name> <user-secret> <org> <channel> <chaincode-name>\n")
	//	os.Exit(1)
	//}
	sdk, err := fabsdk.New(config.FromFile("./config.yaml"))
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s\n", err)
		os.Exit(1)
	}
	defer sdk.Close()

	clientChannelContext := sdk.ChannelContext(channelName, fabsdk.WithUser(user), fabsdk.WithOrg(org))
	ledgerClient, err := ledger.New(clientChannelContext)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create channel %s client %#v", channelName, err)
	}
	queryChannelInfo(ledgerClient)
	queryChannelConfig(ledgerClient)
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create channel %s", channelName)
	}
	//if org == "trader" {
	//	//todo: how to get traderID or companyName
	//	RegisterTrader(sdk, client, secret, user)
	//
	//} else if org == "company" {
	//	//todo: hardocod names
	//	RegisterIssuer(sdk, client, secret, user)
	//} else {
	//	return cc, nil, fmt.Errorf("org name not valid")
	//}
	return sdk, client, nil
}

func queryChannelConfig(ledgerClient *ledger.Client) {
	resp1, err := ledgerClient.QueryConfig()
	if err != nil {
		fmt.Printf("Failed to queryConfig: %s", err)
	}
	fmt.Println("ChannelID: ", resp1.ID())
	fmt.Println("Channel Orderers: ", resp1.Orderers())
	fmt.Println("Channel Versions: ", resp1.Versions())
}

func queryChannelInfo(ledgerClient *ledger.Client) {
	resp, err := ledgerClient.QueryInfo()
	if err != nil {
		fmt.Printf("Failed to queryInfo: %s", err)
	}
	fmt.Println("BlockChainInfo:", resp.BCI)
	fmt.Println("Endorser:", resp.Endorser)
	fmt.Println("Status:", resp.Status)
}

func enrollUser(sdk *fabsdk.FabricSDK, user string) error {
	ctx := sdk.Context()
	mspClient, err := msp.New(ctx)
	if err != nil {
		return fmt.Errorf("Failed to create msp client: %s\n", err)
	}

	_, err = mspClient.GetSigningIdentity(user)
	if err == msp.ErrUserNotFound {
		fmt.Println("Going to enroll user")
		//err = mspClient.Enroll(user, msp.WithSecret(secret))
		err = mspClient.Enroll(user)

		if err != nil {
			return fmt.Errorf("Failed to enroll user: %s\n", err)
		} else {
			fmt.Printf("Success enroll user: %s\n", user)
		}

	} else if err != nil {
		return fmt.Errorf("Failed to get user: %s\n", err)
	} else {
		fmt.Printf("User %s already enrolled, skip enrollment.\n", user)
	}
	return nil
}
func main() {
	userName := "User1"
	orgName := "trader"
	ccName := "stock"
	//todo: call setup
	fmt.Println("setting up...")
	sdk, client, err := setup(userName, orgName, ccName)
	if err != nil {
		fmt.Println("can't setup chaincode")
	}
	err = enrollUser(sdk, userName)
	if err != nil {
		fmt.Printf("can't enroll user %s\n", err)
	}

	mhmmd := RegisterTrader(ccName, client, "mhmmd")
	mhmmdCards := []sm.Card{{
		TraderID:    "mhmmd",
		Count:       300,
		StockSymbol: "afrouz",
		Dividend:    100,
		DividendPayments: []sm.DividendPayment{{
			Percentage: 0.8,
			PDate:      time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		}, {
			Percentage: 0.2,
			PDate:      time.Date(2020, 11, 12, 0, 0, 0, 0, time.UTC),
		}},
	}}
	err = mhmmd.AddCards(ccName, client, mhmmdCards)
	if err != nil {
		fmt.Printf("can't add cards of %s ,%s\n", mhmmd.TraderID, err)
	}
	mhmmdQuery, err := mhmmd.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", mhmmd.TraderID, err)
	}
	fmt.Println(mhmmdQuery)
	msft := RegisterIssuer(ccName, client, "micorsoft", "msft")
	mhmmdQuery, err = mhmmd.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", mhmmd.TraderID, err)
	}
	fmt.Println(mhmmdQuery)

	msftPayments := []sm.DividendPayment{{
		Percentage: 0.5,
		PDate:      time.Date(2020, 10, 13, 0, 0, 0, 0, time.UTC),
	}, {
		Percentage: 0.5,
		PDate:      time.Date(2020, 10, 28, 0, 0, 0, 0, time.UTC),
	}}
	err = msft.GeneralMeeting(ccName, client, 100, msftPayments)
	if err != nil {
		fmt.Printf("generalmeeting did not hold of %s ,%s\n", msft.CompanyName, err)
	}
	msftQuery, err := msft.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", msft.CompanyName, err)
	}
	fmt.Println(msftQuery)
	moosa := RegisterTrader(ccName, client, "moosa")
	moosaCards := []sm.Card{{
		TraderID:    "moosa",
		Count:       500,
		StockSymbol: "msft",
		Dividend:    200,
		DividendPayments: []sm.DividendPayment{{
			Percentage: 0.6,
			PDate:      time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		}, {
			Percentage: 0.4,
			PDate:      time.Date(2020, 11, 12, 0, 0, 0, 0, time.UTC),
		}},
	}}
	err = moosa.AddCards(ccName, client, moosaCards)
	if err != nil {
		fmt.Printf("can't add cards of %s ,%s\n", moosa.TraderID, err)
	}
	moosaQuery, err := moosa.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", moosa.TraderID, err)
	}
	fmt.Println(moosaQuery)
	msftQuery, err = msft.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", msft.CompanyName, err)
	}
	fmt.Println(msftQuery)
	err = Trading(ccName, client, moosa.TraderID, mhmmd.TraderID, 200, msft.StockSymbol)
	if err != nil {
		fmt.Printf("can't trade of %s ,%s , %s,%s\n", moosa.TraderID, mhmmd.TraderID, msft.StockSymbol, err)
	}
	moosaQuery, err = moosa.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", moosa.TraderID, err)
	}
	fmt.Println(moosaQuery)
	mhmmdQuery, err = mhmmd.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", mhmmd.TraderID, err)
	}
	fmt.Println(mhmmdQuery)
	msftQuery, err = msft.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", msft.CompanyName, err)
	}
	fmt.Println(msftQuery)

}
