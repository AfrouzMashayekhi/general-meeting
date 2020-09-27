package generalMeetingSDK

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"os"
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
	//	//todo: how to get traderID or compnayname
	//	RegisterTrader(sdk, client, secret, user)
	//
	//} else if org == "company" {
	//	//todo: hardocoded names
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

func enrollUser(sdk *fabsdk.FabricSDK, user string, secret string) {
	ctx := sdk.Context()
	mspClient, err := msp.New(ctx)
	if err != nil {
		fmt.Printf("Failed to create msp client: %s\n", err)
	}

	_, err = mspClient.GetSigningIdentity(user)
	if err == msp.ErrUserNotFound {
		fmt.Println("Going to enroll user")
		err = mspClient.Enroll(user, msp.WithSecret(secret))

		if err != nil {
			fmt.Printf("Failed to enroll user: %s\n", err)
		} else {
			fmt.Printf("Success enroll user: %s\n", user)
		}

	} else if err != nil {
		fmt.Printf("Failed to get user: %s\n", err)
	} else {
		fmt.Printf("User %s already enrolled, skip enrollment.\n", user)
	}

}
func main() {
	// todo:get input?
	//todo: call setup
	//todo:call enroll user
	//todo: call registeration

}
