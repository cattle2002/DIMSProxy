package pfilter

import (
	"DIMSProxy/model"
)

func FilterStatus(productID int64, productName string) (bool, error) {
	find, err := model.Find(productID, productName)
	return find, err
}
func FilterConditions(productID int64, productName string) (bool, error) {
	conditions, err := model.Conditions(productID, productName)
	return conditions, err
}

// FilterInsertPK 查看数据库里面是否存在用户的公钥如果不存在就添加,否则不添加
func FilterInsertPK(user string, timeStamp int64) (bool, error) {
	pk, err := model.FindPK(user, timeStamp)
	return pk, err
}
