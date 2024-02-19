package manager

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"DIMSProxy/model"
	"DIMSProxy/protocol"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func InertSellerPK2DBBak(user string, timeStamp int64, pk string, sk string) error {
	url := fmt.Sprintf("http://127.0.0.1:%s/api/v1/cert2f", strconv.Itoa(config.Conf.Local.CertPort))
	var req protocol.Cert2fReq
	var res protocol.Cert2fRes
	req.User = user
	req.TimeStamp = timeStamp
	req.PublicKey = pk
	req.PrivateKey = sk
	marshal, _ := json.Marshal(req)
	resp, err := http.Post(url, "application/json", bytes.NewReader(marshal))
	if err != nil {
		//log.Loggerx.Errorf("request http://127.0.0.1:%s/api/v1/cert2f error:%s",err.Error())
		return err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(all, &res)
	if err != nil {
		return err
	}
	if res.Code != protocol.FSuccessCode {
		return errors.New(res.Msg)
	} else {
		return nil
	}
}
func InertSellerPK2DB(user string, timeStamp int64, pk string, sk string) error {
	err := model.CreateCert(user, timeStamp, pk, sk)
	return err
}
func Watcher() {
	//todo 初始化库的位置
	//NewAlgoLib()
	http.HandleFunc("/api/v1/websocket/1", Watch)
	addr := config.Conf.Local.EthHost + ":" + strconv.Itoa(config.Conf.Local.ManagerServicePort)
	log.Logger.Infof("manager service addr:%s/api/v1/websocket/1", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Logger.Errorf("manager service run failed:%s", err.Error())
		return
	}
}
