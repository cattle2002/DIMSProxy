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
	"errors"
	"fmt"
	"time"
)

func NeedDouble(request *protocol.HttpCalcRequest) (*protocol.HttpCalcResponse, error) {
	var ptbsc *PTBSC
	var err error
	var sk string
	ca, err := model.FindLastCA(config.Conf.Local.User)
	if err != nil {
		return nil, err
	}
	if ca.ID == 0 {
		return nil, errors.New("not found user keypair")
	}

	//log.Logger.Tracef("平台携带的买家证书时间戳:%d，用户自己的时间戳:%d", request.Payload.BuyerCaTimeStamp, ca.TimeStamp)
	log.Logger.Tracef("plat user ca timestamp:%d:%s,local user ca timestamp:%d:%s", request.Payload.BuyerCaTimeStamp, util.DateFormat(request.Payload.BuyerCaTimeStamp), ca.TimeStamp, util.DateFormat(ca.TimeStamp))

	sk, err = config.GetPrivateKeyPem()
	if err != nil {
		log.Logger.Errorf("get user privatekey error:%s", err.Error())
		return nil, err
	}
	//var err error
	if request.Payload.NeedDecrypt {
		//sellerPk := request.Payload.SellerKey
		sellerPk, _ := GetUserPublicKey(request)

		ptbsc, err = AsymmetricDecryptDoublePlus(sellerPk, sk, request.Payload.CipherSymmetricKey)
		if err != nil {
			log.Logger.Errorf("double  key decrypt error:%s", err.Error())
			return nil, err
		}

		if ptbsc.Buyer == request.Payload.Buyer && ptbsc.Seller == request.Payload.Seller {
			plus, err := file.DownloadPlus(request.Payload.ProductUrl)
			if err != nil {
				return nil, err
			}

			keyDecrypt, err := SymmtricKeyDecryptPlus(string(plus), "aes", ptbsc.Pwd)
			if err != nil {
				errMsg := fmt.Sprintf("aes decrypt product error:%s", err.Error())
				return nil, errors.New(errMsg)
			}
			//decrypt, err := base64.StdEncoding.DecodeString(C.GoString(keyDecrypt))
			decrypt, err := base64.StdEncoding.DecodeString(keyDecrypt)
			if err != nil {
				errMsg := fmt.Sprintf("base64 decode  error:%s", err.Error())
				return nil, errors.New(errMsg)
			}
			reader := bytes.NewReader(decrypt)
			decryptName := fmt.Sprintf("%s.dnc", request.Payload.ProductName)
			err = file.UploadBinary(context.Background(), config.Conf.Minio.ProductUpload, decryptName, reader, int64(len(decrypt)))
			//err = file.UploadBinary(context.Background(), config.Conf.Minio.ProductUpload, "data.enc", reader, int64(len(decrypt)))
			if err != nil {
				//log.Logger.Errorf("上传解密后的数据产品失败:%s", err.Error())
				return nil, fmt.Errorf("upload  decrypt product error:%s", err.Error())
			}
			url, err := file.GetObjectUrl(context.Background(), config.Conf.Minio.ProductUpload, decryptName, time.Second*60*60*24)
			//url, err := file.GetObjectUrl(context.Background(), config.Conf.Minio.ProductUpload, "data.enc", time.Second*60*60*24)
			if err != nil {
				//log.Logger.Errorf("获取解密数据产品的下载链接失败:%s", err.Error())
				return nil, fmt.Errorf("get decrypt product  download url error:%s", err.Error())
			}

			//log.Loggerxp.Infof("确权授权离线数据产品使用:卖家:%s 对称密钥:%s 买家证书时间戳：%d 卖家证书时间戳: %d 卖家公钥：%s", request.Payload.Seller, string("0000000000000000"), request.Payload.BuyerCaTimeStamp, request.Payload.SellerCaTimeStamp, request.Payload.SellerKey)
			log.Logger.Infof("confirm online product:productName:%s Seller:%s SymmtricKey:%s buyerca：%d,buyercaDate:%s, sellerca: %d,sellercaDate:%s, sellerca：%s",
				request.Payload.ProductName, request.Payload.Seller, ptbsc.Pwd, request.Payload.BuyerCaTimeStamp, util.DateFormat(request.Payload.BuyerCaTimeStamp), request.Payload.SellerCaTimeStamp,
				util.DateFormat(int64(request.Payload.SellerCaTimeStamp)), request.Payload.SellerKey)
			err = model.CreateLog(request.Payload.ProductID, request.Payload.ProductName, model.ConfirmGrantOnlineEncryptPro,
				time.Now().UnixMilli(), request.Payload.Seller, ptbsc.Pwd, int64(request.Payload.BuyerCaTimeStamp),
				int64(request.Payload.SellerCaTimeStamp), request.Payload.SellerKey)
			if err != nil {
				log.Logger.Errorf("create product  use log to db error:%s", err.Error())
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
			return nil, errors.New("grantAnd confirm failed")
		}
	} else {
		sellerPk, _ := GetUserPublicKey(request)
		ptbsc, err = AsymmetricDecryptDoublePlus(sellerPk, sk, request.Payload.CipherSymmetricKey)
		if err != nil {
			//log.Logger.Errorf("双密钥解密失败:%s", err.Error())
			return nil, fmt.Errorf("double decrypt  error:%s", err.Error())
		}

		err := model.CreateLog(request.Payload.ProductID, request.Payload.ProductName, model.ConfirmGrantOnlineNotEncryptPro,
			time.Now().UnixMilli(), request.Payload.Seller, ptbsc.Pwd, int64(request.Payload.BuyerCaTimeStamp),
			int64(request.Payload.SellerCaTimeStamp), request.Payload.SellerKey)

		//log.Logger.Trace(err)
		if err != nil {
			log.Logger.Errorf("create product  use log to dberror:%s", err.Error())
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
