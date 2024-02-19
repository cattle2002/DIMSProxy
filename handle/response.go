package handle

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"DIMSProxy/packet"
	"DIMSProxy/protocol"
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func ErrorResponse(errmsg string, id int64, seller string, productID int64, demand string, reply string) error {
	var res protocol.MonitorProductRes
	res.Command = protocol.MonitorDataRet
	res.ErrorMsg = errmsg
	res.RetCode = protocol.CoreErrorCode
	res.Program = protocol.Monitor
	res.Payload.ID = id
	res.Payload.Buyer = config.Conf.Local.User
	res.Payload.Seller = seller
	res.Payload.Demand = demand
	res.Payload.ProductID = productID
	res.Payload.Log = reply
	marshal, err := json.Marshal(res)
	if err != nil {
		return err
	}
	log.Logger.Tracef("ErrorReponse msg:%s", string(marshal))
	es, err := packet.Packet(marshal)
	if err != nil {
		return err
	}
	err = WsConn.WriteMessage(websocket.TextMessage, es)
	if err != nil {
		return err
	}
	return nil
}
func SuccessResponse(id int64, seller string, productID int64, demand string, reply string) error {

	var res protocol.MonitorProductRes
	res.Command = protocol.MonitorDataRet
	res.RetCode = protocol.CoreSuccessCode
	res.Payload.ID = id
	res.Program = protocol.Monitor
	res.Payload.Buyer = config.Conf.Local.User
	res.Payload.Seller = seller
	res.Payload.Demand = demand
	res.Payload.ProductID = productID
	res.Payload.Log = reply
	marshal, err := json.Marshal(res)
	if err != nil {
		return err
	}
	log.Logger.Tracef("SuccessReponse msg:%s", string(marshal))
	htmlJson := TransHtmlJson(marshal)
	es, err := packet.Packet(htmlJson)
	if err != nil {
		return err
	}
	err = WsConn.WriteMessage(websocket.TextMessage, es)
	if err != nil {
		return err
	}
	log.Logger.Tracef("return log json2:%s", es)
	return nil
}
func TransHtmlJson(data []byte) []byte {
	data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)
	data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
	data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
	return data
}
