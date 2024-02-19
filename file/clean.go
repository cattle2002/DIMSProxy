package file

import (
	"DIMSProxy/config"
	"container/list"
	"time"
)

type ObjectInfo struct {
	BucketName string
	ObjectName string
	//UploadTime int64
	CleanTime int64
}

var CleanCh chan *ObjectInfo

// ss := time.Now().Add(time.Hour * 24 * 3).Unix()
// //add := time.Now().After(time.d)
func Sub(begin int64, now int64) bool {
	sub := now - begin
	if sub/int64(24*60*60*config.Conf.Minio.LifeDay) >= 1 {
		return true
	} else {
		return false
	}
}

var CleanList *list.List

func Add() {
	l := list.New()
	CleanCh = make(chan *ObjectInfo, 0)
	CleanList = l
	for {
		v := <-CleanCh
		l.PushBack(v)
	}
}
func Clean() {
	for {
		<-time.After(time.Hour * 6)
		for element := CleanList.Front(); element != nil; element = element.Next() {

			if time.Now().Unix() >= element.Value.(*ObjectInfo).CleanTime {
				go Remove(element.Value.(*ObjectInfo).BucketName, element.Value.(*ObjectInfo).ObjectName)
			}
		}
	}
}
