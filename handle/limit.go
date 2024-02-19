package handle

import (
	"DIMSProxy/packet"
	"DIMSProxy/protocol"
	"DIMSProxy/util"
	"encoding/json"
	"errors"
	"time"

	"github.com/gorilla/websocket"
)

// todo 获取数据产品的限制信息
func ProductInfoRequest(productID int64, buyer string, productName string) error {
	req := protocol.LimitReq{
		Cmd:     protocol.Limit,
		Program: protocol.Monitor,
		Payload: protocol.LimitReqPayload{
			ID:          util.MsgID(),
			Buyer:       buyer,
			ProductName: productName,
			ProductID:   productID,
		},
	}
	marshal, _ := json.Marshal(req)
	packets, err := packet.Packet(marshal)
	if err != nil {
		return err
	}
	err = WsConn.WriteMessage(websocket.TextMessage, packets)
	if err != nil {
		return err
	}
	return nil
}

func PInfo() (int64, int64, error) {
	select {
	case v := <-PInfoCh:
		var res protocol.LimitRes
		err := json.Unmarshal(v, &res)
		if err != nil {
			return 0, 0, err
		}
		if res.RetCode != protocol.CoreSuccessCode {
			return 0, 0, errors.New(res.ErrorMsg)
		}
		return res.Payload.Detail.CanUseNumberLocal, res.Payload.Detail.CanUseTimeLocal, nil
	case <-time.After(time.Second * 30):
		return 0, 0, errors.New("获取产品信息超时")
	}
}
