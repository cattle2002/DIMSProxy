package fhandle

import (
	"DIMSProxy/config"
	"DIMSProxy/file"
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cattle2002/easycrypto/ecrypto"
)

func Encryptx(symmetricKey []byte, filePosition string, tp string, hexKey string) (*protocol.HttpOfflineEncryptRes, error) {
	if tp == "cus" {
		resp, err := http.Get(filePosition)
		if err != nil {
			return nil, fmt.Errorf("get http file error:%s", err.Error())
		}

		all, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Logger.Errorf("read file stream error:%s", err.Error())
			return nil, err
		}
		i, _ := hex.DecodeString(hexKey)

		encrypt, err := ecrypto.AesEncrypt(all, i, []byte("0000000000000000"))
		if err != nil {
			return nil, err
		}
		pathParts := strings.Split(filePosition, "/")
		// 取得最后一个斜杠之后的部分
		lastPart := pathParts[len(pathParts)-1]
		log.Logger.Infof("upload offline product name:%s", lastPart)
		split := strings.Split(lastPart, "%")

		err = file.UploadBinary(context.Background(), config.Conf.Minio.OfflineProductDataEncrypt, split[0], bytes.NewReader(encrypt), int64(len(encrypt)))
		if err != nil {
			//log.Logger.Errorf("上传二进制文件到minio失败:%s", err.Error())
			return nil, fmt.Errorf("upload file to minio error:%s", err.Error())
		}
		url, err := file.GetObjectUrl(context.Background(), config.Conf.Minio.OfflineProductDataEncrypt, split[0], time.Second*60*60*24*7)
		//url, err := ominio.GetObjectUrl(context.Background(), config.Conf.Minio.OfflineProductDataEncrypt, split[0], 0)
		if err != nil {
			//log.Logger.Errorf("获取文件下载链接失败:%s", err.Error())
			return nil, fmt.Errorf("get file download url error:%s", err.Error())
		}
		//log.Logger.Tracef("离线文件下载连接:%s", url)
		var res protocol.HttpOfflineEncryptRes
		res.Code = protocol.FSuccessCode
		res.Cmd = protocol.OffEncryptRet
		res.Data.Url = url
		res.Data.AlgoType = "aes"
		return &res, nil
	}
	if tp == "rand" {
		resp, err := http.Get(filePosition)
		if err != nil {
			return nil, fmt.Errorf("get http file error:%s", err.Error())
		}

		all, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Logger.Errorf("read file stream error:%s", err.Error())
			return nil, err
		}

		decodeString, _ := hex.DecodeString(hexKey)
		encrypt, err := ecrypto.AesEncrypt(all, symmetricKey, decodeString)
		if err != nil {
			return nil, err
		}
		pathParts := strings.Split(filePosition, "/")
		// 取得最后一个斜杠之后的部分
		lastPart := pathParts[len(pathParts)-1]
		log.Logger.Infof("上传的离线数据资源名称:%s", lastPart)
		split := strings.Split(lastPart, "%")

		err = file.UploadBinary(context.Background(), config.Conf.Minio.OfflineProductDataEncrypt, split[0], bytes.NewReader(encrypt), int64(len(encrypt)))
		if err != nil {
			return nil, fmt.Errorf("upload file to minio error:%s", err.Error())
		}
		url, err := file.GetObjectUrl(context.Background(), config.Conf.Minio.OfflineProductDataEncrypt, split[0], time.Second*60*60*24*7)
		//url, err := ominio.GetObjectUrl(context.Background(), config.Conf.Minio.OfflineProductDataEncrypt, split[0], 0)
		if err != nil {
			log.Logger.Errorf("获取文件下载链接失败:%s", err.Error())
			return nil, fmt.Errorf("get file download url error:%s", err.Error())
		}
		//log.Logger.Tracef("离线文件下载连接:%s", url)
		var res protocol.HttpOfflineEncryptRes
		res.Code = protocol.FSuccessCode
		res.Cmd = string(protocol.OffEncryptRet)
		res.Data.Url = url
		res.Data.AlgoType = "aes"

		return &res, nil
	}
	return nil, nil

}
