package manager

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func convert(cls []CertListResPayload) *[]protocol.CertShowResPayload {
	var csl []protocol.CertShowResPayload
	for i := 0; i < len(cls); i++ {
		var ele protocol.CertShowResPayload
		ele.User = cls[i].User
		ele.PublicKey = cls[i].PublicKey
		csl = append(csl, ele)
	}
	return &csl
}

func CertShowHandle() {
	list, err := HttpCertUserList()
	if err != nil {
		CertShowResponseError(err.Error())
		return
	}
	var res protocol.CertShowRes
	if list.Code != protocol.FSuccessCode {
		res.IpAddr = config.Conf.Local.EthHost
		res.Code = protocol.FErrorCode
		res.Cmd = string(CertShowRet)
		res.Msg = list.Msg
		marshal, _ := json.Marshal(res)
		err = ManagerConn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			log.Logger.Errorf("return msg to manager error:%s", err.Error())
			return
		}
	} else {
		res.IpAddr = config.Conf.Local.EthHost
		res.Cmd = string(CertShowRet)
		res.Code = protocol.FSuccessCode
		res.Msg = protocol.FSuccessMsg
		i := convert(*list.Data)
		res.Data = *i
		marshal, _ := json.Marshal(res)
		err = ManagerConn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			log.Logger.Errorf("return msg to manager error:%s", err.Error())
			return
		}
	}
}

func CertOwnerHandle() {
	pk, sk, err := HttpRequestOwnerPKSK()
	if err != nil {
		CertOwnerResponseError(err.Error())
		return
	}
	CertOwnerResponseSuccess(pk, sk)
}
