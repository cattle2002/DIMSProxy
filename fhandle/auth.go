package fhandle

import (
	"DIMSProxy/config"
	"DIMSProxy/file"
	"DIMSProxy/handle"
	"DIMSProxy/log"
	"DIMSProxy/model"
	"DIMSProxy/pfilter"
	"DIMSProxy/protocol"
	"DIMSProxy/util"
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/cattle2002/easycrypto/ecrypto"
)

type PTBSC struct {
	Pwd        string `json:"Pwd"`
	TimeStamp  int64  `json:"TimeStamp"`
	Buyer      string `json:"Buyer"`
	Seller     string `json:"Seller"`
	ContentMD5 string `json:"ContentMD5"`
}

/*
用报文携带的公钥，解密用自己私钥解密，确权 授权 验证
*/
func VerifyOffOrOn(request *protocol.HttpCalcRequest) int {
	/* 确权 从请求报文里面取出如果NeedConfirmPermission==True代表这是需要确权的数据产品,在实际的处理流程当中,需要使用自己的数据拥有者的公钥进行解密
	 */
	if request.Payload.NeedConfirmPermission {
		return 1
	}
	/* 授权 从请求报文里面取出如果NeedGrantPermission==True代表这是需要授权的数据产品，在实际的处理流程当中,需要使用自己的私钥进行解密
	 */
	if request.Payload.NeedGrantPermission {
		return 2
	}
	/* 确权授权 从请求报文里面取出如果 NeedConfirmAndGrantPermission==true代表该数据产品既需要授权也需要确权,在处理过程中，使用数据拥有者的公钥进行解密,然后
	使用数据购买者的私钥进行解密*/
	if request.Payload.NeedConfirmAndGrantPermission {
		return 3
	}
	return 0
}

func CalcPlusPermission(request *protocol.HttpCalcRequest) (*protocol.HttpCalcResponse, error) {
	pk, err := pfilter.FilterInsertPK(request.Payload.Seller, int64(request.Payload.SellerCaTimeStamp))
	if err != nil {
		//log.Logger.Errorf("查找用户公钥失败:%s", err.Error())
		return nil, fmt.Errorf("find user public key error:%s", err.Error())

	} else {
		if !pk {
			err := model.CreateCert(request.Payload.Seller, int64(request.Payload.SellerCaTimeStamp), request.Payload.SellerKey, "")
			if err != nil {
				return nil, fmt.Errorf("create user public to  db error:%s", err.Error())
			}
		}
	}
	if request.Payload.ProductType == Off {

		var res protocol.HttpCalcResponse
		plus, err := file.DownloadPlus(request.Payload.ProductUrl)
		if err != nil {
			return nil, err
		}

		decrypt, err := ecrypto.AesDecrypt(plus, []byte("0000000000000000"), []byte("0000000000000000"))
		if err != nil {
			return nil, err
		}
		err = file.UploadBinary(context.Background(), config.Conf.Minio.ProductUpload, request.Payload.ProductName, bytes.NewReader(decrypt), int64(len(decrypt)))
		if err != nil {
			return nil, err
		}
		url, err := file.GetObjectUrl(context.Background(), config.Conf.Minio.ProductUpload, request.Payload.ProductName, time.Second*60*60*24*7)
		if err != nil {
			return nil, err
		}
		//log.Logger.Infof("授权离线数据产品使用:卖家:%s 对称密钥:%s 买家证书时间戳：%d 卖家证书时间戳: %d 卖家公钥：%s", request.Payload.Seller, "0000000000000000", request.Payload.BuyerCaTimeStamp, request.Payload.SellerCaTimeStamp, request.Payload.SellerKey)
		log.Logger.Infof("grant offline product use,seller:%s symmetric key:%s buyer ca timeStamp:%d,buyer ca date:%s"+
			" seller ca timestamp: %d,seller ca date:%s,seller publickey:%s", request.Payload.Seller, "0000000000000000", request.Payload.BuyerCaTimeStamp, util.DateFormat(request.Payload.BuyerCaTimeStamp), request.Payload.SellerCaTimeStamp, util.DateFormat(int64(request.Payload.SellerCaTimeStamp)), request.Payload.SellerKey)

		err = model.CreateLog(request.Payload.ProductID, request.Payload.ProductName, model.GrantOfflinePro, time.Now().UnixMilli(),
			request.Payload.Seller, hex.EncodeToString([]byte("0000000000000000")), int64(request.Payload.BuyerCaTimeStamp), int64(request.Payload.SellerCaTimeStamp),
			request.Payload.SellerKey)
		if err != nil {
			log.Logger.Errorf("write  product use log to db  error:%s", err.Error())
		}
		//todo 减少本地可使用次数
		one, err := model.ReduceOne(request.Payload.ProductID, request.Payload.ProductName)
		if err != nil {
			log.Logger.Errorf("reduce db product use number error:%s", err.Error())
			return nil, err
		}
		go handle.ReduceSync(request.Payload.ProductID, request.Payload.ProductName, one)
		res.Cmd = protocol.CalcRet
		res.RetCode = protocol.FSuccessCode
		res.ErrorMsg = protocol.FSuccessMsg
		res.Payload.ID = request.Payload.ID
		res.Payload.HaveData = false
		res.Payload.Url = url
		return &res, nil
	} else {
		fl := VerifyOffOrOn(request)
		if fl == 1 {
			permission, err := NeedConfirmPermission(request)
			if err != nil {

				return nil, err
			} else {
				//todo 减少本地可使用次数
				one, err := model.ReduceOne(request.Payload.ProductID, request.Payload.ProductName)
				if err != nil {
					log.Logger.Errorf("reduce db product use number error:%s", err.Error())
					return nil, err
				}

				go handle.ReduceSync(request.Payload.ProductID, request.Payload.ProductName, one)
				return permission, nil
			}
		}
		if fl == 2 {
			grant, err := NeedGrant(request)
			if err != nil {
				return nil, err
			} else {
				//todo 减少本地可使用次数
				one, err := model.ReduceOne(request.Payload.ProductID, request.Payload.ProductName)
				if err != nil {
					log.Logger.Errorf("reduce db product use number error :%s", err.Error())
					return nil, err
				}
				go handle.ReduceSync(request.Payload.ProductID, request.Payload.ProductName, one)
				return grant, nil
			}
		}
		if fl == 3 {
			double, err := NeedDouble(request)
			if err != nil {
				return nil, err
			} else {
				//todo 减少本地可使用次数
				one, err := model.ReduceOne(request.Payload.ProductID, request.Payload.ProductName)
				if err != nil {
					log.Logger.Errorf("reduce db product use number error:%s", err.Error())
					return nil, err
				}
				go handle.ReduceSync(request.Payload.ProductID, request.Payload.ProductName, one)
				return double, nil
			}
		}
		return nil, errors.New("internal error")
	}
}
