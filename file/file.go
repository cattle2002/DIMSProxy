package file

import (
	"DIMSProxy/config"
	"DIMSProxy/log"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

func DownloadBinary(ctx context.Context, bucketName string, objectName string) ([]byte, error) {
	object, err := MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}
	stat, err := object.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size == int64(len(bs)) {
		return bs, nil
	} else {
		return nil, errors.New("download file error")
	}
}

func UploadBinary(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64) error {

	_, err := MinioClient.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{})
	if err != nil {
		log.Logger.Errorf("Upload Error:%s", err.Error())
		return err
	}
	var info ObjectInfo
	info.BucketName = bucketName
	info.ObjectName = objectName
	expire := config.Conf.Minio.LifeDay * 24 * 60 * 60
	info.CleanTime = time.Now().Add(time.Duration(expire) * time.Second).Unix()
	CleanCh <- &info
	return err
}

func GetObjectUrl(ctx context.Context, bucketName string, objectName string, expires time.Duration) (string, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment")
	urls, err := MinioClient.PresignedGetObject(ctx, bucketName, objectName, expires, reqParams)
	if err != nil {
		return "", err
	}
	return urls.String(), err
}

func DownloadPlus(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func Remove(buckName string, objectName string) {
	err := MinioClient.RemoveObject(context.Background(), buckName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Logger.Errorf("clean expired product faile:%s", err.Error())
		return
	}
}
