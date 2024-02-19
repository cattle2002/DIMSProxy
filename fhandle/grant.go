package fhandle

import (
	"DIMSProxy/config"
	"DIMSProxy/file"
	"DIMSProxy/log"
	"DIMSProxy/model"
	"DIMSProxy/protocol"
	"DIMSProxy/util"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/wenzhenxi/gorsa"
)

func NeedGrant(request *protocol.HttpCalcRequest) (*protocol.HttpCalcResponse, error) {
	var ptbsc PTBSC

	if request.Payload.NeedDecrypt {
		//todo 获取自己的私钥
		skPem, err := config.GetPrivateKeyPem()
		if err != nil {
			return nil, err
		}
		decrypt, err := gorsa.PriKeyDecrypt(request.Payload.CipherSymmetricKey, skPem)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(decrypt), &ptbsc)
		if err != nil {
			return nil, err
		}
		if request.Payload.Seller == ptbsc.Seller && request.Payload.Buyer == ptbsc.Buyer {
			plus, err := file.DownloadPlus(request.Payload.ProductUrl)
			if err != nil {
				return nil, err
			}

			keyDecrypt, err := SymmtricKeyDecryptPlus(string(plus), "aes", ptbsc.Pwd)
			if err != nil {
				panic(err)
			}
			//decrypt, err := base64.StdEncoding.DecodeString(C.GoString(keyDecrypt))
			decrypt, err := base64.StdEncoding.DecodeString(keyDecrypt)
			if err != nil {
				return nil, fmt.Errorf("base64Decode error:%s", err.Error())
			}
			reader := bytes.NewReader([]byte(decrypt))
			err = file.UploadBinary(context.Background(), config.Conf.Minio.ProductUpload, "data.enc", reader, int64(len(decrypt)))
			if err != nil {
				return nil, errors.New("上传解密后的数据产品失败")
			}
			url, err := file.GetObjectUrl(context.Background(), config.Conf.Minio.ProductUpload, "data.enc", time.Second*60*60*24)
			if err != nil {
				return nil, fmt.Errorf("get decrypt product download  url error:%s", err.Error())
			}

			//log.Loggerxp.Infof("授权离线数据产品使用:卖家:%s 对称密钥:%s 买家证书时间戳：%d 卖家证书时间戳: %d 卖家公钥：%s", request.Payload.Seller, string("0000000000000000"), request.Payload.BuyerCaTimeStamp, request.Payload.SellerCaTimeStamp, request.Payload.SellerKey)
			log.Logger.Infof("grant online product:productName:%s seller:%s symmetricKey:%s buyerca：%d,buyercaDate:%s sellerca: %d,sellercaDate:%s sellerca：%s",
				request.Payload.ProductName, request.Payload.Seller, ptbsc.Pwd, request.Payload.BuyerCaTimeStamp, util.DateFormat(request.Payload.BuyerCaTimeStamp), request.Payload.SellerCaTimeStamp,
				util.DateFormat(int64(request.Payload.SellerCaTimeStamp)), request.Payload.SellerKey)
			err = model.CreateLog(request.Payload.ProductID, request.Payload.ProductName, model.GrantOnlineEncryptPro, time.Now().UnixMilli(), request.Payload.Seller, ptbsc.Pwd,
				request.Payload.BuyerCaTimeStamp, int64(request.Payload.SellerCaTimeStamp), request.Payload.SellerKey)
			if err != nil {
				log.Logger.Errorf("create productd use log to db error:%s", err.Error())
			}
			return &protocol.HttpCalcResponse{
				Cmd:      protocol.CalcRet,
				RetCode:  protocol.FSuccessCode,
				ErrorMsg: "",
				Payload: protocol.HttpCalcResponsePayload{
					ID:       request.Payload.ID,
					HaveData: false,
					Data:     nil,
					Url:      url,
				},
			}, nil
		} else {
			return nil, errors.New("grantAnd Confirm failed")
		}
	} else {
		err := model.CreateLog(request.Payload.ProductID, request.Payload.ProductName, model.GrantOnlineNotEncryptPro, time.Now().UnixMilli(), request.Payload.Seller, ptbsc.Pwd,
			int64(request.Payload.BuyerCaTimeStamp), int64(request.Payload.SellerCaTimeStamp), request.Payload.SellerKey)
		if err != nil {
			log.Logger.Errorf("create productd use log to db error:%s", err.Error())
		}
		return &protocol.HttpCalcResponse{
			Cmd:      protocol.CalcRet,
			RetCode:  protocol.FSuccessCode,
			ErrorMsg: "",
			Payload: protocol.HttpCalcResponsePayload{
				ID:       request.Payload.ID,
				HaveData: false,
				Data:     nil,
				Url:      request.Payload.ProductUrl,
			},
		}, nil
	}
}
