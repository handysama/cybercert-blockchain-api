package certificatetemplate

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

type CertificateTemplate struct {
	TemplateSource interface{} `json:"template_source"`
	SourceType     string      `json:"source_type"`
	Version        string      `json:"version"`
	IssuerId       string      `json:"issuer_id"`
	IssuerName     string      `json:"issuer_name"`
}

type HistoryQueryRecord struct {
	Value     *CertificateTemplate `json:"value"`
	TxId      string               `json:"txid"`
	Timestamp time.Time            `json:"timestamp"`
	IsDelete  bool                 `json:"is_delete"`
}

type HistoryQueryResult struct {
	TransactionID string
	Payload       []HistoryQueryRecord
}

type QueryTemplateResult struct {
	TransactionID string
	Payload       CertificateTemplate
}

type QueryResult struct {
	TransactionID string
	Payload       []byte
}

const (
	CHAINCODE_ID  = "certificate_template"
	CHANNELID_KEY = "chaincode.certificate_template.channel_id"
)

func getOrgConfig(orgId string) core.ConfigProvider {
	return config.FromFile(viper.GetString("org_hlf_config"))
}

func PutTemplate(orgId, templateKey string, templateSource []byte,
	sourceType, version, issuerId, issuerName string) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "PutTemplate",
		Args: [][]byte{
			[]byte(templateKey),
			templateSource,
			[]byte(sourceType),
			[]byte(version),
			[]byte(issuerId),
			[]byte(issuerName),
		},
	}

	return executeCC(orgId, req)
}

func QueryTemplate(orgId, templateKey string) (*QueryTemplateResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "QueryTemplate",
		Args:        [][]byte{[]byte(templateKey)},
	}

	res, err := queryCC(orgId, req)
	if err != nil {
		return nil, err
	}

	parsedResult := &QueryTemplateResult{
		TransactionID: res.TransactionID,
	}

	err = json.Unmarshal(res.Payload, &parsedResult.Payload)
	if err != nil {
		return nil, err
	}

	return parsedResult, nil
}

func GetHistoryForKey(orgId, templateKey string) (*HistoryQueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "GetHistoryForKey",
		Args:        [][]byte{[]byte(templateKey)},
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
