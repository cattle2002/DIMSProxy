package handle

import (
	"DIMSProxy/file"
	"DIMSProxy/log"
	"DIMSProxy/packet"
	"DIMSProxy/protocol"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var RetCode int

func Reader() {
	for {
		if WsStatusBool == false {
			continue
		}
		_, p, err := WsConn.ReadMessage()
		if err != nil {
			WsStatusBool = false
			WsStatusCh <- false
			log.Logger.Errorf("bus read msg error:%s", err.Error())
			continue
		}
		uPacket, err := packet.UPacket(p)
		if err != nil {
			log.Logger.Errorf("unpacket coreserver msg error:%s", err.Error())
			continue
		}
		log.Logger.Infof("read coreserver msg:%s", string(uPacket))

		//todo cmd 区分
		cmd, err := packet.ExtractCmdValue(string(uPacket))
		if err != nil {
			log.Logger.Errorf("extract coreserver msg cmd error:%s", err.Error())
			continue
		}
		if cmd == protocol.LimitRet {
			PInfoCh <- uPacket
		}
		if cmd == protocol.LoginRet {
			//protocol.LoginReq{}
			var loginRes protocol.LoginRes
			err := json.Unmarshal(uPacket, &loginRes)
			if err != nil {
				log.Logger.Errorf("Json Unmarshal Error:%s", err.Error())
			}
			//log.Logger.Tracef("monitor login return：%v", loginRes)
			RetCode = loginRes.RetCode
		}
		if cmd == "MonitorData" {
			//todo
			var req protocol.MonitorProductReq
			err = json.Unmarshal(uPacket, &req)
			if err != nil {
				continue
			}
			f := VerifyUserName(req.Payload.Buyer)
			if f == false {
				err := ErrorResponse("用户名不匹配", req.Payload.ID, req.Payload.Buyer, req.Payload.ProductID, req.Payload.Demand, "")
				if err != nil {
					log.Logger.Errorf("return monitor error msg error:%s", err.Error())
					continue
				}
				continue
			} else {
				seller, logortimes, err := Monitor(req.Payload.ProductID, req.Payload.Demand)
				//log.Logger.Errorf("e:%s", err.Error())
				if err != nil {
					err = ErrorResponse(err.Error(), req.Payload.ID, seller, req.Payload.ProductID, req.Payload.Demand, "")
					if err != nil {
						log.Logger.Errorf("发送消息到核心服务器错误:%s", err.Error())
						continue
					}
					continue
				} else {
					//if req.Payload.Demand == protocol.Time {
					//times, err := strconv.Atoi(logortimes)
					//if err != nil {
					//	log.Logger.Errorf("内部错误")
					//	continue
					//}
					if len(logortimes) >= 1000 {
						log.Logger.Tracef("日志文件大小:%d", len(logortimes))
						reader := bytes.NewReader([]byte(logortimes))
						name := fmt.Sprintf("%s:%s", seller, strconv.Itoa(int(time.Now().UnixMilli())))
						err := file.UploadBinary(context.Background(), "message", name, reader, int64(len(logortimes)))
						if err != nil {
							err = ErrorResponse(err.Error(), req.Payload.ID, seller, req.Payload.ProductID, req.Payload.Demand, "")
							if err != nil {
								log.Logger.Errorf("发送消息到核心服务器错误:%s", err.Error())
								continue
							}
							continue
						} else {
							url, err := file.GetObjectUrl(context.Background(), "message", name, time.Hour*2)
							if err != nil {
								err = ErrorResponse(err.Error(), req.Payload.ID, seller, req.Payload.ProductID, req.Payload.Demand, "")
								if err != nil {
									log.Logger.Errorf("发送消息到核心服务器错误:%s", err.Error())
									continue
								}
								continue
							} else {
								log.Logger.Tracef("URL:%s", url)
								err = SuccessResponse(req.Payload.ID, seller, req.Payload.ProductID, req.Payload.Demand, url)
								if err != nil {
									log.Logger.Errorf("发送消息到核心服务器错误:%s", err.Error())
									continue
								}
							}
						}
						continue
					} else {
						err = SuccessResponse(req.Payload.ID, seller, req.Payload.ProductID, req.Payload.Demand, logortimes)
						if err != nil {
							log.Logger.Errorf("发送消息到核心服务器错误:%s", err.Error())
							continue
						}
					}

				}
			}
		}
	}
}

var Chc chan bool

func WsSend() {
	for {
		<-Chc
		WsConn.WriteMessage(1, []byte("hello,world"))
	}
}
