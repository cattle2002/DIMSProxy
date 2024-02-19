package model

import (
	"DIMSProxy/log"
	"DIMSProxy/protocol"
	"errors"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type ProductInfoModel struct {
	ProductID             int64  `gorm:"column:product_id;unique;not null;primaryKey"`
	ProductName           string `gorm:"column:product_name;unique;not null"`
	Seller                string `gorm:"column:seller;not null"`
	CanUseNumberLocalFlag bool   `gorm:"column:number_flag;not null"`
	CanUseTimeLocalFlag   bool   `gorm:"column:time_flag;not null"`
	CanUseNumberLocal     int64  `gorm:"column:canuse_number;not null"`
	CanUseTimeLocal       int64  `gorm:"column:canuse_time;not null"`
	Status                string `gorm:"column:status;not null"` //补丁: 完善监控
}

func (p *ProductInfoModel) TableName() string {
	return "product_info"
}

func OpenPInfo() {

	DBInfo, err = gorm.Open(sqlite.Open("product_info.db"), &gorm.Config{})
	if err != nil {
		log.Logger.Fatalf("open product info db file error:%s", err.Error())
	}
	//将代码中定义的结构体迁移到表数据库表
	err = DBInfo.AutoMigrate(&ProductInfoModel{})
	if err != nil {
		log.Logger.Fatalf("autoMigrate product info db error:%s", err.Error())
	}
	log.Logger.Info("product info db file ini  success")
	//}
}
func Find(productID int64, productName string) (bool, error) {
	var pb ProductInfoModel
	tx := DBInfo.Where("product_id = ?  and  product_name = ?", productID, productName).Find(&pb)
	if tx.Error != nil {
		return false, tx.Error
	} else {
		if pb.ProductID == 0 {
			return false, errors.New(protocol.ProductNotFound)
		} else {
			if pb.Status == protocol.Stop {
				return false, nil
			} else if pb.Status == protocol.Renew {
				return true, nil
			} else if pb.Status == protocol.OStatus {
				return true, nil
			}
			return false, errors.New("internal error")
		}
	}
}
func Create(productID int64, productName string, seller string, numberFlag bool, timeFlag bool, canUseNumberLocal int64, canUseTimeLocal int64) error {

	tx := DBInfo.Create(&ProductInfoModel{ProductID: productID, ProductName: productName, Seller: seller, CanUseNumberLocalFlag: numberFlag, CanUseTimeLocalFlag: timeFlag, CanUseNumberLocal: canUseNumberLocal, CanUseTimeLocal: canUseTimeLocal, Status: protocol.OStatus})
	return tx.Error
}

func Conditions(productID int64, productName string) (bool, error) {
	var pb ProductInfoModel
	tx := DBInfo.Where("product_id = ? and product_name = ?", productID, productName).Find(&pb)
	if pb.ProductID == 0 {
		return false, errors.New(protocol.ProductNotFound)
	}
	if tx.Error != nil {
		return false, tx.Error
	}
	if pb.CanUseNumberLocalFlag {
		if pb.CanUseNumberLocal > 0 {
			return true, nil
		} else {
			return false, errors.New("local use number is zero")
		}
	} else {
		milli := time.Now().UnixMilli()
		if milli > pb.CanUseTimeLocal {
			return false, errors.New("local use time is expired ")
		} else {
			return true, nil
		}
	}
}

func ReduceOne(productID int64, productName string) (int64, error) {
	tx := DBInfo.Model(&ProductInfoModel{}).Where("product_id = ? and product_name = ?", productID, productName).Updates(map[string]interface{}{"canuse_number": gorm.Expr("canuse_number - ?", 1)})
	if tx.Error != nil {
		return 0, tx.Error
	}
	var m ProductInfoModel
	first := DBInfo.First(&m, "product_id=?", productID)
	return m.CanUseNumberLocal, first.Error
}
func UpdateStatus(productID int64, user string, status string) (string, error) {
	var pb ProductInfoModel
	//tx := DBLimit.Where("product_id = ?  and seller  = ?", productID, user).Find(&pb)
	tx := DBInfo.Where("product_id = ?", productID).Find(&pb)
	if tx.Error != nil {
		return "", tx.Error
	}
	if pb.ProductID == 0 {
		return "", errors.New("not found record")
	}
	//tx = DBLimit.Model(&ProductDataLimitModel{}).Where("product_id = ? and seller = ?", productID, user).UpdateColumn("status", status)
	tx = DBInfo.Model(&ProductInfoModel{}).Where("product_id = ? ", productID).UpdateColumn("status", status)

	return pb.Seller, tx.Error
}
func GetProductTimes(productID int64) (string, int64, error) {
	var pb ProductInfoModel
	tx := DBInfo.Where("product_id = ?", productID).Find(&pb)
	if pb.ProductID == 0 {
		return "", 0, errors.New("not found record")
	}
	return pb.Seller, pb.CanUseNumberLocal, tx.Error
}
