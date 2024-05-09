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
func AlgoRegister() {
	a := "{\"Name\":\"算法名称\",\"Type\":\"EXE\",\"FilePath\":\"D:\\\\workdir\\\\DIMSProxy\\\\algo\\\\algo3.exe\",\"StartupCmd\":\"\",\"ExeInputEOF\":\"\",\"MaxExecTime\":20,\"InputExample\":\"\",\"AlgoFuncName\":\"\",\"Parameters\":\"\",\"CreatedAt\":\"0001-01-01T00:00:00Z\"}"
	registerFunc, err := algo.AlgoRegisterFunc(a)
	if err != nil {
		log.Logger.Errorf("return forntend algo list error:%s", err.Error())
		return
	}
	err = FrontConn.WriteMessage(websocket.TextMessage, []byte(registerFunc))
	if err != nil {
		log.Logger.Errorf("return forntend algo list error:%s", err.Error())
		return
	}
}
