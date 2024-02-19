package fhandle

import (
	"DIMSProxy/protocol"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/wenzhenxi/gorsa"

	"github.com/wumansgy/goEncrypt/aes"
)

const Off = "off"

func GetUserPublicKey(request *protocol.HttpCalcRequest) (string, error) {
	return request.Payload.SellerKey, nil
}

func SymmtricKeyDecryptPlus(data string, algo string, key string) (string, error) {
	decodeString, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	xdata, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	plainText, err := aes.AesCbcDecrypt(xdata, decodeString, []byte("0000000000000000"))
	if err != nil {
		return "", nil
	}
	s := base64.StdEncoding.EncodeToString(plainText)
	return s, nil
}
func AsymmetricDecryptDoublePlus(pk string, sk string, cipherHexSymmetricKey string) (*PTBSC, error) {
	decrypt, err := gorsa.PublicDecrypt(cipherHexSymmetricKey, pk)
	if err != nil {
		return nil, err
	}

	publicDecrypt, err := gorsa.PriKeyDecrypt(decrypt, sk)
	if err != nil {
		return nil, err
	}
	var res PTBSC
	err = json.Unmarshal([]byte(publicDecrypt), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func PTBSCFiled(ptbsc *PTBSC) (string, int64, string, string, string, error) {
	if ptbsc == nil {
		return "", 0, "", "", "", errors.New("PTBSC is nil")
	}
	return ptbsc.Pwd, ptbsc.TimeStamp, ptbsc.Buyer, ptbsc.Seller, ptbsc.ContentMD5, nil
}
