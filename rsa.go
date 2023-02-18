package fun

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"sort"
)

// 对签名数据排序、拼接
func SortingSplicing(params map[string]interface{}) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//
	str := ""
	i := 0
	for _, k := range keys {
		v := params[k]
		if i != 0 {
			str += "&"
		}
		str += k
		str += "="
		str += fmt.Sprintf("%v", v)
		i++
	}

	return str
}


func RsaSha1Sign(originalData string, RsaPrivateKey string) (sign []byte, err error) {
	h := crypto.Hash.New(crypto.SHA1) //进行SHA1的散列
	h.Write([]byte(originalData))
	hashed := h.Sum(nil)

	privateKeyString, err := base64.StdEncoding.DecodeString(RsaPrivateKey)
	if err != nil {
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyString)
	if err != nil {
		return
	}
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hashed)
	return []byte(base64.StdEncoding.EncodeToString(signatureByte)), err
}

/**RSA验签
 * $originalData待签名数据(需要先排序，然后拼接)
 * $sign需要验签的签名,需要base64_decode解码
 * 验签用支付公钥
 * return 验签是否通过 bool值
 */
func RsaSha1Verify(signatureStr, signStr, RsaPublickey string) (bool, error) {
	sign, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		return false, err
	}
	publicKey, _ := base64.StdEncoding.DecodeString(RsaPublickey) //RsaLlkPublickey  RsaPublickey
	pub, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return false, err
	}

	h := crypto.Hash.New(crypto.SHA1) //进行SHA1的散列
	h.Write([]byte(signatureStr))
	hashed := h.Sum(nil)

	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hashed, sign)
	if err != nil {
		return false, err
	}
	return true, nil
}
