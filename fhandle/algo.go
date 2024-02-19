package fhandle

import (
	"DIMSProxy/algo"
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func AlgoList() {
	listFunc, err := algo.AlgoGetListFunc()
	if err != nil {
		encryptError := serviceEncryptError(err.Error())
		err := FrontConn.WriteMessage(websocket.TextMessage, encryptError)
		if err != nil {
			log.Logger.Errorf("return forntend algo list error:%s", err.Error())
			return
		}
		return
	}
	var res protocol.AlgoListRes
	res.Cmd = protocol.HttpAlgoListRet
	//res.Cmd = protocol.Algo
	res.Code = protocol.FSuccessCode
	res.Msg = protocol.FSuccessMsg
	res.Data = listFunc
	marshal, _ := json.Marshal(res)
	err = FrontConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("return forntend algo list error:%s", err.Error())
		return
	}
}
