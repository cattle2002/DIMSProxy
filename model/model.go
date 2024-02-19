package model

import "gorm.io/gorm"

var DBInfo *gorm.DB
var DBLog *gorm.DB
var DBCert *gorm.DB

var err error

type CertModel struct {
	gorm.Model
	User       string `gorm:"column:user;not null"`
	TimeStamp  int64  `gorm:"column:timeStamp;not null"`
	PublicKey  string `gorm:"column:publicKey;not null"`
	PrivateKey string `gorm:"column:privateKey"`
}

func (t *CertModel) TableName() string {
	return "cert"
}

const (
	ConfirmOfflinePro               = "ConfirmOfflinePro"
	GrantOfflinePro                 = "GrantOfflinePro"
	ConfirmOnlineNotEncryptPro      = "ConfirmOnlineNotEncryptPro"
	ConfirmOnlineEncryptPro         = "ConfirmOnlineEncryptPro"
	GrantOnlineEncryptPro           = "GrantOnlineEncryptPro"
	GrantOnlineNotEncryptPro        = "GrantOnlineNotEncryptPro"
	ConfirmGrantOnlineEncryptPro    = "ConfirmGrantOnlineEncryptPro"
	ConfirmGrantOnlineNotEncryptPro = "ConfirmGrantOnlineNotEncryptPro"
)
