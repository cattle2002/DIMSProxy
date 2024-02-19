package model

import (
	"DIMSProxy/log"
	"DIMSProxy/util"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
)

var CertCh chan bool

func OpenCert() {
	certFile := util.ReturnLFileAbs("cert.db")
	_, err = os.Stat(certFile)
	if err != nil {
		go func() {
			<-CertCh
			DBCert, err = gorm.Open(sqlite.Open(certFile), &gorm.Config{})
			if err != nil {
				log.Logger.Fatalf("open cert.db error:%s", err.Error())
			}
			log.Logger.Info("open cert.db success")
		}()
	} else {
		DBCert, err = gorm.Open(sqlite.Open(certFile), &gorm.Config{})
		if err != nil {
			log.Logger.Fatalf("open cert.db error:%s", err.Error())
		}
		log.Logger.Info("open cert.db success")
	}
}
func CreateCert(user string, TimeStamp int64, PublicKey string, PrivateKey string) error {
	ct := CertModel{
		User:       user,
		TimeStamp:  TimeStamp,
		PublicKey:  PublicKey,
		PrivateKey: PrivateKey,
	}
	tx := DBCert.Create(&ct)
	return tx.Error
}
func FindPK(user string, timeStamp int64) (bool, error) {
	var ct CertModel
	tx := DBCert.Where("user = ? and timeStamp = ?", user, timeStamp).Find(&ct)
	if tx.Error != nil {
		return false, tx.Error
	} else {
		if ct.ID == 0 {
			return false, nil
		} else {
			return true, nil
		}
	}
}
func FindLastCA(user string) (*CertModel, error) {
	var certs []CertModel
	var crt CertModel
	tx := DBCert.Where("user = ?", user).Order("timeStamp DESC").Find(&certs)

	if len(certs) == 0 {
		return &crt, nil
	}
	return &certs[0], tx.Error
}
