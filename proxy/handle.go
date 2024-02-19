package proxy

//func handle(w http.ResponseWriter, r *http.Request) {
//
//}
//
//type CommonReq struct {
//	Buyer     string `json:"Buyer"`
//	ProductID int64  `json:"ProductID"`
//	Demand    string `json:"Demand"`
//}
//type CommonRes struct {
//	Code   int    `json:"Code"`
//	Seller string `json:"Seller"`
//	Demand string `json:"Demand"`
//	Msg    string `json:"Msg"`
//	Data   string `json:"Data"`
//}
//
//func Send(url string, contentType string, body string) (string, string, error) {
//	reader := strings.NewReader(body)
//	resp, err := http.Post(url, contentType, reader)
//	if err != nil {
//		return "", " ", err
//	}
//	boddy, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", " ", err
//	}
//	packet, err := packet.UPacket(boddy)
//	if err != nil {
//		return "", " ", err
//	}
//	var res CommonRes
//	err = json.Unmarshal(packet, &res)
//	if err != nil {
//		return "", " ", err
//	}
//	log.Logger.Tracef("recive http msg:%v", res)
//	if res.Code != 0 {
//		return "", " ", errors.New(res.Msg)
//	}
//	if res.Demand == "times" {
//		return res.Seller, res.Data, nil
//	}
//	if res.Demand == "log" {
//		return res.Seller, res.Data, nil
//	}
//	if res.Demand == "stop" {
//		return res.Seller, " ", nil
//	}
//	if res.Demand == "renew" {
//		return res.Seller, " ", nil
//	}
//	return "", "", nil
//}
