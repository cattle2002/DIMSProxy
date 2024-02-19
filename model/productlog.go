package model

import (
	"DIMSProxy/log"
	"encoding/json"
	"errors"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type ProductLogModel struct {
	ProductID         int64  `gorm:"column:product_id;not null"`
	ProductName       string `gorm:"column:product_name;not null"`
	ProductType       string `gorm:"column:product_type;not null"` //产品类型
	CurrentTime       int64  `gorm:"column:current_time;not null"`
	Seller            string `gorm:"column:seller;not null"`
	SymmetricKey      string `gorm:"column:symmetrickey;not null"` //数据产品对称密钥
	BuyerCaTimestamp  int64  `gorm:"column:buyerca_timestamp;not  null"`
	SellerCaTimeStamp int64  `gorm:"column:sellerca_timestamp"`
	SellerPublicKey   string `gorm:"column:seller_publickey"`
}

type ProductLogModelDay struct {
	ProductID         int64
	ProductName       string
	ProductType       string
	CurrentTime       string
	Seller            string
	SymmetricKey      string
	BuyerCaTimestamp  int64
	SellerCaTimeStamp int64
	SellerPublicKey   string
}

func DateFormat(timestamp int) string {
	tm := time.UnixMilli(int64(timestamp))
	return tm.Format("2006-01-02 15:04:05")
}
func convert(fb []ProductLogModel) []ProductLogModelDay {
	var pbe []ProductLogModelDay
	for i := range fb {
		var pbee ProductLogModelDay
		pbee.ProductID = fb[i].ProductID
		pbee.ProductName = fb[i].ProductName
		pbee.Seller = fb[i].Seller
		pbee.SellerPublicKey = fb[i].SellerPublicKey
		pbee.SellerCaTimeStamp = fb[i].SellerCaTimeStamp
		pbee.BuyerCaTimestamp = fb[i].BuyerCaTimestamp
		pbee.SymmetricKey = fb[i].SymmetricKey
		pbee.ProductType = fb[i].ProductType
		pbee.CurrentTime = DateFormat(int(fb[i].CurrentTime))
		pbe = append(pbe, pbee)
	}
	return pbe
}
func (lg *ProductLogModel) TableName() string {
	return "product_log"
}
func OpenPLog() {

	//打开数据库
	DBLog, err = gorm.Open(sqlite.Open("product_log.db"), &gorm.Config{})
	if err != nil {
		log.Logger.Fatalf("open product log db file error:%s", err.Error())
	}
	//将代码中定义的结构体迁移到表数据库表
	err = DBLog.AutoMigrate(&ProductLogModel{})
	if err != nil {
		log.Logger.Fatalf(" autoMigrate product log db file error:%s", err.Error())
	}
	log.Logger.Info("product log db file ini success")
	//}
}
func CreateLog(productID int64, productName string, productType string, currentTime int64, seller string, symmetricKey string,
	buyerCaTimeStamp int64, sellerCaTimeStamp int64, sellerPublicKey string) error {
	tx := DBLog.Debug().Create(&ProductLogModel{ProductID: productID, ProductName: productName, ProductType: productType, CurrentTime: currentTime,
		Seller: seller, SymmetricKey: symmetricKey, BuyerCaTimestamp: buyerCaTimeStamp, SellerCaTimeStamp: sellerCaTimeStamp, SellerPublicKey: sellerPublicKey})
	return tx.Error
}
func FindBatchLog(productID int64) (string, string, error) {
	var fb []ProductLogModel
	tx := DBLog.Debug().Where("product_id = ?", productID).Find(&fb)
	if len(fb) == 0 {
		return "", "", errors.New("not found this product record")
	}
	days := convert(fb)
	marshal, err := json.Marshal(days)
	if err != nil {
		return "", "", errors.New("internal error")
	}

	return fb[0].Seller, string(marshal), tx.Error
}
