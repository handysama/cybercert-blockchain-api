package certificateinfo

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

type CertificateRecord struct {
	CertificateSignature string      `json:"certificate_signature"`
	TemplateRef          string      `json:"template_ref"`
	CourseName           string      `json:"course_name"`
	ModuleName           string      `json:"module_name"`
	CertificateHolder    string      `json:"certificate_holder"`
	Email                string      `json:"email"`
	IsRevoked            bool        `json:"is_revoked"`
	IssuerId             string      `json:"issuer_id"`
	IssuerName           string      `json:"issuer_name"`
	IssuedAt             string      `json:"issued_at"`
	Extras               interface{} `json:"extras"`
}

type HistoryQueryRecord struct {
	Value     *CertificateRecord `json:"value"`
	TxId      string             `json:"txid"`
	Timestamp time.Time          `json:"timestamp"`
	IsDelete  bool               `json:"is_delete"`
}

type HistoryQueryResult struct {
	TransactionID string
	Payload       []HistoryQueryRecord
}

type QueryCertificateResult struct {
	TransactionID string
	Payload       CertificateRecord
}

type QueryRecordsResult struct {
	TransactionID string
	Payload       []CertificateRecord
}

type QueryResult struct {
	TransactionID string
	Payload       []byte
}

const (
	CHAINCODE_ID  = "certificate-info"
	CHANNELID_KEY = "chaincode.certificate_info.channel_id"
)

func getOrgConfig(orgId string) core.ConfigProvider {
	return config.FromFile(viper.GetString("org_hlf_config"))
}

func IssueCertificate(orgId, certificateId, certSignature, templateRef,
	courseName, moduleName, certHolder, email, issuerId,
	issuer, issuedAt string, extras []byte) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "IssueCertificate",
		Args: [][]byte{
			[]byte(certificateId),
			[]byte(certSignature),
			[]byte(templateRef),
			[]byte(courseName),
			[]byte(moduleName),
			[]byte(certHolder),
			[]byte(email),
			[]byte(issuerId),
			[]byte(issuer),
			[]byte(issuedAt),
			extras,
		},
	}

	return executeCC(orgId, req)
}

func QueryCertificate(orgId, certificateId string) (*QueryCertificateResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "QueryCertificate",
		Args:        [][]byte{[]byte(certificateId)},
	}

	res, err := queryCC(orgId, req)
	if err != nil {
		return nil, err
	}

	parsedResult := &QueryCertificateResult{
		TransactionID: res.TransactionID,
	}

	err = json.Unmarshal(res.Payload, &parsedResult.Payload)
	if err != nil {
		return nil, err
	}

	return parsedResult, nil
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

func RevokeCertificate(orgId, certificateId string) (*QueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "RevokeCertificate",
		Args:        [][]byte{[]byte(certificateId)},
	}

	return executeCC(orgId, req)
}

func GetHistoryForKey(orgId, certificateId string) (*HistoryQueryResult, error) {
	req := &channel.Request{
		ChaincodeID: CHAINCODE_ID,
		Fcn:         "GetHistoryForKey",
		Args:        [][]byte{[]byte(certificateId)},
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
