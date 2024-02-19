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

// 确权 卖家私钥加密
func NeedConfirmPermission(request *protocol.HttpCalcRequest) (*protocol.HttpCalcResponse, error) {
	var ptbsc PTBSC

	if request.Payload.NeedDecrypt {
		decrypt, err := gorsa.PublicDecrypt(request.Payload.CipherSymmetricKey, request.Payload.SellerKey)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(decrypt), &ptbsc)
		if err != nil {
			return nil, err
		}
		if ptbsc.Buyer == request.Payload.Buyer && ptbsc.Seller == request.Payload.Seller {
			plus, err := file.DownloadPlus(request.Payload.ProductUrl)
			if err != nil {
				return nil, err
			}

			keyDecrypt, err := SymmtricKeyDecryptPlus(string(plus), "aes", ptbsc.Pwd)
			if err != nil {
				return nil, err
			}
			decrypt, err := base64.StdEncoding.DecodeString(keyDecrypt)
			if err != nil {
				return nil, err
			}
			reader := bytes.NewReader(decrypt)
			decryptName := fmt.Sprintf("%s.dnc", request.Payload.ProductName)
			err = file.UploadBinary(context.Background(), config.Conf.Minio.ProductUpload, decryptName, reader, int64(len(decrypt)))
			//err = file.UploadBinary(context.Background(), config.Conf.Minio.ProductUpload, "data.enc", reader, int64(len(decrypt)))
			if err != nil {
				//return nil, errors.New("上传解密后的数据产品失败")
				return nil, fmt.Errorf("upload decrypt product error:%s", err.Error())
			}
			url, err := file.GetObjectUrl(context.Background(), config.Conf.Minio.ProductUpload, decryptName, time.Second*60*60*24)
			//url, err := file.GetObjectUrl(context.Background(), config.Conf.Minio.ProductUpload, "data.enc", time.Second*60*60*24)
			if err != nil {
				//log.Loggerx.Errorf("获取解密数据产品的下载链接失败:%s", err.Error())
				return nil, fmt.Errorf("get decrypt product error:%s", err.Error())
			}

			log.Logger.Infof("confirm online product:productName:%s seller:%s symmtricKey:%s buyer ca timestamp:%d buyerca date:%s seller ca timestamp: %d sellerca date:%s 卖家公钥：%s",
				request.Payload.ProductName, request.Payload.Seller, ptbsc.Pwd, request.Payload.BuyerCaTimeStamp, util.DateFormat(request.Payload.BuyerCaTimeStamp), request.Payload.SellerCaTimeStamp,
				util.DateFormat(int64(request.Payload.SellerCaTimeStamp)),
				request.Payload.SellerKey)
			//log.Logger.Infof("确权在线数据产品使用:产品名称:%s 卖家:%s 对称密钥:%s 买家证书时间戳：%d 卖家证书时间戳: %d 卖家公钥：%s",
			//	request.Payload.ProductName, request.Payload.Seller, ptbsc.Pwd, request.Payload.BuyerCaTimeStamp, request.Payload.SellerCaTimeStamp,
			//	request.Payload.SellerKey)
			err = model.CreateLog(request.Payload.ProductID, request.Payload.ProductName, model.ConfirmOnlineEncryptPro,
				time.Now().UnixMilli(), request.Payload.Seller, ptbsc.Pwd, int64(request.Payload.BuyerCaTimeStamp),
				int64(request.Payload.SellerCaTimeStamp), request.Payload.SellerKey)
			if err != nil {
				log.Logger.Errorf("create product use log to db  error:%s", err.Error())
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
			return nil, errors.New("confirm failed")
		}
	} else {
		err := model.CreateLog(request.Payload.ProductID, request.Payload.ProductName, model.ConfirmOnlineNotEncryptPro,
			time.Now().UnixMilli(), request.Payload.Seller, ptbsc.Pwd, int64(request.Payload.BuyerCaTimeStamp),
			int64(request.Payload.SellerCaTimeStamp), request.Payload.SellerKey)
		if err != nil {
			log.Logger.Errorf("create product use log to db  error:%s", err.Error())
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
