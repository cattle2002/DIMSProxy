package handle

import "github.com/gorilla/websocket"

var WsStatusCh chan bool
var WsConn *websocket.Conn
var WsStatusBool bool
var PInfoCh chan []byte
