package fhandle

import (
	"DIMSProxy/log"
	"encoding/json"
	"errors"
	"strconv"
)

type Order struct {
	OrderNumber     string //订单编号
	ProductCode     string //产品编号
	QuantityOrdered string //数量
	PriceEach       string //单价
	OrderLineNumber string //订单行
}
type SubStruct struct {
	OrderNumber     string
	ProductCode     string
	QuantityOrdered string
	PriceEach       string
	OrderLineNumber string
}
type SilimarOrderNumber struct {
	OrderNumber      string
	Count            int64
	MidPrice         float64
	SingleTotalPrice float64
	AllPrice         int64
	TotalCount       int64
}

func SimilarOrderLine(data []SubStruct) []SilimarOrderNumber {
	orderNumberMap := make(map[string][]SubStruct)
	for _, s := range data {
		orderNumberMap[s.OrderNumber] = append(orderNumberMap[s.OrderNumber], s)
	}
	sons := make([]SilimarOrderNumber, 0)
	for _, subStructs := range orderNumberMap {
		// fmt.Printf("OrderNumber: %s\n", orderNumber)
		var son SilimarOrderNumber
		// fmt.Println(len(subStructs))
		for _, subStruct := range subStructs {
			// fmt.Printf("\tProductCode: %s, QuantityOrdered: %s, PriceEach: %s, OrderLineNumber: %s\n", subStruct.ProductCode, subStruct.QuantityOrdered, subStruct.PriceEach, subStruct.OrderLineNumber)
			son.OrderNumber = subStruct.OrderNumber
			sc, e := strconv.Atoi(subStruct.OrderLineNumber)
			if e != nil {
				return sons
			}

			son.TotalCount += int64(sc)
			son.Count += 1

			p, e := strconv.ParseFloat(subStruct.PriceEach, 64)
			if e != nil {
				sons = append(sons, son)
				return sons
			}
			q, e := strconv.ParseFloat(subStruct.QuantityOrdered, 64)
			if e != nil {
				sons = append(sons, son)
				return sons
			}
			son.AllPrice += int64(p * q)
		}
		sons = append(sons, son)
	}
	return sons
}
func CalcResult(jsonData []byte) (string, error) {
	var orders []map[string]interface{}
	if err := json.Unmarshal((jsonData), &orders); err != nil {
		//fmt.Println("Error:", err)
		log.Logger.Debug(1)
		return "", err
	}

	convertedData := map[string]interface{}{
		"Fields": []string{"OrderNumber", "Count", "AllPrice"},
		"Rows":   make([][]interface{}, len(orders)),
	}

	fields, ok := convertedData["Fields"].([]string)
	if !ok {
		//fmt.Println("Error: unable to extract fields")
		return "", errors.New("unable to extract fields")
	}

	for i, order := range orders {
		row := make([]interface{}, len(fields))
		for j, field := range fields {
			row[j] = order[field]
		}
		convertedData["Rows"].([][]interface{})[i] = row
	}

	convertedJSON, err := json.MarshalIndent(convertedData, "", "    ")
	return string(convertedJSON), err
}
func CalcResultX(jsonData []byte) (string, error) {
	var orders []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &orders); err != nil {
		//fmt.Println("Error:", err)
		log.Logger.Debug(1)
		return "", err
	}

	convertedData := map[string]interface{}{
		"Fields": []string{"OrderNumber", "Count", "AllPrice"},
		"Rows":   make([][]interface{}, len(orders)),
	}

	for i, order := range orders {
		row := make([]interface{}, len(convertedData["Fields"].([]string)))
		row[0] = order["OrderNumber"]
		row[1] = order["Count"]
		row[2] = order["AllPrice"]
		convertedData["Rows"].([][]interface{})[i] = row
	}

	convertedJSON, err := json.MarshalIndent(convertedData, "", "    ")
	return string(convertedJSON), err
}
