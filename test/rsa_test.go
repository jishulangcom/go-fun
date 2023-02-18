package test

import (
	"encoding/json"
	"fmt"
	"github.com/jishulangcom/go-fun"
	"testing"
)

// Transaction Submit KEY(自家站点)
var transactionSubmitPrivateKey = "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBALu4z/1kcTys46RTSBgS1Ie+f+AJZXPHpFObXw/UsdCFxHng26ilRvTAGwptb14RunxinOPB1ALhwAPGBZBi/EQB9azhxsdNQquWqfgCNgMd87wLA+YFRwid/EQitJJJgpc3K8PhRh/stHQfxFnALPTarh0aeE4HL0gtP1WG92c3AgMBAAECgYEAnBFcqlhVZJKAd9/dclZFZ83TVpm5RPbYAcKM2AaHBswPuxxcwusWAOmuEY9GJbkrh7ocoaZF81doYjrB6XbCLwE5ZO5wC/awuzlChIUEYbuKHYaHklOIRXZdVB5F1VWNuigm9+Bp1Z1tgMPJ3XuPh8vcLsBjYzoW/n52KZjWX/ECQQD2vi8IJvOoq3VTMguBsX1C/WuIP7i6jS8p794SyYj7OlZ5lCDj38/YxSGbTL/1sOQ2hTPqfPjnOWbkIRfA/8srAkEAwsPElz/V3dVYQJDm4B4aOTbA5RJWHcSyvCzme3J50V+f1n+NYT5A7vsfuwD/EtOg4JUAK2oKEJYJx4qDIJAeJQJBAJsDjp6rggakeVgkJ7B6Jnzwox79EXw5+Lh7FuRsst9KnktRcXxX/sdryZo8lJixYh6SfrRBgUoa+PY3iCnbSfMCQQCY9oXcwB7olZk1RFh/JkU0MCN5BVirEoPJtH3j8DlqTe7L2OuhtvHqf0IhbZvnHlyYZY13i1WSRzQxCkEH/wgdAkB8T84E5yKQ4n0PoigggVrCv3SekwZMOxtka5/KfH9ioJOq7XQZN5nZHxX5xIFmngiowOiAIJ/Hpnv72/Oi3Ns+"
var transactionSubmitPublicKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC7uM/9ZHE8rOOkU0gYEtSHvn/gCWVzx6RTm18P1LHQhcR54NuopUb0wBsKbW9eEbp8YpzjwdQC4cADxgWQYvxEAfWs4cbHTUKrlqn4AjYDHfO8CwPmBUcInfxEIrSSSYKXNyvD4UYf7LR0H8RZwCz02q4dGnhOBy9ILT9VhvdnNwIDAQAB"

// Transaction Notification KEY(pay平台)
var transactionNotificationPrivateKey = "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBALu4z/1kcTys46RTSBgS1Ie+f+AJZXPHpFObXw/UsdCFxHng26ilRvTAGwptb14RunxinOPB1ALhwAPGBZBi/EQB9azhxsdNQquWqfgCNgMd87wLA+YFRwid/EQitJJJgpc3K8PhRh/stHQfxFnALPTarh0aeE4HL0gtP1WG92c3AgMBAAECgYEAnBFcqlhVZJKAd9/dclZFZ83TVpm5RPbYAcKM2AaHBswPuxxcwusWAOmuEY9GJbkrh7ocoaZF81doYjrB6XbCLwE5ZO5wC/awuzlChIUEYbuKHYaHklOIRXZdVB5F1VWNuigm9+Bp1Z1tgMPJ3XuPh8vcLsBjYzoW/n52KZjWX/ECQQD2vi8IJvOoq3VTMguBsX1C/WuIP7i6jS8p794SyYj7OlZ5lCDj38/YxSGbTL/1sOQ2hTPqfPjnOWbkIRfA/8srAkEAwsPElz/V3dVYQJDm4B4aOTbA5RJWHcSyvCzme3J50V+f1n+NYT5A7vsfuwD/EtOg4JUAK2oKEJYJx4qDIJAeJQJBAJsDjp6rggakeVgkJ7B6Jnzwox79EXw5+Lh7FuRsst9KnktRcXxX/sdryZo8lJixYh6SfrRBgUoa+PY3iCnbSfMCQQCY9oXcwB7olZk1RFh/JkU0MCN5BVirEoPJtH3j8DlqTe7L2OuhtvHqf0IhbZvnHlyYZY13i1WSRzQxCkEH/wgdAkB8T84E5yKQ4n0PoigggVrCv3SekwZMOxtka5/KfH9ioJOq7XQZN5nZHxX5xIFmngiowOiAIJ/Hpnv72/Oi3Ns+"
var transactionNotificationPublicKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC7uM/9ZHE8rOOkU0gYEtSHvn/gCWVzx6RTm18P1LHQhcR54NuopUb0wBsKbW9eEbp8YpzjwdQC4cADxgWQYvxEAfWs4cbHTUKrlqn4AjYDHfO8CwPmBUcInfxEIrSSSYKXNyvD4UYf7LR0H8RZwCz02q4dGnhOBy9ILT9VhvdnNwIDAQAB"


type transactionSubmitDto struct {
	Currency string `json:"currency"`
	amouAmountnt   string `json:"amount"`
	Sign     string `json:"sign"`
}

type transactionNotificationDto struct {
	Id int64 `json:"id"`
	Currency string `json:"currency"`
	amouAmountnt   string `json:"amount"`
	Sign     string `json:"sign"`
}

func rsa_sha1_sign(privateKey string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	data["currency"] = "USDT"
	data["amount"] = "100"

	origData := fun.SortingSplicing(data)
	signStr, err := fun.RsaSha1Sign(origData, privateKey)
	if err != nil {
		return nil, err
	}
	data["sign"] = string(signStr)

	//Query
	fmt.Println("Query data：", data)
	return data, err
}

// Transaction Submit
func Test_rsa_sha1_sign(t *testing.T) {
	data, err := rsa_sha1_sign(transactionSubmitPrivateKey)
	fmt.Println(data, err)
}

// Transaction Notification
func Test_rsa_sha1_verify(t *testing.T) {
	//模拟取得Notification数据
	originalData, _ := rsa_sha1_sign(transactionNotificationPrivateKey)
	originalDataByte, _ := json.Marshal(originalData)
	fmt.Println("originalData json：", string(originalDataByte))

	//verify开始
	rsaData := transactionNotificationDto{}
	json.Unmarshal(originalDataByte, &rsaData)
	delete(originalData, "sign")
	signatureStr := fun.SortingSplicing(originalData)
	ok, err := fun.RsaSha1Verify(signatureStr, rsaData.Sign, transactionNotificationPublicKey)
	if err != nil {
		fmt.Println("RsaSha1Verify err：", err)
		return
	}
	fmt.Println(ok)
}
