package handle

import (
	"DIMSProxy/log"
	"DIMSProxy/model"
	"DIMSProxy/protocol"
	"errors"
	"strconv"
)

func Monitor(productID int64, tp string) (string, string, error) {
	if tp == protocol.Stop {
		seller, err := model.UpdateStatus(productID, "", protocol.Stop)
		if err != nil {
			return "", "", err
		}
		//todo 返回消息给核心服务器
		return seller, "", nil
	}
	if tp == protocol.Renew {
		seller, err := model.UpdateStatus(productID, "", protocol.Renew)
		if err != nil {
			return "", "", err
		}
		//todo 返回消息给核心服务器
		return seller, "", nil
	}
	if tp == protocol.Logg {
		seller, s, err := model.FindBatchLog(productID)
		if err != nil {
			log.Logger.Trace("cao")
			return "", "", err
		}
		//log.Logger.Tracef("Logg-------------------------------------Logg:%s", s)
		//todo 返回消息给核心服务器
		return seller, s, nil
	}
	if tp == protocol.Time {
		seller, times, err := model.GetProductTimes(productID)
		if err != nil {
			return "", "", err
		}
		itoa := strconv.Itoa(int(times))
		return seller, itoa, nil
		//todo 返回消息给核心服务器

	}
	return "", "", errors.New("内部错误")
}
