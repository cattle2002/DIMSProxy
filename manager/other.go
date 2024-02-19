package manager

import (
	"DIMSProxy/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type CertListResPayload struct {
	User      string `json:"User"`
	PublicKey string `json:"PublicKey"`
}
type CertListRes struct {
	Code int                   `json:"Code"`
	Msg  string                `json:"Msg"`
	Data *[]CertListResPayload `json:"Data"`
}

func HttpCertUserList() (*CertListRes, error) {
	url := fmt.Sprintf("http://127.0.0.1:%s/api/v1/cert/user/list", strconv.Itoa(config.Conf.Local.CertPort))
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res CertListRes
	err = json.Unmarshal(all, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
func HttpRequestOwnerPKSK() (string, string, error) {
	pem, err := config.GetPublicKeyPem()
	if err != nil {
		return "", "", err
	}
	keyPem, err := config.GetPrivateKeyPem()
	if err != nil {
		return "", "", err
	}
	return pem, keyPem, nil
	//if runtime.GOOS == "windows" {
	//	pkDir := config.Conf.Local.CurrentDir + "\\lib\\" + "public.pub"
	//	skDir := config.Conf.Local.CurrentDir + "\\lib\\" + "private.key"
	//	pk, err := os.ReadFile(pkDir)
	//	if err != nil {
	//		return "", "", err
	//	}
	//	sk, err := os.ReadFile(skDir)
	//	if err != nil {
	//		return "", "", err
	//	}
	//	return string(pk), string(sk), nil
	//}
	//if runtime.GOOS == "linux" {
	//	pkDir := config.Conf.Local.CurrentDir + "/lib/" + "public.pub"
	//	skDir := config.Conf.Local.CurrentDir + "/lib/" + "private.key"
	//	pk, err := os.ReadFile(pkDir)
	//	if err != nil {
	//		return "", "", err
	//	}
	//	sk, err := os.ReadFile(skDir)
	//	if err != nil {
	//		return "", "", err
	//	}
	//	return string(pk), string(sk), nil
	//}
	//return "", "", errors.New("not support your os")
}
