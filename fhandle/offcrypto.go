package fhandle

import (
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/gorilla/websocket"
)

func serviceEncryptError(errmsg string) []byte {
	res := protocol.HttpOfflineEncryptRes{
		Cmd:  protocol.OffEncryptRet,
		Code: protocol.FErrorCode,
		Msg:  errmsg,
		Data: protocol.HttpOfflineEncryptResData{},
	}
	marshal, _ := json.Marshal(res)
	return marshal
}
func Encrypt(req protocol.HttpOfflineEncryptReq) {

	if req.Password == "" && req.Encryption == "" {
		encryptError := serviceEncryptError("the symmetricKey not exist")
		err := FrontConn.WriteMessage(websocket.TextMessage, encryptError)
		if err != nil {
			log.Logger.Errorf("write front offline encrypt error:%s", err.Error())
			return
		}
		return
	}
	if req.Encryption == "random" {
		toString := hex.EncodeToString([]byte("0000000000000000"))
		log.Logger.Infof("offline product rand key:%v", toString)

		encrypt, err := Encryptx([]byte("0000000000000000"), req.FileUrl, "rand", toString)
		if err != nil {
			log.Logger.Errorf("rand key encrypt error:%s", err.Error())
			encryptError := serviceEncryptError(err.Error())
			err := FrontConn.WriteMessage(websocket.TextMessage, encryptError)
			if err != nil {
				log.Logger.Errorf("write front offline encrypt error:%s", err.Error())
				return
			}
			return
		}
		log.Logger.Infof("offline encrypt msg:%v", encrypt)
		marshal, err := json.Marshal(encrypt)
		if err != nil {
			log.Logger.Errorf("json marshal error:%s", err.Error())
			return
		}
		//success := serviceEncryptSuccess(encrypt)
		err = FrontConn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			log.Logger.Errorf("write front offline encrypt error:%s", err.Error())
			return
		}
		return
	}
	if req.Encryption == "customize" {
		log.Logger.Infof("offline product cus key:%v", req.Password)

		encrypt, err := Encryptx(nil, req.FileUrl, "cus", req.Password)
		if err != nil {
			log.Logger.Errorf("customize key encrypt error:%s", err.Error())
			encryptError := serviceEncryptError(err.Error())
			err = FrontConn.WriteMessage(websocket.TextMessage, encryptError)
			if err != nil {
				log.Logger.Errorf("write front offline encrypt error:%s", err.Error())
				return
			}
			return
		}
		log.Logger.Infof("offline encrypt msg:%v", encrypt)
		//success := serviceEncryptSuccess(encrypt)
		marshal, _ := json.Marshal(encrypt)
		marshal = []byte(strings.ReplaceAll(string(marshal), "\\u0026", "&"))
		//success := serviceEncryptSuccess(encrypt)
		err = FrontConn.WriteMessage(websocket.TextMessage, marshal)
		if err != nil {
			log.Logger.Errorf("write front offline encrypt error:%s", err.Error())
			return
		}
	}

}
