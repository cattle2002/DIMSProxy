package file

import (
	"DIMSProxy/config"
	"DIMSProxy/log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var Err error

func NewMinioClient() { //初始化minio连接
	mc, err := minio.New(config.Conf.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Conf.Minio.AccessKeyID, config.Conf.Minio.SecretAccessKey, ""),
		Secure: config.Conf.Minio.UseSSL,
	})
	if err != nil {
		log.Logger.Fatalf("connect minio error:%s", err.Error())
	}
	MinioClient = mc
}
