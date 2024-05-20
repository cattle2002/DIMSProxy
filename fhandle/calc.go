package fhandle

import (
	"DIMSProxy/handle"
	"DIMSProxy/log"
	"DIMSProxy/model"
	"DIMSProxy/pfilter"
	"DIMSProxy/protocol"
	"DIMSProxy/util"
	"github.com/gorilla/websocket"
)

func Calc(calcReq protocol.HttpCalcRequest) {
	filter, err := pfilter.FilterStatus(calcReq.Payload.ProductID, calcReq.Payload.ProductName)
	if err != nil {
		if err.Error() == protocol.ProductNotFound {
			//todo 数据库没发现数据产品的信息,开协程去获取信息
			err = handle.ProductInfoRequest(calcReq.Payload.ProductID, calcReq.Payload.Buyer, calcReq.Payload.ProductName)
			if err != nil {
				log.Logger.Errorf("get product  limit error:%s", err.Error())
				calcResponse(protocol.FErrorCode, err.Error(), calcReq.Payload.ID, false, "", "")
				return
			} else {
				lc, lt, err := handle.PInfo()
				if err != nil {
					log.Logger.Errorf("get product limit information  error:%s", err.Error())
					calcResponse(protocol.FErrorCode, err.Error(), calcReq.Payload.ID, false, "", "")
					return
				} else {
					if util.GetCurrentMillSecond(lt) {
						err = model.Create(calcReq.Payload.ProductID, calcReq.Payload.ProductName, calcReq.Payload.Seller, true, false, lc, lt)
						if err != nil {
							log.Logger.Errorf("write product limit info to db error:%s", err.Error())
							calcResponse(protocol.FErrorCode, err.Error(), calcReq.Payload.ID, false, "", "")
							return
						} else {
							//todo 进行calc
							calc, err := CalcPlusPermission(&calcReq)
							if err != nil {
								calcError := ServiceCalcError(err.Error())
								err = FrontConn.WriteMessage(websocket.TextMessage, calcError)
								if err != nil {
									log.Logger.Errorf("calc error:%s", err.Error())
									return
								}
								return
							} else {
								success, err := ServiceCalcSuccess(calc)
								if err != nil {
									log.Logger.Errorf("calc error:%s", err.Error())
									return
								}
								err = FrontConn.WriteMessage(websocket.TextMessage, success)
								if err != nil {
									log.Logger.Errorf("calc error:%s", err.Error())
									return
								}
								return
							}
						}
					} else {
						err = model.Create(calcReq.Payload.ProductID, calcReq.Payload.ProductName, calcReq.Payload.Seller, false, true, lc, lt)
						if err != nil {
							calcResponse(protocol.FErrorCode, err.Error(), calcReq.Payload.ID, false, "", "")
							return
						} else {
							//todo 进行calc
							calc, err := CalcPlusPermission(&calcReq)
							if err != nil {
								calcError := ServiceCalcError(err.Error())
								err = FrontConn.WriteMessage(websocket.TextMessage, calcError)
								if err != nil {
									log.Logger.Errorf("calc error:%s", err.Error())
									return
								}
								return
							} else {
								success, err := ServiceCalcSuccess(calc)
								if err != nil {
									log.Logger.Errorf("calc error:%s", err.Error())
									return
								}
								err = FrontConn.WriteMessage(websocket.TextMessage, success)
								if err != nil {
									log.Logger.Errorf("calc error:%s", err.Error())
									return
								}
								return
							}
						}
					}
				}
			}
		} else {
			calcResponse(protocol.FErrorCode, err.Error(), calcReq.Payload.ID, false, "", "")
			return
		}
	} else {
		if filter {
			_, err = pfilter.FilterConditions(calcReq.Payload.ProductID, calcReq.Payload.ProductName)
			if err != nil {
				calcError := ServiceCalcError(err.Error())
				err = FrontConn.WriteMessage(websocket.TextMessage, calcError)
				if err != nil {
					log.Logger.Errorf("send msg front error:%s", err.Error())
					return
				}
				return
			} else {
				//todo calc
				calc, err := CalcPlusPermission(&calcReq)
				if err != nil {
					calcError := ServiceCalcError(err.Error())
					err = FrontConn.WriteMessage(websocket.TextMessage, calcError)
					if err != nil {
						log.Logger.Errorf("send msg front error:%s", err.Error())
						return
					}
					return
				} else {
					success, err := ServiceCalcSuccess(calc)
					if err != nil {
						log.Logger.Errorf("handle calc  error:%s", err.Error())
						return
					}
					err = FrontConn.WriteMessage(websocket.TextMessage, success)
					if err != nil {
						log.Logger.Errorf("handle calc  error:%s", err.Error())
						return
					}
					return
				}
			}
		} else {
			calcError := ServiceCalcError("该数据产品被禁止使用 ")
			err = FrontConn.WriteMessage(websocket.TextMessage, calcError)
			if err != nil {
				log.Logger.Errorf("handle calc  error:%s", err.Error())
				return
			}
			return
		}
	}
}
