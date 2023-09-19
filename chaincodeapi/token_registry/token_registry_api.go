package tokenregistry

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/spf13/viper"
)

type AccessTokenRegistry struct {
	TokenId           string `json:"token_id"`
	CertificateId     string `json:"certificate_id"`
	Owner             string `json:"owner"`
	Transferable      bool   `json:"transferable"`
	Amount            int64  `json:"amount"`
	MonthlyTokenQuota int64  `json:"monthly_token_quota"`
	AccessQuota       int64  `json:"access_quota"`
	AvailableAccesses int64  `json:"available_accesses"`
	ExpiryDate        int64  `json:"expiry_date"`
	LastUsedAt        int64  `json:"last_used_at"`
	Issuer            string `json:"issuer"`
	IssuerRef         string `json:"issuer_ref"`
	IsRevoked         bool   `json:"is_revoked"`
}

type HistoryQueryRecord struct {
	Value     *AccessTokenRegistry `json:"value"`
	TxId      string               `json:"txid"`
	Timestamp time.Time            `json:"timestamp"`
	IsDelete  bool                 `json:"is_delete"`
}

type HistoryQueryResult struct {
	TransactionID string
	Payload       []HistoryQueryRecord
}

type QueryTokenResult struct {
	TransactionID string
	Payload       AccessTokenRegistry
}

type QueryRecordsResult struct {
	TransactionID string
	Payload       []AccessTokenRegistry
}

type QueryResult struct {
	TransactionID string
	Payload       []byte
}

const (
	CHAINCODE_ID  = "token-registry"
	CHANNELID_KEY = "chaincode.token_registry.channel_id"
)

func getOrgConfig(orgId string) core.ConfigProvider {
	return config.FromFile(viper.GetString("org_hlf_config"))
}

func IssueRootToken(orgId, tokenId, certificateId, owner string) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "IssueRootToken",
		Args: [][]byte{
			[]byte(tokenId),
			[]byte(certificateId),
			[]byte(owner),
		},
	}

	return executeCC(orgId, req)
}

func IssueTransferableToken(orgId, tokenId, issuerTokenId, recipient string, amount, monthlyTokenQuota, expiryDate int64) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "IssueTransferableToken",
		Args: [][]byte{
			[]byte(tokenId),
			[]byte(issuerTokenId),
			[]byte(recipient),
			[]byte(fmt.Sprint(amount)),
			[]byte(fmt.Sprint(monthlyTokenQuota)),
			[]byte(fmt.Sprint(expiryDate)),
		},
	}

	return executeCC(orgId, req)
}

func IssueStandardToken(orgId, tokenId, issuerTokenId, recipient string, amount, accessQuota, expiryDate int64) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "IssueStandardToken",
		Args: [][]byte{
			[]byte(tokenId),
			[]byte(issuerTokenId),
			[]byte(recipient),
			[]byte(fmt.Sprint(amount)),
			[]byte(fmt.Sprint(accessQuota)),
			[]byte(fmt.Sprint(expiryDate)),
		},
	}

	return executeCC(orgId, req)
}

func ChangeTokenOwner(orgId, tokenId, owner string) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "ChangeTokenOwner",
		Args: [][]byte{
			[]byte(tokenId),
			[]byte(owner),
		},
	}

	return executeCC(orgId, req)
}

func ConsumeToken(orgId, tokenId string) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "ConsumeToken",
		Args:        [][]byte{[]byte(tokenId)},
	}

	return executeCC(orgId, req)
}

func RevokeToken(orgId, tokenId string) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "RevokeToken",
		Args:        [][]byte{[]byte(tokenId)},
	}

	return executeCC(orgId, req)
}

func QueryToken(orgId, tokenId string) (*QueryTokenResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "QueryToken",
		Args:        [][]byte{[]byte(tokenId)},
	}

	res, err := queryCC(orgId, req)
	if err != nil {
		return nil, err
	}

	parsedResult := &QueryTokenResult{
		TransactionID: res.TransactionID,
	}

	err = json.Unmarshal(res.Payload, &parsedResult.Payload)
	if err != nil {
		return nil, err
	}

	return parsedResult, nil
}

func QueryTokenStatus(orgId, tokenId string) (string, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "QueryTokenStatus",
		Args:        [][]byte{[]byte(tokenId)},
	}

	res, err := queryCC(orgId, req)
	if err != nil {
		return "", err
	}

	return string(res.Payload), nil
}

func QueryRecords(orgId, queryString string) (*QueryRecordsResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "QueryRecords",
		Args:        [][]byte{[]byte(queryString)},
	}

	res, err := queryCC(orgId, req)
	if err != nil {
		return nil, err
	}

	parsedResult := &QueryRecordsResult{
		TransactionID: res.TransactionID,
	}

	err = json.Unmarshal(res.Payload, &parsedResult.Payload)
	if err != nil {
		return nil, err
	}

	return parsedResult, nil
}

func GetHistoryForKey(orgId, tokenId string) (*HistoryQueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "GetHistoryForKey",
		Args:        [][]byte{[]byte(tokenId)},
	}

	res, err := queryCC(orgId, req)
	if err != nil {
		return nil, err
	}

	parsedResult := &HistoryQueryResult{
		TransactionID: res.TransactionID,
	}

	err = json.Unmarshal(res.Payload, &parsedResult.Payload)
	if err != nil {
		return nil, err
	}

	return parsedResult, nil
}

// Read operation on chaincode
func queryCC(orgId string, req *channel.Request) (*QueryResult, error) {
	sdk, err := fabsdk.New(getOrgConfig(orgId))
	if err != nil {
		return nil, err
	}

	defer sdk.Close()

	channelId := viper.GetString(CHANNELID_KEY)
	orgAdminUserKey := fmt.Sprintf("organizations.%s.org_admin", orgId)
	orgAdminUser := viper.GetString(orgAdminUserKey)
	orgNameKey := fmt.Sprintf("organizations.%s.org_name", orgId)
	orgName := viper.GetString(orgNameKey)

	ctx := sdk.ChannelContext(
		channelId,
		fabsdk.WithUser(orgAdminUser),
		fabsdk.WithOrg(orgName))

	client, err := channel.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new resource management client: %s", err)
	}

	response, err := client.Query(*req)
	if err != nil {
		return nil, err
	}

	res := &QueryResult{
		TransactionID: string(response.TransactionID),
		Payload:       response.Payload,
	}

	return res, nil
}

// Write operation on chaincode
func executeCC(orgId string, req *channel.Request) (*QueryResult, error) {
	sdk, err := fabsdk.New(getOrgConfig(orgId))
	if err != nil {
		return nil, err
	}

	defer sdk.Close()

	channelId := viper.GetString(CHANNELID_KEY)
	orgAdminUserKey := fmt.Sprintf("organizations.%s.org_admin", orgId)
	orgAdminUser := viper.GetString(orgAdminUserKey)
	orgNameKey := fmt.Sprintf("organizations.%s.org_name", orgId)
	orgName := viper.GetString(orgNameKey)

	ctx := sdk.ChannelContext(
		channelId,
		fabsdk.WithUser(orgAdminUser),
		fabsdk.WithOrg(orgName))

	client, err := channel.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new resource management client: %s", err)
	}

	response, err := client.Execute(*req)
	if err != nil {
		return nil, err
	}

	res := &QueryResult{
		TransactionID: string(response.TransactionID),
		Payload:       response.Payload,
	}

	return res, nil
}
