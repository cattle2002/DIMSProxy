package handle

import (
	"DIMSProxy/log"
	"DIMSProxy/packet"
	"DIMSProxy/protocol"
	"DIMSProxy/util"
	"encoding/json"

	"github.com/gorilla/websocket"
)

func ReduceSync(productID int64, userName string, canUseNumberLocal int64) {

	if WsStatusBool == false {
		log.Logger.Errorf("coreServer is disconnect")
		return
	}
	req := protocol.LocalCanUseSyncReq{
		Cmd:     protocol.LocalCanUse,
		Program: protocol.Monitor,
		Payload: protocol.LocalCanUseSyncReqPayload{
			ID:                util.MsgID(),
			ProductID:         productID,
			UserName:          userName,
			CanUseNumberLocal: canUseNumberLocal,
		},
	}
	marshal, _ := json.Marshal(req)
	bs, err := packet.Packet(marshal)
	if err != nil {
		log.Logger.Errorf("packet msg error:%s", err.Error())
		return
	}
	err = WsConn.WriteMessage(websocket.TextMessage, bs)
	if err != nil {
		log.Logger.Errorf("send reduce local count:%s", err.Error())
		return
	}
}
