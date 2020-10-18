package gmSDK

import (
	"fmt"
	sm "github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"time"
)

func Setup(user string, org string, channelName string, secret string) (*fabsdk.FabricSDK, *channel.Client, error) {

	sdk, err := fabsdk.New(config.FromFile("./config.yaml"))
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create new SDK: %s\n", err)
	}

	setupLogLevel()
	err = enrollUser(sdk, user, secret)
	if err != nil {
		fmt.Printf("can't enroll user %s\n", err)
	}
	clientChannelContext := sdk.ChannelContext(channelName, fabsdk.WithUser(user), fabsdk.WithOrg(org))

	ledgerClient, err := ledger.New(clientChannelContext)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create channel %s client %#v", channelName, err)
	}
	queryChannelInfo(ledgerClient)
	//queryChannelConfig(ledgerClient)

	client, errorClient := channel.New(clientChannelContext)

	if errorClient != nil {
		return nil, nil, fmt.Errorf("failed to create channel %s", channelName)
	}

	return sdk, client, nil
}
func setupLogLevel() {
	logging.SetLevel("fabsdk", logging.INFO)
	logging.SetLevel("fabsdk/common", logging.INFO)
	logging.SetLevel("fabsdk/fab", logging.INFO)
	logging.SetLevel("fabsdk/client", logging.INFO)
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

func enrollUser(sdk *fabsdk.FabricSDK, user string, secret string) error {
	ctx := sdk.Context()
	mspClient, err := msp.New(ctx)
	if err != nil {
		return fmt.Errorf("Failed to create msp client: %s\n", err)
	}

	_, err = mspClient.GetSigningIdentity(user)
	if err == msp.ErrUserNotFound {
		fmt.Println("Going to enroll user")
		err = mspClient.Enroll(user, msp.WithSecret(secret))
		//err = mspClient.Enroll(user)

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
	userName := "user1"
	orgName := "trader"
	channelName := "mychannel"
	secret := "user1pw"
	ccName := "stock"
	fmt.Println("setting up...")
	sdk, client, err := Setup(userName, orgName, channelName, secret)
	if err != nil {
		fmt.Println("can't setup chaincode %+v , %+v", sdk, client)
	}
	fmt.Println("register mhmmd trader..")
	mhmmd := RegisterTrader(ccName, client, "mhmmd")
	mhmmdCards := []sm.Card{{
		TraderID:    "mhmmd",
		Count:       300,
		StockSymbol: "afrouz",
		Dividend:    100,
		DividendPayments: []sm.DividendPayment{{
			Percentage: 1.8,
			PDate:      time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		}, {
			Percentage: 0.2,
			PDate:      time.Date(2020, 11, 12, 0, 0, 0, 0, time.UTC),
		}},
	}}

	fmt.Println("update cards of mhmmd %+v", mhmmdCards)
	err = mhmmd.AddCards(ccName, client, mhmmdCards)
	if err != nil {
		fmt.Printf("can't add cards of %s ,%s\n", mhmmd.TraderID, err)
	}
	mhmmdQuery, err := mhmmd.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", mhmmd.TraderID, err)
	}
	fmt.Println("Get mhmmd cards from world state:")
	fmt.Println(mhmmdQuery)
	fmt.Println("register msft issuer")
	msft := RegisterIssuer(ccName, client, "micorsoft", "msft")
	mhmmdQuery, err = mhmmd.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", mhmmd.TraderID, err)
	}
	fmt.Println("Get Cards of mhmmd after registering msft issuer:")
	fmt.Println(mhmmdQuery)

	msftPayments := []sm.DividendPayment{{
		Percentage: 0.5,
		PDate:      time.Date(2020, 10, 13, 0, 0, 0, 0, time.UTC),
	}, {
		Percentage: 0.5,
		PDate:      time.Date(2020, 10, 28, 0, 0, 0, 0, time.UTC),
	}}
	fmt.Println("Register trader Moosa")
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
	fmt.Println("Update cards of Moosa %v", moosaCards)
	err = moosa.AddCards(ccName, client, moosaCards)
	if err != nil {
		fmt.Printf("can't add cards of %s ,%s\n", moosa.TraderID, err)
	}
	moosaQuery, err := moosa.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", moosa.TraderID, err)
	}
	fmt.Println("Get Cards of Moosa from worldstate")
	fmt.Println(moosaQuery)
	fmt.Println("Holding general Meeting for msft")
	err = msft.GeneralMeeting(ccName, client, 100, msftPayments)
	if err != nil {
		fmt.Printf("generalmeeting did not hold of %s ,%s\n", msft.CompanyName, err)
	}
	msftQuery, err := msft.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", msft.CompanyName, err)
	}
	fmt.Println("Get Cards of MSFT after General Meeting")
	fmt.Println(msftQuery)
	fmt.Println("Trading 200 number of MSFT company from Moosa To Mhmmd")
	err = Trading(ccName, client, moosa.TraderID, mhmmd.TraderID, 200, msft.StockSymbol)
	if err != nil {
		fmt.Printf("can't trade of %s ,%s , %s,%s\n", moosa.TraderID, mhmmd.TraderID, msft.StockSymbol, err)
	}
	fmt.Println("Get Cards of Moosa after Trading")
	moosaQuery, err = moosa.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", moosa.TraderID, err)
	}
	fmt.Println(moosaQuery)
	fmt.Println()
	fmt.Println("Get Cards of Mhmmd after trading from World State")
	mhmmdQuery, err = mhmmd.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", mhmmd.TraderID, err)
	}
	fmt.Println(mhmmdQuery)
	msftQuery, err = msft.GetCards(ccName, client)
	if err != nil {
		fmt.Printf("can't query cards of %s ,%s\n", msft.CompanyName, err)
	}
	fmt.Println("Get Cards of MSFT after trading from World State")
	fmt.Println(msftQuery)
	defer sdk.Close()

}
