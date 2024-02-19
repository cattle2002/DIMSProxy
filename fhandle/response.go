package fhandle

import (
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"bytes"
	"encoding/json"

	"github.com/gorilla/websocket"
)

func calcResponse(retCode int, msg string, id int64, haveData bool, data string, url string) {
	var res protocol.HttpCalcResponse
	res.Cmd = protocol.CalcRet
	res.RetCode = retCode
	res.ErrorMsg = msg
	res.Payload.ID = id
	res.Payload.HaveData = haveData
	res.Payload.Data = []byte(data)
	res.Payload.Url = url
	marshal, _ := json.Marshal(res)
	err := FrontConn.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Logger.Errorf("retrun  front calc  error:%s", err.Error())
		return
	}
	log.Logger.Tracef("retrun  front calc  error:%s", string(marshal))

}
func ServiceCalcError(errmsg string) []byte {
	response := protocol.HttpCalcResponse{
		Cmd:      protocol.CalcRet,
		RetCode:  protocol.FErrorCode,
		ErrorMsg: errmsg,
		Payload:  protocol.HttpCalcResponsePayload{},
	}
	marshal, _ := json.Marshal(response)
	return marshal
}
func ServiceCalcSuccess(response *protocol.HttpCalcResponse) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(response)
	if err != nil {
		return nil, err
	}
	marshal := buf.Bytes()
	return marshal, nil
}
