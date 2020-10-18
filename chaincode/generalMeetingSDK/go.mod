module github.com/afrouzMashaykhi/general-meeting/chaincode/generalMeetingSDK

go 1.14

replace github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket => /home/afrouz/go/src/fabric-samples/chaincode/stockmarket

require (
	github.com/afrouzMashaykhi/general-meeting/chaincode/stockmarket v0.0.0-20201003060308-319b1140efd8
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta3
)
